package io

import (
	"bufio"
	"io"
)

const (
	readBufferSize = 128
)

type InputReader struct {
	r         io.Reader
	onClose   func()
	onError   func(err error)
	onReceive func([]byte)
}

func NewInputReader(onClose func(), onError func(err error), onReceive func([]byte)) InputReader {
	i := InputReader{}
	i.onClose = onClose
	i.onError = onError
	i.onReceive = onReceive

	return i
}

func (i *InputReader) From(from io.Reader) {
	i.r = bufio.NewReader(from)
}

func (i InputReader) Listen() {
	buf := make([]byte, 0, readBufferSize)

	for {
		n, err := i.r.Read(buf[:cap(buf)])

		if n == 0 && err == io.EOF {
			i.onClose()

			break
		}

		if err != nil && err != io.EOF {
			i.onError(err)
		}

		i.onReceive(buf[:n])
	}
}
