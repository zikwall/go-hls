package http

import (
	"context"
	"github.com/zikwall/go-hls/src/log"
	"net/http"
	"time"
)

type HandlerProvider struct {
	RootDirectory string
	server        *http.Server
}

type ChunkedResponseWriter struct {
	w http.ResponseWriter
}

func (c ChunkedResponseWriter) Write(b []byte) (int, error) {
	n, err := c.w.Write(b)
	c.w.(http.Flusher).Flush()

	return n, err
}

func (h HandlerProvider) WriteError(w http.ResponseWriter, err ...error) {
	w.WriteHeader(http.StatusNotFound)

	if len(err) > 0 && err[0] != nil {
		_, _ = w.Write([]byte(err[0].Error()))
	}
}

func (h *HandlerProvider) Serve() {
	h.server = &http.Server{Addr: ":1338"}

	http.HandleFunc("/", h.PullHandler)

	if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error(err)
	}
}

func (h HandlerProvider) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Info(err)
	}

	log.Info("Graceful shutdown")
}
