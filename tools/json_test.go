package tools

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	obj := new(JSON)
	json.Unmarshal([]byte(`{"a":1515,"arr":["DQQWD",2,"DQW"]}`), obj.Set())
	fmt.Println(obj.Get("arr", 0).(string))
	obj2 := new(JSON)
	json.Unmarshal([]byte(`["DQQWD",2,"DQW"]`), obj2.Set())
	fmt.Println(obj2.Get(0))
}
