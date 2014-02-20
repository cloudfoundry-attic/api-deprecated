package job

import (
	"time"
)

type Model interface {
	Guid() string
	CreatedAt() time.Time
	Url() string
	Status() string
}
