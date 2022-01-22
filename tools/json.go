package tools

import "fmt"

type JSON struct {
	data interface{}
}

func (j JSON) String() string {
	return fmt.Sprint(j.data)
}
func (j *JSON) Set() interface{} {
	return &j.data
}
func (j JSON) JSON(path ...interface{}) *JSON {
	return &JSON{j.Get(path...)}
}
func (j JSON) Get(path ...interface{}) interface{} {
	current := j.data
	for _, name := range path {
		if current != nil {
			switch v := current.(type) {

			case []interface{}:
				if name.(int) >= 0 && len(v) > name.(int) {
					current = v[name.(int)]
				} else {
					return nil
				}
			case map[string]interface{}:
				if n, ok := v[name.(string)]; ok {
					current = n
				} else {
					return nil
				}
			default:
				return nil
			}
		}
	}
	return current
}
