package readable

import "bytes"

// convert string to readable.Readable
func NewStringReadable(s string) Readable {
	return NewBytesBufferReadable(bytes.NewBufferString(s))
}
