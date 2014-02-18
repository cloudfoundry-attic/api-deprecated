package hood_example

import (
	"fmt"
	"github.com/eaigner/hood"
	"os"
)

type Apps struct {
	Id      hood.Id
	Name    string
	Guid    string
	SpaceId int64
	StackId int64
}

type Spaces struct {
	Id hood.Id
}

type Stacks struct {
	Id hood.Id
}

func GenerateUUID() string {
	file, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	file.Read(b)
	file.Close()

	uuid := fmt.Sprintf("%x", b)
	return uuid
}
