package readable

import "bytes"

func NewByteArrayReadable(data []byte) Readable {
	return NewBytesBufferReadable(bytes.NewBuffer(data))
}
