package request

import (
	"net/http"
	"strings"
)

type monday struct {
	Client
	header Header // client 公共请求头
	client *http.Client
}

// Create a instance of a HTTP request client
func NewClient() Client {
	client := new(monday)
	client.header = NewHeader()
	client.client = new(http.Client)
	return client
}
func (m *monday) Header() Header {
	return m.header
}

// Create a simple HTTP GET method request
func (m monday) Get(url string, params map[string]interface{}) Action {
	return m.HTTP("GET", url, params)
}

// Create a simple HTTP Post method request
func (m monday) Post(url string, params map[string]interface{}) Action {
	return m.HTTP("POST", url, params)
}

// Create a request with a method of HTTP that you would like to, etc. GET, POST, PUT...
func (m monday) HTTP(method, url string, params map[string]interface{}) Action {
	return m.NewAction(method, url, params)
}

type mondayAction struct {
	Action
	header Header
	method string
	url    string
	params Params
	client monday
}

func (m monday) NewAction(method, url string, params map[string]interface{}) Action {
	action := new(mondayAction)
	action.url = url
	action.method = strings.ToUpper(method)
	action.params = NewParams(params)
	action.header = m.Header().Clone()
	action.client = m
	return action
}

func (ma mondayAction) Header() Header {
	return ma.header
}
func (ma mondayAction) Do(body Body) (*http.Response, error) {
	// body must be empty if HTTP method equal GET or OPTIONS
	if ma.method == "GET" || ma.method == "OPTIONS" {
		body = nil
	}
	if rq, err := http.NewRequest(ma.method, ma.url, body); err == nil {
		if body != nil {
			rq.ContentLength = body.ContentLength()
			body.SetHeader(ma.header)
			defer body.Reset()
		}
		query := rq.URL.Query()
		for k, v := range ma.params.Items() {
			query.Set(k, v)
		}
		rq.URL.RawQuery = query.Encode()
		for k, v := range ma.header.Items() {
			rq.Header.Set(k, v)
		}
		return ma.client.client.Do(rq)
	} else {
		return nil, err
	}

}
