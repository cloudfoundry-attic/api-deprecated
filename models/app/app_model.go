package app

type Model interface {
	Getter
	Setter
}

const (
	PACKGE_STATE_PENDING = "PENDING"
	PACKGE_STATE_STAGED  = "STAGED"
	PACKGE_STATE_FAILED  = "FAILED"
)

type Getter interface {
	Guid() string
	PackageHash() string
	PackageState() string
}

type Setter interface {
	SetPackageHash(string)
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

func (m *model) PackageHash() string {
	return m.record.PackageHash
}

func (m *model) PackageState() string {
	return m.record.PackageState
}

// SETTERS
func (m *model) SetPackageHash(hash string) {
	m.record.PackageHash = hash
	m.record.PackageState = PACKGE_STATE_PENDING
}

// model.RECORDER
func (m *model) Record() interface{} {
	return &m.record
}

func (m *model) SetRecord(r interface{}) {
	rec, ok := r.(*Record)
	if !ok {
		panic("record must be of type Record")
	}
	m.record = *rec
}
