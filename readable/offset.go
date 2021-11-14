package readable

import "fmt"

type offsetReadable struct {
	readable Readable
	offset   int64
	readed   int64
}

// offset data position
// as possible as do not use this function
func NewOffsetReadable(readable Readable, offset int64) Readable {
	offsetReadable := new(offsetReadable)

	offsetReadable.readable = readable
	if offset > 0 {
		if offset >= readable.Len() {
			return NewLenReadable(readable, 0)
		}
		offsetReadable.offset = offset
		offsetReadable.readed = offset

	}

	return offsetReadable
}
func (o *offsetReadable) Read(p []byte) (int, error) {
	var num int
	var err error
	for num, err = o.readable.Read(p); o.readed > 0; num, err = o.readable.Read(p) {
		o.readed -= int64(num)
		if o.readed <= 0 {
			num, _ = NewByteArrayReadable(p[num+int(o.readed) : num]).Read(p)
			break
		}
	}

	return num, err
}
func (o offsetReadable) Len() int64 {
	return o.readable.Len() - o.offset
}
func (o offsetReadable) String() string {
	return fmt.Sprintf("[BINARY_DATA(%d)]", o.Len())
}
func (o offsetReadable) Reset() {
	o.readable.Reset()
}
