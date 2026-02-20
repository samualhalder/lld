package notificationchannel

import "github.com/samualhalder/lld/notification_system/notification"

type NotificationChannel interface {
	Send(notification.Notification) error
}
