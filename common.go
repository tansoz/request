package request

const VERSION = "1.2.1"

var HEADERS = map[string][]string{
	"User-Agent": {"request/" + VERSION},
}
