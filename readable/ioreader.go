package readable

import (
	"bytes"
	"io"
)

func NewIOReaderReadable(reader io.Reader) Readable {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)
	return NewBytesBufferReadable(buffer)
}
