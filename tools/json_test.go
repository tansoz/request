package tools

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	obj := new(JSON)
	json.Unmarshal([]byte(`{"author":"RommHui","name":"request","version":[1,2,0]}`), obj.Set())
	fmt.Println(obj.Get("author").(string))
	fmt.Println(obj.Get("version", 1))
	fmt.Println("obj:", obj.Get())

	subobj := obj.JSON("version")
	fmt.Println("subobj:", subobj)
	fmt.Println(subobj.Get(1))

	obj2 := new(JSON)
	json.Unmarshal([]byte(`["RommHui",1.20,"request"]`), obj2.Set())
	fmt.Println(obj2.Get(0))
	fmt.Println(obj2.Get(2))
	fmt.Println(obj2.Get())
}
