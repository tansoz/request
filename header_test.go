package request_test

import (
	"fmt"
	"testing"

	"github.com/tansoz/request"
)

func TestNewHeader(t *testing.T) {

	header := request.NewHeader()
	fmt.Println(header)
	header.Items()["good"] = "good"
	header.Set("user-agent", "request/1.0")
	fmt.Println(header)

}
