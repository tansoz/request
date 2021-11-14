package request

const VERSION = "1.1.2"

var HEADERS = map[string][]string{
	"User-agent": {"request/" + VERSION},
}
