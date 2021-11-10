package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
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
	Reset()
}

type ReadCloseLener interface {
	Len() int
	Read([]byte) (int, error)
	Reset()
}

type rawBody struct {
	Body
	data    ReadCloseLener
	headers map[string]string
}

func NewRawBody(data ReadCloseLener, headers map[string]string) Body {
	body := new(rawBody)

	body.data = data
	body.headers = headers

	return body
}
func (r rawBody) Read(p []byte) (int, error) {
	return r.data.Read(p)
}
func (r rawBody) SetHeader(header Header) {
	for k, v := range r.headers {
		header.Set(k, v)
	}
}
func (r rawBody) ContentLength() int64 {
	return int64(r.data.Len())
}
func (r rawBody) Reset() {
	r.data.Reset()
}

// A object encode to JSON as HTTP request body
func NewJSONBody(data interface{}) Body {
	p := new(bytes.Buffer)
	json.NewEncoder(p).Encode(data)
	return NewRawBody(p, map[string]string{
		"content-type": "application/json; charset=utf-8",
	})
}

// A map encode to http query as HTTP request body
func NewQueryBody(data map[string]interface{}) Body {
	p := new(bytes.Buffer)
	query := make(url.Values)
	for k, v := range NewParams(data).Items() {
		query.Set(k, v)
	}
	p.Write([]byte(query.Encode()))
	return NewRawBody(p, map[string]string{
		"content-type": "application/x-www-form-urlencoded; charset=utf-8",
	})
}

// A object encode to XML document as HTTP request body
func NewXMLBody(data interface{}) Body {
	p := new(bytes.Buffer)
	xml.NewEncoder(p).Encode(data)
	return NewRawBody(p, map[string]string{
		"content-type": "application/xml; charset=utf-8",
	})
}

func boundary() string {
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
		m.boundary = boundary()
	}
	return m.boundary
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
func (m multipartBody) Reset() {

}
