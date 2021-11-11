package request

import (
	"net/http"

	"github.com/tansoz/request/body"
)

type Client interface {
	Get(url string, query map[string]interface{}) (*http.Response, error)
	Post(url string, query map[string]interface{}, body body.Body) (*http.Response, error)
	HTTP(method, url string, query map[string]interface{}, host string) Action
}

type clientImpl struct {
	headers http.Header
	client  *http.Client
}

func NewClient(headers map[string][]string) Client {
	cli := new(clientImpl)

	if headers == nil {
		headers = http.Header(HEADERS).Clone()
	}

	cli.headers = headers
	cli.client = new(http.Client)

	return cli
}
func (cli *clientImpl) Get(url string, query map[string]interface{}) (*http.Response, error) {
	return cli.HTTP("GET", url, query, "").Do(nil)
}
func (cli *clientImpl) Post(url string, query map[string]interface{}, body body.Body) (*http.Response, error) {
	return cli.HTTP("POST", url, query, "").Do(body)
}
func (cli *clientImpl) HTTP(method, url string, query map[string]interface{}, host string) Action {
	return cli.NewAction(method, url, query, host)
}
