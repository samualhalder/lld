package metadata

import "time"

type Metadata struct {
	Size        int
	CreateTime  time.Time
	UpdateTime  time.Time
	Permissions []string
}

func (m *Metadata) AssignCT() {
	m.CreateTime = time.Now()
}
func (m *Metadata) AssignUT() {
	m.UpdateTime = time.Now()
}
