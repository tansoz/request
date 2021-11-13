package request

import (
	"net/http"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/option"
	"github.com/tansoz/request/readable"
	"github.com/tansoz/request/response"
)

type Request interface {
	Do(body.Body) response.Response
}

type requestImpl struct {
	options []option.Option
}

func Go(options ...option.Option) Request {
	request := new(requestImpl)

	request.options = options

	return request
}
func (r *requestImpl) Do(content body.Body) response.Response {

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
