package notificationchannel

import (
	"fmt"

	"github.com/samualhalder/lld/notification_system/notification"
)

type MailChannel struct {
}

func (m *MailChannel) Send(notification notification.Notification) error {
	sendTo := notification.GetSender()
	message := notification.GetPayload()
	if sendTo == "" || message == "" {
		return fmt.Errorf("all filedls in message not presant")
	}
	fmt.Printf("Sending a mail to : %s Body is : %s\n", sendTo, message)
	return nil
}
