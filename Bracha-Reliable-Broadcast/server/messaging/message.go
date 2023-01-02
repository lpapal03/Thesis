package messaging

import (
	"errors"
)

type Message struct {
	Sender   string
	Receiver string
	Tag      string
	Content  []string
}

func CreateMessageString(tag string, content []string) []string {
	return append([]string{tag}, content...)
}

func StringToMessage(m []string) (Message, error) {
	if len(m) == 0 {
		return Message{}, errors.New("Empty message")
	}
	return Message{Sender: m[0], Receiver: "", Tag: m[1], Content: m[2:]}, nil
}
