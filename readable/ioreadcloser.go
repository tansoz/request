package readable

import "io"

type ioReadCloserReadable struct {
	Readable
	closer io.ReadCloser
}

// convert io.ReadCloser to readable.Readable
func NewIOReadCloserReadable(IOReadCloser io.ReadCloser, len int64, self bool) Readable {
	readable := new(ioReadCloserReadable)

	readable.Readable = NewIOReaderReadable(IOReadCloser, len, self)
	readable.closer = IOReadCloser

	return readable
}
func (i ioReadCloserReadable) Reset() {
	i.closer.Close()
}
