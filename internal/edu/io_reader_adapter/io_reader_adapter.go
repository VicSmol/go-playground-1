package io_reader_adapter

type IoReaderAdapter interface {
	BytesRead() int
	LinesRead() int
}
