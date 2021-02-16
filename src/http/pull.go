package http

import (
	"github.com/zikwall/go-hls/src/io"
	nativeio "io"
	"log"
	"net/http"
	"path"
	"strings"
)

func (h HandlerProvider) PullHandler(w http.ResponseWriter, r *http.Request) {
	defer w.(http.Flusher).Flush()

	paths := strings.Split(r.URL.String(), "/")

	file, info, exist, err := io.GetFile(
		path.Join(h.RootDirectory, paths[len(paths)-1]),
	)

	if !exist {
		h.WriteError(w, err)

		return
	}

	contentType, err := io.GetFileContentType(file)

	if err != nil {
		h.WriteError(w, err)

		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Last-Modified", info.ModTime().Format(http.TimeFormat))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	if _, err := nativeio.Copy(ChunkedResponseWriter{w}, file); err != nil {
		log.Println(err)
	}
}
