package request_test

import (
	"bytes"
	"fmt"
	"request"
	"testing"
)

func TestNewClient(t *testing.T) {

	cli := request.NewClient()
	fmt.Println(cli.Header())
	if resp := cli.Get("http://www.tansoz.cn/?r=asd&", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewJSONBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		fmt.Println(resp)
	}
	if resp := cli.Get("http://sw.rommhui.com/info.php", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewJSONBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	}
	if resp := cli.Post("http://sw.rommhui.com/info.php", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewQueryBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	}

}
