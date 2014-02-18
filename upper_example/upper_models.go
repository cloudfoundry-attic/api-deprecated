package upper_example

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
	Id      int64  `field:"id"`
	Name    string `field:"name"`
	Guid    string `field:"guid"`
	SpaceId int64  `field:"space_id"`
	StackId int64  `field:"stack_id"`
}

type Space struct {
	Id int64 `db:"id"`
}

type Stack struct {
	Id int64 `db:"id"`
}
