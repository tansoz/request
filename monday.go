package request

import (
	"net/http"
	"strings"
)

type monday struct {
	Client
	header Header // client 公共请求头
}

func NewClient() Client {
	client := new(monday)
	client.header = NewHeader()
	return client
}
func (m *monday) Header() Header {
	return m.header
}
func (m monday) Get(url string, params map[string]interface{}) Action {
	return m.HTTP("GET", url, params)
}
func (m monday) Post(url string, params map[string]interface{}) Action {
	return m.HTTP("POST", url, params)
}
func (m monday) HTTP(method, url string, params map[string]interface{}) Action {
	return m.NewAction(method, url, params)
}

type mondayAction struct {
	Action
	header Header
	method string
	url    string
	params Params
}

func (m monday) NewAction(method, url string, params map[string]interface{}) Action {
	action := new(mondayAction)
	action.url = url
	action.method = strings.ToUpper(method)
	action.params = NewParams(params)
	action.header = m.Header().Clone()
	return action
}

func (ma mondayAction) Header() Header {
	return ma.header
}
func (ma mondayAction) Do(body Body) *http.Response {
	// body must be empty if HTTP method equal GET
	if ma.method == "GET" {
		body = nil
	}
	if rq, err := http.NewRequest(ma.method, ma.url, body); err == nil {
		if body != nil {
			rq.ContentLength = body.ContentLength()
			body.SetHeader(ma.header)
		}
		query := rq.URL.Query()
		for k, v := range ma.params.Items() {
			query.Set(k, v)
		}
		rq.URL.RawQuery = query.Encode()
		for k, v := range ma.header.Items() {
			rq.Header.Set(k, v)
		}
		if resp, err := http.DefaultClient.Do(rq); err == nil {
			return resp
		}
	}

	return nil
}
