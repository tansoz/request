package request_test

import (
	"fmt"
	"request"
	"testing"
)

func TestNewHeader(t *testing.T) {

	header := request.NewHeader()
	fmt.Println(header)
	header.Items()["good"] = "good"
	header.Set("user-agent", "request/1.0")
	fmt.Println(header)

}
