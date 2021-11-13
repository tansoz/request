package body

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewMultipartBodyField(t *testing.T) {

	f := MultipartBodyField("good", []Field{
		QueryBodyField("data", map[string]interface{}{
			"id": 15,
		}, nil),
	}, "", map[string][]string{"add": {"的权威渠道"}})
	fmt.Println(f.Len())
	fmt.Println(f)

	b := MultipartBody([]Field{f, FileBodyField("file", "ok.png", "../testfile.txt", "", 0, -1, nil)}, "", nil)
	fmt.Println(b.Len())

	fmt.Println(b)
}
func TestNewMultipartBodyField2(t *testing.T) {
	b := MultipartBody([]Field{
		FileBodyField("file", "blob.png", "../testfile.txt", "plain/text", 0, -1, map[string][]string{"ok": {"Hello"}}),
	}, "", map[string][]string{"x-qwdqw": {"完全的青蛙的"}})
	fmt.Println(b.Len())
	fmt.Println(b)
	var tmp bytes.Buffer
	tmp.ReadFrom(b)
	fmt.Println(tmp.Len())
}
