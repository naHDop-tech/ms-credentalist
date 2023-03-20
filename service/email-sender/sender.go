package email_sender

type Sender interface {
	Sent(from string, to []string, msg []byte) error
}
