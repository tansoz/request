package request_test

import (
	"fmt"
	"request"
	"testing"
)

func TestNewParams(t *testing.T) {
	params := request.NewParams(nil)
	fmt.Println(params)
	params.Set("id", 1)
	params.Set("url", "http://www.baidu.com")
	params.Set("ur?l", "http://www.baidu.com")
	fmt.Println(params)
}