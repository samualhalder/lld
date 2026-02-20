package notification

type MailNotification struct {
	Email  string
	Body   string
	Status string
}

func NewMail(sendTo string, body string) *MailNotification {
	return &MailNotification{
		Email:  sendTo,
		Body:   body,
		Status: "Pending",
	}
}

func (m *MailNotification) GetSender() string {
	return m.Email
}
func (m *MailNotification) GetPayload() string {
	return m.Body
}
