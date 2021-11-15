package request

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/option"
)

var URL = "http://127.0.0.1/info.php"

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
		option.URL("https://httpbin.org/post"),
		option.Query(map[string]interface{}{
			"id": 262626,
		}),
		option.Timeout(15000),
		// option.ProxyURL("socks5://127.0.0.1:6565"),
		option.Cookies(map[string]string{
			"asd":   "wqdwd",
			"asd48": "wqdwd",
		}),
		option.Headers(map[string][]string{
			"power-by":       {"rommhui"},
			"cookie":         {"rommhui"},
			"content-length": {"1561"},
			"user-agent":     {"1561"},
		}),
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
func TestPost2(t *testing.T) {
	multipartBody := body.MultipartBody(
		[]body.Field{
			body.FileBodyField("files", "testfile.txt", "go.mod", "text/plain", 0, -1, nil),
			body.ParamField("id", 15, nil),
			body.QueryBodyField("data", map[string]interface{}{
				"str":    "hello world",
				"number": 15,
				"bool":   false,
			}, nil),
		},
		"",
		nil,
	)
	if resp := New(
		option.Method("POST"),
		option.URL("https://httpbin.org/post"),
		option.Headers(map[string][]string{
			"User-Agent": {"the request library by Romm Hui"},
		}),
	).Go(multipartBody); resp.Error() == nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Readable())
		fmt.Println(tmp.String())
	} else {
		fmt.Println(resp.Error())
	}
}
