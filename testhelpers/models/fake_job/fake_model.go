package fake_job

import "time"

type Model struct {
	Outputs struct {
		Guid      string
		CreatedAt time.Time
		Url       string
		Status    string
	}
}

func (m *Model) Guid() string {
	return m.Outputs.Guid
}
func (m *Model) CreatedAt() time.Time {
	return m.Outputs.CreatedAt
}
func (m *Model) Url() string {
	return m.Outputs.Url
}
func (m *Model) Status() string {
	return m.Outputs.Status
}
