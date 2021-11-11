package readable

import (
	"bytes"
)

type bufferReadable struct {
	buffer *bytes.Buffer
	len    int64
}

func NewBytesBufferReadable(buffer *bytes.Buffer) Readable {

	readable := new(bufferReadable)
	readable.buffer = buffer
	readable.len = int64(buffer.Len())

	return readable
}
func (b *bufferReadable) Read(p []byte) (int, error) {
	return b.buffer.Read(p)
}
func (b bufferReadable) Len() int64 {
	return b.len
}
func (b *bufferReadable) String() string {
	return b.buffer.String()
}
func (b *bufferReadable) Reset() {
	b.buffer.Reset()
}
