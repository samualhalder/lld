package content

type ContentType interface {
	GetType() string
	GetContent() string
}

type Voice struct {
	Voice string
}

func (v *Voice) GetType() string {
	return "voice"
}
func (v *Voice) GetContent() string {
	return v.Voice
}

type Text struct {
	Text string
}

func (t *Text) GetType() string {
	return "text"
}
func (t *Text) GetContent() string {
	return t.Text
}
