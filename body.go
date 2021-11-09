package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)

type Body interface {
	io.Reader
	SetHeader(header Header)
	ContentLength() int64
}

type jsonBody struct {
	Body
	data *bytes.Buffer
	len  int
}

func NewJSONBody(data interface{}) Body {
	body := new(jsonBody)
	body.data = new(bytes.Buffer)
	json.NewEncoder(body.data).Encode(data)
	body.len = body.data.Len()
	return body
}

func (j jsonBody) Read(bytes []byte) (int, error) {
	return j.data.Read(bytes)
}
func (j jsonBody) ContentLength() int64 {
	return int64(j.len)
}
func (j jsonBody) SetHeader(header Header) {
	header.Set("content-type", "application/json; charset=utf-8")
}

type queryBody struct {
	Body
	data *bytes.Buffer
	len  int
}

func NewQueryBody(data map[string]interface{}) Body {

	body := new(queryBody)

	body.data = new(bytes.Buffer)
	query := make(url.Values)
	for k, v := range NewParams(data).Items() {
		query.Set(k, v)
	}
	body.len, _ = body.data.Write([]byte(query.Encode()))

	return body
}
func (q queryBody) ContentLength() int64 {
	return int64(q.len)
}
func (q queryBody) Read(bytes []byte) (int, error) {
	return q.data.Read(bytes)
}
func (q queryBody) SetHeader(header Header) {
	header.Set("content-type", "application/x-www-form-urlencoded; charset=utf-8")
}
