package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/Lucas-Palomo/go-discord-logger/internal"
	"log"
	"net/http"
)

type Webhook struct {
	internal.Dispatcher
	url string
}

func NewWebhook(url string) *Webhook {
	return &Webhook{
		url: url,
	}
}

func (webhook *Webhook) Send(content string) {

	type Payload struct {
		Content string `json:"content"`
	}

	payload, err := json.Marshal(Payload{Content: content})

	if err != nil {
		log.Fatalln("Failed to prepare message", err)
	}

	req, err := http.NewRequest("POST", webhook.url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln("Failed to send message", err)
	}

}
