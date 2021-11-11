package readable

import "bytes"

func NewStringReadable(s string) Readable {
	return NewBytesBufferReadable(bytes.NewBufferString(s))
}
