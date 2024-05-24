package message

import (
	"fmt"
	"github.com/ztrue/tracerr"
	"strings"
	"time"
)

type Message struct {
	subject    string
	level      string
	moment     time.Time
	content    []string
	error      error
	stackTrace string
}

// NewMessage is a constructor for Message
func NewMessage(subject string, err error, content ...string) *Message {
	message := &Message{
		subject: subject,
		moment:  time.Now(),
		content: content,
	}
	if err != nil {
		message.addError(err)
	}

	return message
}

// SetLevel defines the severity of a Message, automatically defined by the Logger
func (message *Message) SetLevel(level string) {
	message.level = level
}

// addError is used to add an error and the stacktrace
func (message *Message) addError(err error) {

	text := ""
	stacktrace := tracerr.StackTrace(tracerr.Wrap(err))

	for i := 0; i < len(stacktrace); i++ {
		text = text + "\n\t" + stacktrace[i].String()

	}

	message.error = err
	message.stackTrace = text
}

// GetSubject return the Message subject
func (message *Message) GetSubject() string {
	return message.subject
}

// GetLevel return the Message level
func (message *Message) GetLevel() string {
	return message.level
}

// GetMoment return the Message moment
func (message *Message) GetMoment() time.Time {
	return message.moment
}

// GetContent return the Message content
func (message *Message) GetContent() []string {
	return message.content
}

// GetError return the Message error
func (message *Message) GetError() error {
	return message.error
}

// GetStackTrace return the Message stackTrace
func (message *Message) GetStackTrace() string {
	return message.stackTrace
}

// Sprint return the Message formatted as a string
func (message *Message) Sprint() string {
	return fmt.Sprintf(
		"\n```txt\n"+
			"Subject: %s\n"+
			"Level: %s\n"+
			"Moment: %s\n"+
			"Content:\n\t%s\n"+
			"Error: \n\t%v\n"+
			"Stack Trace:%s"+
			"\n```",
		message.subject,
		message.level,
		message.moment.Format(time.RFC3339),
		strings.Join(message.content, "\n\t"),
		message.error,
		message.stackTrace,
	)
}
