package notificationchannel

func GetChannel(channel string) NotificationChannel {
	switch channel {
	case "sms":
		return &MessageChannel{}
	case "mail":
		return &MailChannel{}
	default:
		return nil
	}
}
