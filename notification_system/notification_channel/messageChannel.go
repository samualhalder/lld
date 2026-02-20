package notificationchannel

import (
	"fmt"

	"github.com/samualhalder/lld/notification_system/notification"
)

type MessageChannel struct {
}

func (m *MessageChannel) Send(notification notification.Notification) error {
	sendTo := notification.GetSender()
	message := notification.GetPayload()
	if sendTo == "" || message == "" {
		return fmt.Errorf("all filedls in message not presant")
	}
	fmt.Printf("Sending a message to : %s meassage is : %s\n", sendTo, message)
	return nil
}
