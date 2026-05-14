package strategies

import "github.com/samualhalder/lld/splitwise/models"

type EqualSplit struct {
}

func (e *EqualSplit) Split(users []*models.User, amount float32, metadata map[*models.User]int) (map[*models.User]float32, error) {
	resultMap := make(map[*models.User]float32)
	totalUsers := len(users)
	split := amount / float32(totalUsers)
	for _, user := range users {
		resultMap[user] = split
	}
	return resultMap, nil
}
