package discord_logger

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

type Logger struct {
	token     string
	channel   string
	session   *discordgo.Session
	formatter func(message *Message) string
}

// NewLogger is a constructor for Logger
//
// Case an error occurs while creating a new discord session, a call to log.Fatal will be triggered.
func NewLogger(token string, channel string) *Logger {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalln("Failed to create a discord session,", err)
	}

	return &Logger{
		token:   token,
		channel: channel,
		session: session,
	}
}

// SetFormatter defines the string ouput format of the Message
func (logger *Logger) SetFormatter(formatter func(message *Message) string) {
	logger.formatter = formatter
}

// Sprint returns the formatted Message
func (logger *Logger) Sprint(message *Message) string {
	if logger.formatter != nil {
		return logger.formatter(message)
	}
	return message.Sprint()
}

// SendMessage send a Message
//
// Case an error occurs while sending the message, a call to log.Fatal will be triggered.
func (logger *Logger) SendMessage(message *Message) {
	_, err := logger.session.ChannelMessageSend(logger.channel, logger.Sprint(message))
	if err != nil {
		log.Fatalln("Failed to send message", err)
	}
}

// Info Send a Message with the severity of INFO
func (logger *Logger) Info(message *Message) {
	message.SetLevel("INFO")
	logger.SendMessage(message)
}

// Warn Send a Message with the severity of WARN
func (logger *Logger) Warn(message *Message) {
	message.SetLevel("WARN")
	logger.SendMessage(message)
}

// Error Send a Message with the severity of ERROR
func (logger *Logger) Error(message *Message) {
	message.SetLevel("ERROR")
	logger.SendMessage(message)
}

// Fatal Send a Message with the severity of FATAL followed by a call to os.Exit(1)
func (logger *Logger) Fatal(message *Message) {
	message.SetLevel("FATAL")
	logger.SendMessage(message)
	os.Exit(1)
}

// Panic Send a Message with the severity of Panic followed by a call to panic()
func (logger *Logger) Panic(message *Message) {
	message.SetLevel("PANIC")
	logger.SendMessage(message)
	panic(message.error)
}
