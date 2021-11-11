package readable

type httpReadable struct {
	Readable
}

// not polyfilling yet
func NewHTTPReadable(method, url string) Readable {

	readable := new(httpReadable)

	return readable
}
