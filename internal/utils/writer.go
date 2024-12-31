package utils

import (
	"io"
	"os"
)

type RealTimeWriter struct {
	w io.Writer
}

func NewRealTimeWriter(w io.Writer) *RealTimeWriter {
	return &RealTimeWriter{w: w}
}

func (w *RealTimeWriter) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	if err != nil {
		return n, err
	}
	// Ensure immediate output flush
	if f, ok := w.w.(*os.File); ok {
		f.Sync()
	}
	return n, nil
}
