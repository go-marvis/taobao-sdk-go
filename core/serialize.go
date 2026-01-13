package core

import "encoding/json"

type Serializer interface {
	Marshal(val any) ([]byte, error)
	Unmarshal(data []byte, val any) error
}

type defaultSerializer struct{}

func (d *defaultSerializer) Marshal(val any) ([]byte, error) {
	return json.Marshal(val)
}

func (d *defaultSerializer) Unmarshal(data []byte, val any) error {
	return json.Unmarshal(data, val)
}
