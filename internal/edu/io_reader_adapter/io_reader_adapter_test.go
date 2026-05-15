package io_reader_adapter

import (
	"bytes"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOReaderAdapter(t *testing.T) {
	t.Run("should read bytes and lines", func(t *testing.T) {
		str := "Hello\nWorld\n12345"
		adapter := NewIoReaderAdapterImpl(bytes.NewBufferString(str))
		buffer := make([]byte, 14)
		n, err := adapter.Read(buffer)

		assert.Equal(t, n, 14)
		assert.Equal(t, err, nil)
	})

	t.Run("should return bytes and lines for multiple read", func(t *testing.T) {
		str := "Hello\nWorld\n1234\n"
		adapter := NewIoReaderAdapterImpl(bytes.NewBufferString(str))
		buffer := make([]byte, 5)
		bytes := 0

		for {
			n, err := adapter.Read(buffer)

			bytes += n

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			if n == 0 {
				break
			}

		}

		assert.Equal(t, bytes, len(str))
	})

	t.Run("should return lines", func(t *testing.T) {
		str := "Hello\nWorld\n1234\n"
		adapter := NewIoReaderAdapterImpl(bytes.NewBufferString(str))
		buffer := make([]byte, 5)

		for {
			n, err := adapter.Read(buffer)

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			if n == 0 {
				break
			}

		}

		assert.Equal(t, adapter.LinesRead(), 3)
	})

	t.Run("should return bytes", func(t *testing.T) {
		str := "Hello\nWorld\n1234\n"
		adapter := NewIoReaderAdapterImpl(bytes.NewBufferString(str))
		buffer := make([]byte, 5)

		for {
			n, err := adapter.Read(buffer)

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			if n == 0 {
				break
			}

		}

		assert.Equal(t, adapter.BytesRead(), len(str))
	})
}
