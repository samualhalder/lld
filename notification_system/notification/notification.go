package notification

type Notification interface {
	GetSender() string
	GetPayload() string
}
