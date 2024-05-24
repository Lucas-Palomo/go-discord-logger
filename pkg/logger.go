package discordlogger

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

// sendMessage send a Message
//
// Case an error occurs while sending the message, a call to log.Fatal will be triggered.
func (logger *Logger) sendMessage(message *Message) {
	_, err := logger.session.ChannelMessageSend(logger.channel, logger.Sprint(message))
	if err != nil {
		log.Fatalln("Failed to send message", err)
	}
}

// Info Sends a Message with the severity of INFO
func (logger *Logger) Info(message *Message) {
	message.setLevel("INFO")
	logger.sendMessage(message)
}

// Warn Sends a Message with the severity of WARN
func (logger *Logger) Warn(message *Message) {
	message.setLevel("WARN")
	logger.sendMessage(message)
}

// Error Sends a Message with the severity of ERROR
func (logger *Logger) Error(message *Message) {
	message.setLevel("ERROR")
	logger.sendMessage(message)
}

// Fatal Sends a Message with the severity of FATAL followed by a call to os.Exit(1)
func (logger *Logger) Fatal(message *Message) {
	message.setLevel("FATAL")
	logger.sendMessage(message)
	os.Exit(1)
}

// Panic Sends a Message with the severity of Panic followed by a call to panic()
func (logger *Logger) Panic(message *Message) {
	message.setLevel("PANIC")
	logger.sendMessage(message)
	panic(message.error)
}

// Custom Sends a Message with a customized severity level
func (logger *Logger) Custom(level string, message *Message) {
	message.setLevel(level)
	logger.sendMessage(message)
}
