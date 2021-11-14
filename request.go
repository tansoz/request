package request

import (
	"net/http"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/option"
	"github.com/tansoz/request/readable"
	"github.com/tansoz/request/response"
)

type Request interface {
	Go(body.Body) response.Response
}

type requestImpl struct {
	options []option.Option
}

func New(options ...option.Option) Request {
	request := new(requestImpl)

	request.options = options

	return request
}
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

func Get(url string, query map[string]interface{}) response.Response {
	return New(
		option.GET,
		option.URL(url),
		option.Query(query),
	).Go(nil)
}
func Post(url string, params map[string]interface{}) response.Response {
	return New(
		option.POST,
		option.URL(url),
	).Go(body.QueryBody(params, nil))
}
