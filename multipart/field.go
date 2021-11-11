package multipart

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/common"
	"github.com/tansoz/request/readable"
)

type FieldType uint8

const (
	VALUE FieldType = iota
	FILE
)

type Field interface {
	readable.Readable
	Type() FieldType
}

type fieldImpl struct {
	readable.Readable
	fieldType FieldType
}

func NewRawField(readables []readable.Readable, headers map[string][]string, fieldType FieldType) Field {

	field := new(fieldImpl)

	field.fieldType = fieldType
	bodyReadable := readable.NewListReadable(readables)
	header := http.Header(headers)
	sb := new(strings.Builder)
	header.Set("content-length", fmt.Sprint(bodyReadable.Len()))
	for k, vs := range header {
		for _, v := range vs {
			sb.WriteString(k + ": " + v + "\r\n")
		}
	}
	sb.WriteString("\r\n")
	field.Readable = readable.NewListReadable([]readable.Readable{readable.NewStringReadable(sb.String()), bodyReadable, readable.NewStringReadable("\r\n")})

	return field
}
func (f *fieldImpl) Type() FieldType {
	return f.fieldType
}

func NewNameRawField(name string, content readable.Readable, headers map[string][]string) Field {
	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["content-disposition"] = []string{"form-data; name=\"" + common.Escape(name) + "\""}
	return NewRawField([]readable.Readable{content}, headers, VALUE)
}

func NewQueryBodyField(name string, query map[string]interface{}, headers map[string][]string) Field {
	content := body.NewQueryBody(query, headers)
	return NewNameRawField(name, content, content.Header())
}

func NewMultipartBodyField(name string, fields []Field, boundary string, headers map[string][]string) Field {
	body := NewMultipartBody(fields, boundary, headers)
	return NewNameRawField(name, body, body.Header())
}
func NewFileBodyField(name, filename, path, contentType string, offset, len int64, headers map[string][]string) Field {
	return NewBinaryFileBodyField(name, filename, body.NewFileBody(filename, path, contentType, offset, len, nil), contentType, headers)
}
func NewBinaryFileBodyField(name, filename string, content readable.Readable, contentType string, headers map[string][]string) Field {
	binaryFileBody := body.NewBinaryFileBody(filename, content, contentType, headers)

	header := binaryFileBody.Header()
	header.Set("Content-Disposition", "form-data; name=\""+common.Escape(name)+"\"; filename=\""+common.Escape(filename)+"\"")

	return NewRawField([]readable.Readable{binaryFileBody}, header, FILE)
}
