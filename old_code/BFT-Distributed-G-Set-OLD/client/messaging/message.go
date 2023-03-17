package messaging

import (
	"errors"
)

type Message struct {
	Sender  string
	Tag     string
	Content []string
}

func CreateMessageString(tag string, content []string) []string {
	return append([]string{tag}, content...)
}

func ParseMessageString(m []string) (Message, error) {
	if len(m) == 0 {
		return Message{}, errors.New("Message is empty")
	}
	tag := m[1]
	if tag == GET_RESPONSE || tag == ADD_RESPONSE {
		return Message{Sender: m[0], Tag: m[1]}, nil
	}
	return Message{}, errors.New("error parsing message")
}
