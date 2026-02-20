package main

import (
	"github.com/samualhalder/lld/notification_system/notification"
	notificationchannel "github.com/samualhalder/lld/notification_system/notification_channel"
)

func main() {
	mail := notification.NewMail("sam@gmail.com", "Hi from here")
	sms := notification.NewSms("74398788", "Hello guys")

	mailChannel := notificationchannel.GetChannel("mail")
	smsChannel := notificationchannel.GetChannel("sms")
	mailChannel.Send(mail)
	smsChannel.Send(sms)
}
