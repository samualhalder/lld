package strategies

import "github.com/samualhalder/lld/splitwise/models"

type SplitStrategy interface {
	Split([]*models.User, float32, map[*models.User]int) (map[*models.User]float32, error)
}
