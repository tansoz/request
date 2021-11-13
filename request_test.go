package request

import (
	"fmt"
	"testing"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/option"
)

func TestGo(t *testing.T) {

	Go(
		option.Method("GET"),
		option.URL("http://sw.rommhui.com/info.php"),
		option.Query(map[string]interface{}{
			"good": 262626,
		}),
	).Do(nil).File("info.txt")

}

func TestGo2(t *testing.T) {

	Go(
		option.Method("GET"),
		option.URL("http://sw.rommhui.com/info.php"),
		option.Query(map[string]interface{}{
			"id": 262626,
		}),
		option.Headers(map[string][]string{
			"power-by": {"rommhui"},
		}),
	).Do(nil).File("info.txt")

}

func TestGo3(t *testing.T) {

	r := Go(
		option.POST,
		option.URL("http://sw.rommhui.com/info.php"),
		option.Query(map[string]interface{}{
			"id": 262626,
		}),
		option.Headers(map[string][]string{
			"power-by": {"rommhui"},
		}),
		option.Timeout(15000),
		option.RemoteAddr("1.1.1.30:80"),
		option.Proxy("socks5://127.0.0.1:6565"),
	).Do(body.MultipartBody([]body.Field{
		body.FileBodyField("image", "blob.png", "testfile.txt", "", 0, -1, nil),
		body.ParamField("well", 1516616, nil),
	}, "", nil)).File("info.txt")
	fmt.Println(r)
}
