package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zikwall/go-hls/src/http"
	"github.com/zikwall/go-hls/src/io"
	"github.com/zikwall/go-hls/src/log"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	application := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "root-file-directory",
				Usage:   "",
				EnvVars: []string{"ROOT_FILE_DIRECTORY"},
				Value:   "./tmp",
			},
		},
	}

	application.Action = func(ctx *cli.Context) error {
		absolutePath, err := filepath.Abs(ctx.String("root-file-directory"))

		if err != nil {
			return err
		}

		httpHandlerProvider := http.HandlerProvider{
			RootDirectory: absolutePath,
		}

		go func() {
			httpHandlerProvider.Serve()
		}()

		signal := buildWaitNotifier()

		go func() {
			reader := io.NewInputReader(
				func() {
					log.Info("Close reader")

					// Send a signal to end the application
					signal <- syscall.SIGINT
				},
				func(err error) {
					log.Warning(err)
				},
				func(bytes []byte) {
					log.Info(string(bytes))
				},
			)

			reader.From(io.FromTCP(1339))
			reader.Listen()
		}()

		congratulations()

		<-signal

		httpHandlerProvider.Shutdown()

		log.Info("Stopped")

		return nil
	}

	if err := application.Run(os.Args); err != nil {
		log.Error(err)
	}
}
