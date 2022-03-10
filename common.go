package request

const VERSION = "1.2.0"

var HEADERS = map[string][]string{
	"User-Agent": {"request/" + VERSION},
}
