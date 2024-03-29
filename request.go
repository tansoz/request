package request

import (
	"net/http"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/option"
	"github.com/tansoz/request/readable"
	"github.com/tansoz/request/response"
)

type Request interface {
	Go(body.Body) response.Response // send the HTTP request to the target web server
}

type requestImpl struct {
	options []option.Option
}

// new a request
func New(options ...option.Option) Request {
	request := new(requestImpl)

	request.options = options

	return request
}

// send request and get response
func (r *requestImpl) Go(content body.Body) response.Response {

	client := new(http.Client)
	transport := new(http.Transport)
	request := new(http.Request)

	client.Transport = transport

	option.Headers(HEADERS).Set(request, client, transport)

	for _, opt := range r.options {
		opt.Set(request, client, transport)
	}

	if content != nil {
		option.Headers(content.Header()).Set(request, client, transport)
		request.ContentLength = content.Len()
		request.Body = readable.ConvertToIOReadCloser(content)
	}

	return response.NewResponse(client.Do(request))
}

// GET request
func Get(url string, query map[string]interface{}) response.Response {
	return New(
		option.GET,
		option.URL(url),
		option.Query(query),
	).Go(nil)
}

// POST request
func Post(url string, params map[string]interface{}) response.Response {
	return New(
		option.POST,
		option.URL(url),
	).Go(body.QueryBody(params, nil))
}

// POST file
func PostFile(url string, name string, filename string, path string, mime string) response.Response {
	return New(
		option.POST,
		option.URL(url),
	).Go(body.SimpleMultipartBody(
		body.FileBodyField(name, filename, path, mime, 0, -1, nil),
	))
}
