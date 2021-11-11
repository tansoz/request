package request

const VERSION = "1.0.0"

var HEADERS = map[string][]string{
	"user-agent": {"request/" + VERSION},
}
