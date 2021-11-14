package readable

import (
	"fmt"
	"io"
)

type lenReadable struct {
	readable Readable
	unread   int64
	len      int64
}

// limit readable.Readable readable length
func NewLenReadable(readable Readable, len int64) Readable {
	lenReadable := new(lenReadable)

	lenReadable.readable = readable
	lenReadable.len = readable.Len()
	if len >= 0 && lenReadable.len > len {
		lenReadable.len = len
	}
	lenReadable.unread = lenReadable.len

	return lenReadable
}
func (l *lenReadable) Read(p []byte) (int, error) {
	var num int
	var err error

	if l.unread > 0 {
		num, err = l.readable.Read(p)
		l.unread -= int64(num)
	}
	if l.unread <= 0 {
		num += int(l.unread)
		err = io.EOF
	}

	return num, err
}
func (l *lenReadable) Len() int64 {
	return l.len
}
func (l *lenReadable) String() string {
	return fmt.Sprintf("[BINARY_DATA(%d)]", l.len)
}
func (l *lenReadable) Reset() {
	l.readable.Reset()
}
