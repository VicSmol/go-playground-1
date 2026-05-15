package io_reader_adapter

import (
	"bytes"
	"io"
)

type IoReaderAdapterImpl struct {
	reader io.Reader
	bytes  int
	lines  int
}

func NewIoReaderAdapterImpl(reader io.Reader) *IoReaderAdapterImpl {
	return &IoReaderAdapterImpl{
		reader: reader,
		bytes:  0,
		lines:  0,
	}
}

func (r *IoReaderAdapterImpl) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	r.bytes += n
	r.lines += bytes.Count(p[:n], []byte{'\n'})

	return n, err
}

func (r *IoReaderAdapterImpl) BytesRead() int {
	return r.bytes
}

func (r *IoReaderAdapterImpl) LinesRead() int {
	return r.lines
}
