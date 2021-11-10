package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
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

func newBoundary() string {
	return strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(time.Now().Unix()))), "=", "")
}

type multipartBody struct {
	Body
	boundary string
	fields   []Field
	readpos  int
}

func NewMultipartBody(fields []Field, boundary string) Body {

	body := new(multipartBody)

	body.boundary = boundary
	body.fields = make([]Field, len(fields)*2+1)
	p := 0
	for _, f := range fields {
		body.fields[p] = NewRawFeild("--" + body.getBounary() + "\r\n")
		body.fields[p+1] = f
		p += 2
	}

	if p > 0 {
		body.fields[p] = NewRawFeild("--" + body.getBounary() + "--")
	}

	return body
}
func (m multipartBody) getBounary() string {
	if m.boundary == "" {
		m.boundary = newBoundary()
	}
	return "--" + m.boundary
}
func (m *multipartBody) Read(bytes []byte) (int, error) {
	if l := len(m.fields); l > 0 && m.readpos < l {
		num, err := m.fields[m.readpos].Read(bytes)
		if (err != nil && err == io.EOF) || m.fields[m.readpos].Len() == 0 {
			m.fields[m.readpos].Reset()
			m.readpos += 1
		} else if err != nil {
			return num, err
		}
		return num, nil
	}
	return 0, io.EOF
}
func (m multipartBody) ContentLength() int64 {
	l := int64(0)

	for _, f := range m.fields {
		l += f.Len()
	}

	return l
}
func (m multipartBody) SetHeader(header Header) {
	header.Set("content-type", "multipart/form-data; boundary="+m.getBounary())
}
