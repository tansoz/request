package readable

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewBufferReadable(t *testing.T) {
	tmp := bytes.NewBufferString("Hello World!")
	readable := NewBytesBufferReadable(tmp)
	fmt.Println(readable.Len())
	fmt.Println(readable.String())
	b := make([]byte, 3)
	for num, err := readable.Read(b); err == nil; num, err = readable.Read(b) {
		fmt.Println(num, string(b[:num]))
	}
}

func TestNewFileReadable(t *testing.T) {
	readable := NewFileReadable("testfile.txt", 110)
	fmt.Println(readable.Len())
	fmt.Println(readable.String())
	b := make([]byte, 3)
	for num, err := readable.Read(b); err == nil; num, err = readable.Read(b) {
		fmt.Println(num, string(b[:num]))
	}
}
func TestNewFileReadable2(t *testing.T) {
	readable := NewListReadable([]Readable{
		NewStringReadable("XXXX"),
		NewFileReadable("testfile.txt", 4),
		NewStringReadable("PPPPQWDDQW"),
	})
	fmt.Println(readable.Len())
	fmt.Println(readable.String())
	b := make([]byte, 10)
	for num, err := readable.Read(b); err == nil || num > 0; num, err = readable.Read(b) {
		fmt.Print(string(b[:num]))
	}
	readable = NewLenReadable(NewListReadable([]Readable{
		NewStringReadable("XXXX"),
		NewFileReadable("testfile.txt", 4),
		NewStringReadable("PPPPQWDDQW"),
	}), 8)
	readable = NewLenReadable(readable, 50)
	fmt.Println(readable.Len())
	fmt.Println(readable.String())
	b = make([]byte, 10)
	for num, err := readable.Read(b); err == nil || num > 0; num, err = readable.Read(b) {
		fmt.Print(string(b[:num]))
	}
}
func TestNewOffsetReadable(t *testing.T) {

	readable := NewOffsetReadable(NewStringReadable("Hello World!"), 5)
	fmt.Println(readable.Len())
	fmt.Println(readable)

}

func TestNewIOReaderReadable(t *testing.T) {

	readable := NewIOReaderReadable(NewStringReadable("Hello World!"), 150, false)
	fmt.Println(readable.Len())
	fmt.Println(readable)
	var tmp bytes.Buffer
	tmp.ReadFrom(readable)
	fmt.Println(tmp.Len(), tmp.Bytes())
}

func TestNewIOReadCloserReadable(t *testing.T) {

	readable := NewIOReaderReadable(ioutil.NopCloser(bytes.NewBufferString("0123456789")), -1, true)
	fmt.Println(readable.Len())
	fmt.Println(readable)
	var tmp bytes.Buffer
	tmp.ReadFrom(readable)
	fmt.Println(tmp.Len(), tmp.Bytes())
}
