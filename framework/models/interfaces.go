package models

type Recorder interface {
	Record() interface{}
	SetRecord(interface{})
}
