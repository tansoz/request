package request_test

import (
	"fmt"
	"request"
	"testing"
)

func TestNewValueField(t *testing.T) {
	/*f := request.NewValueField("啊实打实的", 1561156156)
	fmt.Println(f)
	fmt.Println(request.NewValueFields(map[string]interface{}{
		"asda\"sd":   "adqdqw",
		"adqwwqd":    516515616516,
		"adqwwq阿松大d": []byte{1, 15, 60, 49, 40, 0, 185},
		"qwqwddqw":   []string{"asdasd", "dqwwwqdwq", "qwdqwdqwdwdq"},
		"qwqwddqwqwddqw": map[string]interface{}{
			"asdas":       "ASDQWD",
			"QWQWDDWQWDQ": 784848,
		},
	}))*/
	bytes := make([]byte, 3)
	f := request.NewRawFeild("123456789012345678901234567890123456789012345678901234567890123456789015")
	for num, err := f.Read(bytes); err == nil; num, err = f.Read(bytes) {
		fmt.Println(string(bytes[0:num]))
	}
}
func TestNewFileField(t *testing.T) {
	bytes := make([]byte, 30)
	f := request.NewFileFields("files", "blob.png", "field_test.go", "", 0, -1)
	fmt.Println(f)
	for num, err := f.Read(bytes); err == nil; num, err = f.Read(bytes) {
		fmt.Println(string(bytes[0:num]))
	}
}
