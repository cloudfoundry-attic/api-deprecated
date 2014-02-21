package job

import (
	"path"
	"time"
)

type JobModel interface {
	Getter
}

type Recorder interface {
	Record() interface{}
	SetRecord(interface{})
}

type Getter interface {
	Guid() string
	CreatedAt() time.Time
	Url() string
	Status() string
}

type jobModel struct {
	record JobRecord
}

func NewModel() JobModel {
	return NewModelFromRecord(NewJobRecord())
}

func NewModelFromRecord(record JobRecord) JobModel {
	return &jobModel{
		record: record,
	}
}

// GETTERS
func (m *jobModel) Guid() string {
	return m.record.Guid
}

func (m *jobModel) CreatedAt() time.Time {
	return m.record.CreatedAt
}

func (m *jobModel) Url() string {
	return path.Join("/v2/jobs", m.record.Guid)
}

func (m *jobModel) Status() string {
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

// RECORDER
func (m *jobModel) Record() interface{} {
	return &m.record
}

func (m *jobModel) SetRecord(record interface{}) {
	jobRecord, ok := record.(*JobRecord)
	if !ok {
		panic("record must be of type JobRecord")
	}
	m.record = *jobRecord
}
