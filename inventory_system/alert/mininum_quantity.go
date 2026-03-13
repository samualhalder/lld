package alert

import "fmt"

type MinimumQuantity struct {
	MinimumQuantity int
}

func (m *MinimumQuantity) MakeAllert(quantity int) error {
	if quantity <= m.MinimumQuantity {
		return fmt.Errorf("Cant remove this amount of product due to minimun quantity limit")

	}
	return nil
}
