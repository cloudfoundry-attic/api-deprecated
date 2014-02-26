package job

import (
	"path"
	"time"
)

type Model interface {
	Getter
}

type Getter interface {
	Guid() string
	CreatedAt() time.Time
	Url() string
	Status() string
}

type model struct {
	record Record
}

func NewModel() Model {
	return NewModelFromRecord(NewRecord())
}

func NewModelFromRecord(record Record) Model {
	return &model{
		record: record,
	}
}

// GETTERS
func (m *model) Guid() string {
	return m.record.Guid
}

func (m *model) CreatedAt() time.Time {
	return m.record.CreatedAt
}

func (m *model) Url() string {
	return path.Join("/v2/jobs", m.record.Guid)
}

func (m *model) Status() string {
	if m.record.LastError != "" {
		return "failed"
	}
	if m.record.Id == 0 {
		return "finished"
	}
	if m.record.LockedAt.IsZero() {
		return "queued"
	}

	return "running"
}

// model.RECORDER
func (m *model) Record() interface{} {
	return &m.record
}

func (m *model) SetRecord(record interface{}) {
	rec, ok := record.(*Record)
	if !ok {
		panic("record must be of type Record")
	}
	m.record = *rec
}
