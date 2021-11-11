package body

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tansoz/request/common"
	"github.com/tansoz/request/readable"
)

type Body interface {
	readable.Readable
	Header() http.Header
}

type rawBody struct {
	readable.Readable
	headers http.Header
}

func NewRawBody(readable readable.Readable, headers map[string][]string) Body {
	body := new(rawBody)

	body.Readable = readable
	body.headers = headers

	return body
}

func (r *rawBody) Header() http.Header {
	return r.headers
}

func NewQueryBody(query map[string]interface{}, headers map[string][]string) Body {

	vquery := make(url.Values)
	for k, v := range query {
		vquery.Add(k, fmt.Sprint(v))
	}

	if headers == nil {
		headers = make(map[string][]string)
	}

	headers["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}

	return NewRawBody(readable.NewStringReadable(vquery.Encode()), headers)
}

func NewJSONBody(object interface{}, headers map[string][]string) Body {

	tmp := new(bytes.Buffer)

	if headers == nil {
		headers = make(map[string][]string)
	}

	json.NewEncoder(tmp).Encode(object)

	headers["Content-Type"] = []string{"application/json; charset=utf-8"}

	return NewRawBody(readable.NewBytesBufferReadable(tmp), headers)
}

func NewXMLBody(object interface{}, headers map[string][]string) Body {

	tmp := new(bytes.Buffer)

	if headers == nil {
		headers = make(map[string][]string)
	}

	xml.NewEncoder(tmp).Encode(object)

	headers["Content-Type"] = []string{"application/xml; charset=utf-8"}

	return NewRawBody(readable.NewBytesBufferReadable(tmp), headers)
}

func NewFileBody(filename, path, contentType string, offset, len int64, headers map[string][]string) Body {
	return NewBinaryFileBody(filename, readable.NewLenReadable(readable.NewFileReadable(path, offset), len), contentType, headers)
}

func NewBinaryFileBody(filename string, body readable.Readable, contentType string, headers map[string][]string) Body {
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	if headers == nil {
		headers = make(map[string][]string)
	}

	headers["Content-Disposition"] = []string{fmt.Sprintf("attachment; filename=\"%s\"", common.Escape(filename))}
	headers["Content-Type"] = []string{contentType}

	return NewRawBody(body, headers)
}
