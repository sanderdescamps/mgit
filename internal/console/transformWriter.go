package console

import "io"

type TransformWriter struct {
	io.Writer
	writer io.Writer
	f      func([]byte) []byte
}

func NewTransformWriter(writer io.Writer, f func([]byte) []byte) io.Writer {
	result := &TransformWriter{writer: writer, f: f}
	return result
}

func (t TransformWriter) Write(p []byte) (n int, err error) {
	return t.writer.Write(t.f(p))
}
