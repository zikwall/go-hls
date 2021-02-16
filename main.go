package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zikwall/go-hls/src/http"
	"github.com/zikwall/go-hls/src/log"
	"os"
	"path/filepath"
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

		congratulations()
		waitSystemNotify()

		httpHandlerProvider.Shutdown()

		log.Info("Stopped")

		return nil
	}

	if err := application.Run(os.Args); err != nil {
		log.Error(err)
	}
}
