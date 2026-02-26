package chat

import (
	"fmt"

	"github.com/samualhalder/lld/chat_system/message"
	"github.com/samualhalder/lld/chat_system/user"
)

type OneToOne struct {
	users    []*user.User
	messages []*message.Message
}

func (o *OneToOne) AddUser(user1 *user.User) error {
	if len(o.users) > 2 {
		return fmt.Errorf("only two user can be there")
	}
	o.users = append(o.users, user1)
	return nil
}

func (o *OneToOne) Send(message *message.Message) error {
	o.messages = append(o.messages, message)
	return nil
}

func (o *OneToOne) Get() []*message.Message {
	return o.messages
}
