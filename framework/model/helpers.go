package model

import (
	uuid "github.com/nu7hatch/gouuid"
)

func NewGuid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("unable to generate guid: " + err.Error())
	}
	return uuid.String()
}
