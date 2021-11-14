package response

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"os"

	"github.com/tansoz/request/readable"
)

var (
	ErrHadReaded = errors.New("Response body already was readed")
)

type Response interface {
	Header() http.Header         // response headers
	Code() int                   // response status code
	File(string) error           // save response body to file
	XML(interface{}) error       // parse response body to XML document
	JSON(interface{}) error      // parse response body to JSON
	Readable() readable.Readable // response body as a readable.Readable
	Error() error                // request error
	ContentLength() int64        // response body data length
}
type responseImpl struct {
	response *http.Response
	err      error
	readed   bool
}

// new a response.Response
func NewResponse(response *http.Response, err error) Response {
	resp := new(responseImpl)

	resp.response = response
	resp.err = err
	resp.readed = err != nil

	return resp
}

// get length of the response body
func (resp responseImpl) ContentLength() int64 {
	return resp.response.ContentLength
}

// get request error
func (resp responseImpl) Error() error {
	return resp.err
}

// get response headers
func (resp responseImpl) Header() http.Header {
	return resp.response.Header.Clone()
}

// get response status code
func (resp responseImpl) Code() int {
	return resp.response.StatusCode
}

// save file from response body
func (resp *responseImpl) File(filename string) error {
	if resp.err != nil {
		return resp.err
	}
	var reader readable.Readable
	defer func() {
		if reader != nil {
			reader.Reset()
		}
	}()

	if reader = resp.Readable(); reader == nil {
		return ErrHadReaded
	} else if file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0333); err == nil {
		file.Truncate(0)
		b := make([]byte, 1024)
		for num, readerr := reader.Read(b); readerr == nil || num > 0; num, readerr = reader.Read(b) {
			_, writeerr := file.Write(b[0:num])
			if writeerr != nil {
				return writeerr
			}
		}
		return nil
	} else {
		return err
	}
}

// make response body as XML document
func (resp *responseImpl) XML(i interface{}) error {
	if resp.err != nil {
		return resp.err
	}
	if reader := resp.Readable(); reader != nil {
		defer reader.Reset()
		return xml.NewDecoder(reader).Decode(i)
	}
	return ErrHadReaded
}

// make response body as JSON
func (resp *responseImpl) JSON(i interface{}) error {
	if resp.err != nil {
		return resp.err
	}
	if reader := resp.Readable(); reader != nil {
		defer reader.Reset()
		return json.NewDecoder(reader).Decode(i)
	}
	return ErrHadReaded
}

// get a response body reader
func (resp *responseImpl) Readable() readable.Readable {
	if !resp.readed {
		resp.readed = true
		return readable.NewIOReadCloserReadable(resp.response.Body, resp.response.ContentLength, true)
	}
	return nil
}
