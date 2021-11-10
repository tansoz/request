package request_test

import (
	"fmt"
	"testing"

	"github.com/tansoz/request"
)

func TestNewMultipartBody(t *testing.T) {

	body := request.NewMultipartBody([]request.Field{
		request.NewValueField("qwdqwd", 15616),
	}, "")
	fmt.Println(body)

}
