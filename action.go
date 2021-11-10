package request

import "net/http"

type Action interface {
	Header() Header
	Do(body Body) (*http.Response, error)
}
