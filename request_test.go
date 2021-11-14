package request

import (
	"fmt"
	"testing"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/option"
)

var URL = "http://sw.rommhui.com/info.php"

func TestGo(t *testing.T) {

	New(
		option.Method("GET"),
		option.URL(URL),
		option.Query(map[string]interface{}{
			"good": 262626,
		}),
	).Go(nil).File("info.txt")

}

func TestGo2(t *testing.T) {

	New(
		option.Method("GET"),
		option.URL(URL),
		option.Query(map[string]interface{}{
			"id": 262626,
		}),
		option.Headers(map[string][]string{
			"power-by": {"rommhui"},
		}),
	).Go(nil).File("info.txt")

}

func TestGo3(t *testing.T) {

	r := New(
		option.POST,
		option.URL(URL),
		option.Query(map[string]interface{}{
			"id": 262626,
		}),
		option.Headers(map[string][]string{
			"power-by": {"rommhui"},
		}),
		option.Timeout(15000),
		option.ProxyURL("socks5://127.0.0.1:6565"),
	).Go(body.MultipartBody([]body.Field{
		body.FileBodyField("image", "blob.png", "testfile.txt", "", 0, -1, nil),
		body.ParamField("well", 1516616, nil),
	}, "", nil)).File("info.txt")
	fmt.Println(r)
}
func TestPost(t *testing.T) {
	r := Post(URL, map[string]interface{}{
		"id": 115,
	}).File("info.txt")
	fmt.Println(r)
}
