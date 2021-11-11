package request

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tansoz/request/body"
	"github.com/tansoz/request/multipart"
	"github.com/tansoz/request/readable"
)

func TestNewClient(t *testing.T) {

	cli := NewClient(map[string][]string{
		"asdas": {"adqwqwdwqd", "安德森取得完全的"},
	})
	fmt.Println(cli.Get("http://sw.rommhui.com", nil))
}
func TestNewClient2(t *testing.T) {

	cli := NewClient(map[string][]string{
		"asdas": {"adqwqwdwqd", "安德森取得完全的"},
	})
	fmt.Println(cli.Post("http://sw.rommhui.com/info.php", map[string]interface{}{}, body.NewRawBody(readable.NewFileReadable("readable/testfile.txt", -1), nil)))

}
func TestNewClient3(t *testing.T) {

	cli := NewClient(map[string][]string{
		"asdas": {"adqwqwdwqd", "安德森取得完全的"},
	})
	fmt.Println(cli.Post("http://sw.rommhui.com/info.php", map[string]interface{}{}, body.NewFileBody("blob.bin", "readable/testfile.txt", "", 0, 5, nil)))

}
func TestNewClient4(t *testing.T) {

	cli := NewClient(map[string][]string{
		"asdas": {"adqwqwdwqd", "安德森取得完全的"},
	})
	if resp, err := cli.HTTP("Post", "http://sw.rommhui.com/info.php", map[string]interface{}{
		"id":     551,
		"asdasd": "48wqdqwdqwd",
	}, "").Do(multipart.NewMultipartBody([]multipart.Field{
		multipart.NewFileBodyField("file", "blob.png", "readable/testfile.txt", "plain/text", 0, -1, map[string][]string{"ok": {"Hello"}}),
		multipart.NewQueryBodyField("data", map[string]interface{}{
			"user_id":  155,
			"username": "15dqwwd",
		}, map[string][]string{
			"asd": {"qwdqwdwdq"},
		}),
	}, "", nil)); err == nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	} else {
		panic(err)
	}

}
func TestNewClient5(t *testing.T) {

	cli := NewClient(nil)
	if resp, err := cli.HTTP("Post", "http://sw.rommhui.com/info.php", map[string]interface{}{
		"id":     551,
		"asdasd": "48wqdqwdqwd",
	}, "www.cloudflare.com").Do(multipart.NewMultipartBody([]multipart.Field{
		// NewFileBodyField("file", "blob.png", "readable/testfile.txt", "plain/text", 0, -1, map[string][]string{"ok": {"Hello"}}),
		multipart.NewMultipartBodyField("files", []multipart.Field{
			multipart.NewFileBodyField("file1", "blob.png", "readable/testfile.txt", "plain/text", 0, -1, map[string][]string{"ok": {"Hello"}}),
			multipart.NewFileBodyField("file2", "blob.png", "readable/testfile.txt", "plain/text", 0, -1, map[string][]string{"ok": {"Hello"}}),
		}, "", nil),
		multipart.NewQueryBodyField("data", map[string]interface{}{
			"user_id":  155,
			"username": "15dqwwd",
		}, map[string][]string{
			"asd": {"qwdqwdwdq"},
		}),
	}, "", nil)); err == nil {
		var tmp bytes.Buffer
		tmp.ReadFrom(resp.Body)
		fmt.Println(tmp.String())
	} else {
		panic(err)
	}

}
