package request

import (
	"strings"
)

type Header interface {
	Set(name, value string) Header
	Get(name string) string
	Items() map[string]string
	Clone() Header
}

func formatName(name string) string {
	return strings.ToLower(name)
}

type headerImpl struct {
	Header
	headers map[string]string
}

func NewHeader() Header {
	header := new(headerImpl)

	header.headers = make(map[string]string)

	for _, h := range HEADERS {
		header.Set(h[0], h[1])
	}

	return header
}
func (header headerImpl) Set(name, value string) Header {
	header.headers[formatName(name)] = value
	return header
}
func (header headerImpl) Get(name string) string {
	return header.headers[formatName(name)]
}
func (header headerImpl) Items() map[string]string {
	return header.items()
}
func (header headerImpl) items() map[string]string {
	items := make(map[string]string)
	for k, v := range header.headers {
		items[k] = v
	}
	return items
}
func (header headerImpl) Clone() Header {

	h := new(headerImpl)

	h.headers = header.Items()

	return h
}
