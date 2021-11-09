package request

type Client interface {
	Header() Header
	Get(url string, params map[string]interface{}) Action
	Post(url string, params map[string]interface{}) Action
	HTTP(method, url string, params map[string]interface{}) Action
}
