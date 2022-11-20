package messaging

type Message struct {
	Sender   string
	Receiver string
	Tag      string
	Content  []string
}

func CreateMessage(tag string, content []string) []string {
	return append([]string{tag}, content...)
}

func ParseMessage(m []string) Message {
	return Message{Sender: m[0], Receiver: "", Tag: m[1], Content: m[2:]}
}
