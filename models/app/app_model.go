package app

type Model interface {
	Getter
}

type Getter interface {
	Guid() string
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
