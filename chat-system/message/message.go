package message

import (
	"time"

	"github.com/samualhalder/lld/chat_system/content"
	"github.com/samualhalder/lld/chat_system/user"
)

type MessageStatus string

const (
	Sent MessageStatus = "Sent"
	Read MessageStatus = "Read"
)

type Message struct {
	User    *user.User
	Status  MessageStatus
	Time    time.Time
	Content content.ContentType
}
