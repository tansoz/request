package body

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tansoz/request/common"
	"github.com/tansoz/request/readable"
)

type FieldType uint8

const (
	VALUE FieldType = iota
	FILE
)

// multipart field
type Field interface {
	readable.Readable
	Type() FieldType
}

type fieldImpl struct {
	readable.Readable
	fieldType FieldType
}

// raw field of multipart
func RawField(readables []readable.Readable, headers map[string][]string, fieldType FieldType) Field {

	field := new(fieldImpl)

	field.fieldType = fieldType
	bodyReadable := readable.NewListReadable(readables)
	header := http.Header(headers)
	sb := new(strings.Builder)
	header.Set("Content-Length", fmt.Sprint(bodyReadable.Len()))
	for k, vs := range header {
		for _, v := range vs {
			sb.WriteString(k + ": " + v + "\r\n")
		}
	}
	sb.WriteString("\r\n")
	field.Readable = readable.NewListReadable([]readable.Readable{readable.NewStringReadable(sb.String()), bodyReadable, readable.NewStringReadable("\r\n")})

	return field
}

// type of field value or file
func (f *fieldImpl) Type() FieldType {
	return f.fieldType
}

// basic field of multipart
func NameRawField(name string, content readable.Readable, headers map[string][]string) Field {
	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["Content-Disposition"] = []string{"form-data; name=\"" + common.Escape(name) + "\""}
	return RawField([]readable.Readable{content}, headers, VALUE)
}

// param field
func ParamField(name string, value interface{}, headers map[string][]string) Field {
	return NameRawField(name, readable.NewStringReadable(fmt.Sprint(value)), headers)
}

// query field, data like 'id=12&name=foo'
func QueryBodyField(name string, query map[string]interface{}, headers map[string][]string) Field {
	content := QueryBody(query, headers)
	return NameRawField(name, content, content.Header())
}

// multipart body field
func MultipartBodyField(name string, fields []Field, boundary string, headers map[string][]string) Field {
	content := MultipartBody(fields, boundary, headers)
	return NameRawField(name, content, content.Header())
}

// file field
func FileBodyField(name, filename, path, contentType string, offset, len int64, headers map[string][]string) Field {
	return BinaryFileBodyField(name, filename, FileBody(filename, path, contentType, offset, len, nil), contentType, headers)
}

// a readable.Readable object as a file
func BinaryFileBodyField(name, filename string, content readable.Readable, contentType string, headers map[string][]string) Field {
	binaryFileBody := NewBinaryFileBody(filename, content, contentType, headers)

	header := binaryFileBody.Header()
	header.Set("Content-Disposition", "form-data; name=\""+common.Escape(name)+"\"; filename=\""+common.Escape(filename)+"\"")

	return RawField([]readable.Readable{binaryFileBody}, header, FILE)
}
