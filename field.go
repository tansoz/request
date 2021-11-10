package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
)

type FieldType uint8

const (
	FILE FieldType = iota
	VALUE
)

type Field interface {
	io.Reader
	Len() int64
	String() string
	Reset()
}

var escape_regexp = regexp.MustCompile("[\\\"]")

func escape(s string) string {
	return escape_regexp.ReplaceAllString(s, "\\${0}")
}
func NewValueField(name string, value interface{}) Field {
	value_str, ext := toString(value)
	lvalue := len([]byte(value_str))
	return NewRawFeild(fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"\r\n%sContent-Length: %d\r\n\r\n%v\r\n", escape(name), ext, lvalue, value_str))
}
func toString(i interface{}) (string, string) {
	switch t := i.(type) {
	case string:
		return t, ""
	case url.Values:
		return t.Encode(), "Content-Type: application/x-www-form-urlencoded; charset=utf-8\r\n"
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return fmt.Sprint(t), ""
	case []byte:
		return string(t), "Content-Type: application/octet-stream\r\n"
	default:
		var tmp bytes.Buffer
		json.NewEncoder(&tmp).Encode(i)
		return tmp.String(), "Content-Type: application/json; charset=utf-8\r\n"

	}
}
func NewValueFields(fields map[string]interface{}) []Field {
	list := []Field{}
	for k, v := range fields {
		list = append(list, NewValueField(k, v))
	}
	return list
}

type rawField struct {
	Field
	data *bytes.Buffer
}

func NewRawFeild(plain string) Field {

	field := new(rawField)

	field.data = new(bytes.Buffer)
	field.data.WriteString(plain)

	return field
}
func (r rawField) Read(bytes []byte) (int, error) {
	return r.data.Read(bytes)
}
func (r rawField) Reset() {
	r.data.Reset()
}
func (r rawField) Len() int64 {
	return int64(r.data.Len())
}
func (r rawField) String() string {
	return r.data.String()
}

type fileFeild struct {
	Field
	fp       *os.File
	data     *bytes.Buffer
	padding  *bytes.Buffer
	len      int64
	list     []io.Reader
	readpos  int
	filesize int64
}

func NewFileFields(name, filename, path, mime string, offset, len int64) Field {
	field := new(fileFeild)

	if fp, err := os.OpenFile(path, os.O_RDONLY, 0666); err == nil {
		field.data = new(bytes.Buffer)
		field.padding = bytes.NewBuffer([]byte("\r\n"))
		field.fp = fp
		field.fp.Seek(offset, 0)
		fileInfo, _ := fp.Stat()
		filesize := fileInfo.Size()
		if offset > 0 {
			filesize -= offset
		}
		if len >= 0 && filesize > len {
			filesize = len
		}
		if mime == "" {
			mime = "application/octet-stream"
		}

		field.data.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"\r\n%sContent-Length: %d\r\n\r\n", escape(name), escape(filename), "Content-Type: "+mime+"\r\n", filesize))

		field.list = []io.Reader{field.data, field.fp, field.padding}
		field.len = int64(field.data.Cap()) + 2 + filesize
		field.filesize = filesize

	} else {
		panic(err)
	}

	return field
}
func (f *fileFeild) Read(bytes []byte) (int, error) {

	if l := len(f.list); l > 0 && f.readpos < l {
		num, err := f.list[f.readpos].Read(bytes)
		if f.readpos == 1 {
			fmt.Println(num)
			f.filesize -= int64(num)
		}
		if (err != nil && err == io.EOF) || f.CheckEOF() {
			if f.readpos == 1 {
				num += int(f.filesize)
			}
			f.readpos += 1
		} else if err != nil {
			return num, err
		}
		return num, nil
	}
	return 0, io.EOF
}
func (f fileFeild) CheckEOF() bool {
	if f.readpos == 0 || f.readpos == 2 {
		return f.data.Len() == 0
	}
	return f.readpos == 1 && f.filesize <= 0
}
func (f fileFeild) Reset() {
	f.data.Reset()
	f.fp.Close()
	f.padding.Reset()
}
func (f fileFeild) Len() int64 {
	return f.len
}
func (f fileFeild) String() string {
	return f.data.String() + "[BINARY_DATA]\r\n"
}
