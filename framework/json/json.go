package json

import (
	"encoding/json"
	"time"
)

const TIMESTAMP_FORMAT = "2006-01-02T15:04:05+00:00"

type Resource struct {
	Metadata Map `json:"metadata"`
	Entity   Map `json:"entity"`
}

type Map map[string]interface{}

func (m Map) MarshalJSON() (result []byte, err error) {
	replaceTimeWithFormattedString(m)
	jsonMap := map[string]interface{}(m)
	return json.Marshal(jsonMap)
}

func (m Map) Set(key, val interface{}) {
	stringKey := key.(string)
	m[stringKey] = val
}

func (m Map) Each(cb func(key, val interface{})) {
	for key, val := range m {
		cb(key, val)
	}
}

type Array []interface{}

func (a Array) Set(key, val interface{}) {
	intKey := key.(int)
	a[intKey] = val
}

func (a Array) Each(cb func(key, val interface{})) {
	for key, val := range a {
		cb(key, val)
	}
}

type Mutable interface {
	Set(key, val interface{})
}

type Iterable interface {
	Each(func(key, val interface{}))
}

type MutableIterator interface {
	Mutable
	Iterable
}

func replaceTimeWithFormattedString(obj interface{}) {

	switch obj := obj.(type) {
	case MutableIterator:
		obj.Each(func(key, val interface{}) {
			timeVal, ok := val.(time.Time)
			if ok {
				obj.Set(key, timeVal.Format(TIMESTAMP_FORMAT))
			} else {
				replaceTimeWithFormattedString(val)
			}
		})
	}
}
