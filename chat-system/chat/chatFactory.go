package chat

import "fmt"

func CreateChat(ty string) (Chat, error) {
	switch ty {
	case "one-to-one":
		return &OneToOne{}, nil
	case "group-chat":
		return &GroupChat{}, nil
	default:
		return nil, fmt.Errorf("no such chat type")
	}
}
