package option

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrRequestNull   = errors.New("the http.Request is null")
	ErrClientNull    = errors.New("the http.Client is null")
	ErrTransportNull = errors.New("the http.Transport is null")
)

type Option interface {
	Set(*http.Request, *http.Client, *http.Transport) error // internal call function
}

type methodOptionImpl struct {
	method string
}

// set method of request
func Method(method string) Option {
	return &methodOptionImpl{method: strings.ToUpper(method)}
}

func (m methodOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if r == nil {
		return ErrRequestNull
	}
	r.Method = m.method
	return nil
}

// default setting request method option
var (
	GET     = Method("GET")
	POST    = Method("POST")
	OPTIONS = Method("OPTIONS")
	PUT     = Method("PUT")
	DELETE  = Method("DELETE")
	HEAD    = Method("HEAD")
	TRACE   = Method("TRACE")
)

type urlOptionImpl struct {
	url string
}

// set URL
func URL(url string) Option {
	return &urlOptionImpl{url: url}
}
func (u urlOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if r == nil {
		return ErrRequestNull
	}
	purl, err := url.Parse(u.url)
	if err != nil {
		return err
	}
	query := purl.Query()
	if r.URL != nil {
		for k, vs := range r.URL.Query() {
			for _, v := range vs {
				query.Add(k, v)
			}
		}
	}
	r.URL = purl
	r.URL.RawQuery = query.Encode()
	return nil
}

type headerOptionImpl struct {
	headers http.Header
}

// set request headers
func Headers(headers map[string][]string) Option {
	return &headerOptionImpl{headers: http.Header(headers)}
}
func (h headerOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if r == nil {
		return ErrRequestNull
	}
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	for k, vs := range h.headers {
		for _, v := range vs {
			r.Header.Add(k, v)
		}
	}
	return nil
}

type queryOptionImpl struct {
	query map[string]interface{}
}

// set query data
func Query(query map[string]interface{}) Option {
	return &queryOptionImpl{query: query}
}
func (q queryOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if r == nil {
		return ErrRequestNull
	}
	if r.URL == nil {
		r.URL = new(url.URL)
	}
	query := r.URL.Query()
	for k, v := range q.query {
		query.Set(k, fmt.Sprint(v))
	}
	r.URL.RawQuery = query.Encode()
	return nil
}

type hostOptionImpl struct {
	host string
}

// set host name
func Host(host string) Option {
	return &hostOptionImpl{host: host}
}
func (h hostOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if r == nil {
		return ErrRequestNull
	}

	r.Host = h.host

	return nil
}

type timeoutOptionImpl struct {
	timeout time.Duration
}

// set request timeout
func Timeout(milliseconds int64) Option {
	option := new(timeoutOptionImpl)
	option.timeout = time.Duration(milliseconds * int64(time.Millisecond))
	return option
}
func (to timeoutOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if c != nil {
		c.Timeout = to.timeout
	} else {
		return ErrClientNull
	}
	return nil
}

type remoteAddrOptionImpl struct {
	addr string
}

// set remote address
// func RemoteAddr(addr string) Option {
// 	option := new(remoteAddrOptionImpl)

// 	option.addr = addr

// 	return option
// }
// func (ra remoteAddrOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
// 	if t == nil {
// 		return ErrTransportNull
// 	}

// 	t.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
// 		return net.Dial(network, ra.addr)
// 	}
// 	t.DialTLS = func(network, addr string) (net.Conn, error) {
// 		return tls.Dial(network, ra.addr, t.TLSClientConfig)
// 	}

// 	return nil
// }

type proxyOptionImpl struct {
	addr string
}

// set proxy
// addr is supported socks5,https,http. examples, socks5://127.0.0.1:8080, https://127.0.0.1:8080, https://127.0.0.1:8080
func ProxyURL(addr string) Option {
	option := new(proxyOptionImpl)

	option.addr = addr

	return option
}
func (p proxyOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {
	if t == nil {
		return ErrTransportNull
	}
	t.Proxy = func(r *http.Request) (*url.URL, error) {
		return url.Parse(p.addr)
	}
	return nil
}

type cookieOptionImpl struct {
	cookies map[string]string
}

// set cookie
func Cookies(cookies map[string]string) Option {

	option := new(cookieOptionImpl)

	option.cookies = cookies

	return option
}
func (ck cookieOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {

	if r == nil {
		return ErrRequestNull
	}

	for k, v := range ck.cookies {
		r.AddCookie(&http.Cookie{Name: k, Value: v})
	}

	return nil
}

type basicAuthOptionImpl struct {
	username string
	password string
}

// set basic auth information username and password
func BasicAuth(username, password string) Option {

	option := new(basicAuthOptionImpl)

	option.username = username
	option.password = password

	return option

}

func (b basicAuthOptionImpl) Set(r *http.Request, c *http.Client, t *http.Transport) error {

	if r == nil {
		return ErrRequestNull
	}

	r.SetBasicAuth(b.username, b.password)
	return nil
}
