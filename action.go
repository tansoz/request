package request

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tansoz/request/body"
)

type Action interface {
	Do(body.Body) (*http.Response, error)
}

type actionImpl struct {
	url    string
	method string
	client *clientImpl
	query  map[string]interface{}
	host   string
	// remoteAddr string
}

func (cli *clientImpl) NewAction(method, url string, query map[string]interface{}, host string) Action {
	action := new(actionImpl)

	action.client = cli
	action.method = strings.ToUpper(method)
	action.url = url
	action.query = query
	action.host = host
	// action.remoteAddr = remoteAddr

	return action
}
func (a *actionImpl) Do(body body.Body) (*http.Response, error) {

	if req, err := http.NewRequest(a.method, a.url, body); err == nil {
		// if a.remoteAddr != "" {
		// 	req.RemoteAddr = a.remoteAddr
		// }
		if a.host != "" {
			req.Host = a.host
		}

		query := req.URL.Query()
		for k, v := range a.query {
			query.Add(k, fmt.Sprint(v))
		}
		req.URL.RawQuery = query.Encode()

		req.Header = a.client.headers.Clone()
		for k, arr := range HEADERS {
			for _, v := range arr {
				req.Header.Add(k, v)
			}
		}
		if body != nil {
			for k, arr := range body.Header() {
				for _, v := range arr {
					req.Header.Add(k, v)
				}
			}
			req.ContentLength = body.Len()
		}
		return a.client.client.Do(req)
	} else {
		return nil, err
	}
}
