package readable

import (
	"bytes"
	"fmt"
	"io"
)

type ioReaderReadable struct {
	io.Reader
	len    int64
	readed int64
	self   bool
}

// convert io.Reader to readable.Readable
func NewIOReaderReadable(reader io.Reader, len int64, self bool) Readable {

	readable := new(ioReaderReadable)
	readable.Reader = reader
	readable.len = len
	readable.self = self
	if len < 0 {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(reader)
		return NewBytesBufferReadable(buffer)
	}
	return readable
}
func (i *ioReaderReadable) Read(p []byte) (int, error) {
	var num int
	var err error
	if i.self {
		return i.Reader.Read(p)
	}
	if i.readed < i.len {
		num, err = i.Reader.Read(p)
		i.readed += int64(num)
		for x := num; x < len(p) && err == io.EOF && i.readed < i.len; x, num, i.readed = x+1, num+1, i.readed+1 {
			p[x] = 0
		}
	}
	if i.readed >= i.len {
		num -= int(i.readed - i.len)
		err = io.EOF
	}
	return num, err
}
func (i ioReaderReadable) Len() int64 {
	return i.len
}
func (i ioReaderReadable) String() string {
	return fmt.Sprintf("[BINARY_DATA(%d)]", i.Len())
}
func (i ioReaderReadable) Reset() {}
