package model

import (
	"fmt"
	"os"
)

func GenerateUUID() string {
	file, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	file.Read(b)
	file.Close()

	uuid := fmt.Sprintf("%x", b)
	return uuid
}

type App struct {
	Id      int64 `db:"id"`
	Name    string
	Guid    string
	SpaceId int64
	StackId int64
}

type Space struct {
	Id int64 `db:"id"`
}

type Stack struct {
	Id int64 `db:"id"`
}
