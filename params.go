package request

import (
	"fmt"
	"net/url"
	"strings"
)

type Params interface {
	String() string
	Set(name string, value interface{}) Params
	Get(name string) interface{}
	Items() map[string]string
}

type paramsImpl struct {
	Params
	params map[string]interface{}
}

func NewParams(params map[string]interface{}) Params {
	p := new(paramsImpl)

	if params == nil {
		p.params = make(map[string]interface{})
	} else {
		p.params = params
	}

	return p
}
func (p paramsImpl) String() string {
	params := []string{}

	for k, v := range p.params {
		params = append(params, url.QueryEscape(k)+"="+url.QueryEscape(fmt.Sprint(v)))
	}

	return strings.Join(params, "&")
}
func (p paramsImpl) Set(name string, value interface{}) Params {
	p.params[name] = value
	return p
}
func (p paramsImpl) Get(name string) interface{} {
	return p.params[name]
}
func (p paramsImpl) Items() map[string]string {
	return p.items()
}
func (p paramsImpl) items() map[string]string {
	items := make(map[string]string)
	for k, v := range p.params {
		items[k] = fmt.Sprint(v)
	}
	return items
}
