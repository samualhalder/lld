package main

import (
	"fmt"
	"time"

	"github.com/samualhalder/lld/chat_system/chat"
	"github.com/samualhalder/lld/chat_system/content"
	"github.com/samualhalder/lld/chat_system/message"
	"github.com/samualhalder/lld/chat_system/user"
)

func main() {
	user1 := user.NewUser(1, "Samual")
	user2 := user.NewUser(2, "Ram")
	user3 := user.NewUser(3, "Sham")
	chat12, err := chat.CreateChat("one-to-one")
	if err != nil {
		println(err.Error())
		return
	}
	chat12.AddUser(user1)
	chat12.AddUser(user2)
	msg := &message.Message{
		User:    user1,
		Status:  message.Sent,
		Content: &content.Text{Text: "Hi from sameal"},
		Time:    time.Now(),
	}
	chat12.Send(msg)
	msg = &message.Message{
		User:    user2,
		Status:  message.Sent,
		Content: &content.Text{Text: "Hi from Ram"},
		Time:    time.Now(),
	}
	chat12.Send(msg)
	messages := chat12.Get()
	for _, msg := range messages {
		fmt.Printf("from %s : %s\n", msg.User.Name, msg.Content.GetContent())
	}
	group, err := chat.CreateChat("group-chat")
	if err != nil {
		println(err.Error())
		return
	}
	group.AddUser(user1)
	group.AddUser(user2)
	group.AddUser(user3)
	group.Send(msg)
	messages = group.Get()
	for _, msg := range messages {
		fmt.Printf("from %s : %s\n", msg.User.Name, msg.Content.GetContent())
	}
}
