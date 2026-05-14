package strategies

import (
	"fmt"

	"github.com/samualhalder/lld/splitwise/enums"
)

func GetSplitStrategy(strategy enums.SplitType) (SplitStrategy, error) {
	switch strategy {
	case enums.EQUAL:
		return &EqualSplit{}, nil
	default:
		return nil, fmt.Errorf("unknown split")
	}
}
