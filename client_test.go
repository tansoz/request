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
	if resp, _ := cli.Get("http://www.tansoz.cn/?r=asd&", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewJSONBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	}
	if resp, _ := cli.Post("http://sw.rommhui.com/info.php", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewJSONBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	}
	if resp, _ := cli.Post("http://sw.rommhui.com/info.php", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewQueryBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		var tmp bytes.Buffer
		fmt.Println(resp)
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	}
	if resp, _ := cli.HTTP("OPTIONS", "http://sw.rommhui.com/info.php", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewQueryBody(map[string]interface{}{
		"id":       1,
		"password": "123456",
	})); resp != nil {
		var tmp bytes.Buffer
		fmt.Println(resp)
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	}

}
func TestNewClient2(t *testing.T) {

	cli := request.NewClient()
	fmt.Println(cli.Header())
	if resp, err := cli.HTTP("POST", "http://sw.rommhui.com/info.php", map[string]interface{}{
		"good": "asdasd",
	}).Do(request.NewMultipartBody([]request.Field{
		request.NewValueField("asdads", 1515),
		request.NewValueField("aads", 815174845),
		request.NewValueField("aads18", []int{15, 1515, 15185, 1818, 84}),
		request.NewFileField("files", "blob.jpg", "field_test.go", "", 900, 151),
	}, "")); err == nil {
		var tmp bytes.Buffer
		fmt.Println(resp)
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	} else {
		fmt.Println(err)
	}

}
