package readable

import "bytes"

// convert []byte to readable.Readble
func NewByteArrayReadable(data []byte) Readable {
	return NewBytesBufferReadable(bytes.NewBuffer(data))
}
