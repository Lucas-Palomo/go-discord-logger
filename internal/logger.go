package internal

import (
	"github.com/Lucas-Palomo/go-discord-logger/pkg/message"
)

type Logger interface {
	Custom(level string, message *message.Message)
	Info(message *message.Message)
	Warn(message *message.Message)
	Error(message *message.Message)
	Fatal(message *message.Message)
	Panic(message *message.Message)
	SetFormatter(formatter func(message *message.Message) string)
	sprint(message *message.Message) string
	send(message *message.Message)
}
