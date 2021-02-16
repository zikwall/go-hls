package main

import (
	"github.com/zikwall/go-hls/src/http"
	"log"
	"path/filepath"
)

func main() {
	absolutePath, err := filepath.Abs("./tmp")

	if err != nil {
		log.Fatal(err)
	}

	httpHandlerProvider := http.HandlerProvider{
		RootDirectory: absolutePath,
	}

	go func() {
		httpHandlerProvider.Serve()
	}()

	waitSystemNotify()

	httpHandlerProvider.Shutdown()

	log.Println("Stopped")
}
