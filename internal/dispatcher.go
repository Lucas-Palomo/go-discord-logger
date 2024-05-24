package internal

type Dispatcher interface {
	Send(content string)
}
