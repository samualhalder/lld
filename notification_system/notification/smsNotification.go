package notification

type SMSNotific struct {
	phoneNumber string
	message     string
	Status      string
}

func NewSms(phone, message string) *SMSNotific {
	return &SMSNotific{
		phoneNumber: phone,
		message:     message,
		Status:      "Pending",
	}
}

func (m *SMSNotific) GetSender() string {
	return m.phoneNumber
}
func (m *SMSNotific) GetPayload() string {
	return m.message
}
