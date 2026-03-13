package models

type History struct {
	logs []*Log
}

func (h *History) Add(log *Log) {
	h.logs = append(h.logs, log)
}

func (h *History) List() []*Log {
	return h.logs
}
