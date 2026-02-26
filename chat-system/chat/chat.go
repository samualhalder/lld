package chat

import (
	"github.com/samualhalder/lld/chat_system/message"
	"github.com/samualhalder/lld/chat_system/user"
)

type Chat interface {
	Send(message *message.Message) error
	Get() []*message.Message
	AddUser(user *user.User) error
}
