package readable

import "io"

type ioReadCloser struct {
	Readable
}

// convert io.ReadCloser to readable.Readable
func ConvertToIOReadCloser(readable Readable) io.ReadCloser {
	return ioReadCloser{readable}
}

func (i ioReadCloser) Close() error {
	i.Reset()
	return nil
}
