package request

const VERSION = "1.0.1"

var HEADERS = map[string][]string{
	"User-agent": {"request/" + VERSION},
}
