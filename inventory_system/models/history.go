package models

type History struct {
	logs []*Log
}

func (h *History) Add(product *Product, opType Operation) {
	log := NewLog(product.SKU, opType)
	h.logs = append(h.logs, log)
}

func (h *History) List() []*Log {
	return h.logs
}
