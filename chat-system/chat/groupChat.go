package chat

import (
	"github.com/samualhalder/lld/chat_system/message"
	"github.com/samualhalder/lld/chat_system/user"
)

type GroupChat struct {
	users    []*user.User
	messages []*message.Message
}

func (o *GroupChat) AddUser(user *user.User) error {
	o.users = append(o.users, user)
	return nil
}

func (o *GroupChat) Send(message *message.Message) error {
	o.messages = append(o.messages, message)
	return nil
}

func (o *GroupChat) Get() []*message.Message {
	return o.messages
}
