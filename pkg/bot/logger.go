package bot

import (
	"github.com/Lucas-Palomo/go-discord-logger/internal"
	"github.com/Lucas-Palomo/go-discord-logger/pkg/message"
	"os"
)

type Logger struct {
	internal.Logger
	bot       *Bot
	formatter func(message *message.Message) string
}

// NewLogger is a constructor for Logger using a Bot
func NewLogger(bot *Bot) *Logger {
	return &Logger{
		bot: bot,
	}
}

// SetFormatter defines the string ouput format of the Message
func (logger *Logger) SetFormatter(formatter func(message *message.Message) string) {
	logger.formatter = formatter
}

// Sprint returns the formatted Message
func (logger *Logger) Sprint(message *message.Message) string {
	if logger.formatter != nil {
		return logger.formatter(message)
	}
	return message.Sprint()
}

// sendMessage send a Message
//
// Case an error occurs while sending the message, a call to log.Fatal will be triggered.
func (logger *Logger) sendMessage(message *message.Message) {
	logger.bot.Send(logger.Sprint(message))
}

// Info Sends a Message with the severity of INFO
func (logger *Logger) Info(message *message.Message) {
	message.SetLevel("INFO")
	logger.sendMessage(message)
}

// Warn Sends a Message with the severity of WARN
func (logger *Logger) Warn(message *message.Message) {
	message.SetLevel("WARN")
	logger.sendMessage(message)
}

// Error Sends a Message with the severity of ERROR
func (logger *Logger) Error(message *message.Message) {
	message.SetLevel("ERROR")
	logger.sendMessage(message)
}

// Fatal Sends a Message with the severity of FATAL followed by a call to os.Exit(1)
func (logger *Logger) Fatal(message *message.Message) {
	message.SetLevel("FATAL")
	logger.sendMessage(message)
	os.Exit(1)
}

// Panic Sends a Message with the severity of Panic followed by a call to panic()
func (logger *Logger) Panic(message *message.Message) {
	message.SetLevel("PANIC")
	logger.sendMessage(message)
	panic(message.GetError())
}

// Custom Sends a Message with a customized severity level
func (logger *Logger) Custom(level string, message *message.Message) {
	message.SetLevel(level)
	logger.sendMessage(message)
}
