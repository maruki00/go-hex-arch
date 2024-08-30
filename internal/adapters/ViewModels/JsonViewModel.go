package viewmodels

import "encoding/json"

type JsonViewModel struct {
	data any
}

func (obj *JsonViewModel) GetResponse() any {
	json.Marshal(obj.data)
	return obj.data
}
