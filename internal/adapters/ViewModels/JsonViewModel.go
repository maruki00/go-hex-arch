package viewmodels

import "encoding/json"

type JsonViewModel struct {
	Data any
}

func (obj *JsonViewModel) GetResponse() any {
	json.Marshal(obj.Data)
	return obj.Data
}
