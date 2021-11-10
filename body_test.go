package request_test

import (
	"fmt"
	"request"
	"testing"
)

func TestNewMultipartBody(t *testing.T) {

	body := request.NewMultipartBody([]request.Field{
		request.NewValueField("qwdqwd", 15616),
	}, "")
	fmt.Println(body)

}
