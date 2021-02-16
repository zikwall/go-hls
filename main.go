package main

import (
	"github.com/zikwall/go-hls/src/http"
	"github.com/zikwall/go-hls/src/log"
	"path/filepath"
)

func main() {
	absolutePath, err := filepath.Abs("./tmp")

	if err != nil {
		log.Error(err)
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
}
