# Discord Logger

This project is a simple logger that uses a Discord bot to send logs to a channel

## Features

- [x] Dispatch messages using a discord bot
- [x] Dispatch messages using a discord webhook
- [x] Error support with stack trace
- [x] Supports log levels such as:
  - Info
  - Warn
  - Error
  - Fatal
    - Followed by a call to os.Exit(1)
  - Panic
    - Followed by a call to panic
- [x] Supports a custom level
- [x] Logging with extra content besides subject and error
- [x] Logging with a customized output format


## Requirements

---
There two requirements to send logs using a Discord Bot
- A token from a Discord bot
- The channel ID to send the logs
---
There is a requirement to send logs using a Discord webhook
- Webhook URL
  -  The URL must have the following pattern https://discord.com/api/webhooks/{webhook.id}/{webhook.token}
---

## Usage

Add the library using the following command

```shell
go get github.com/Lucas-Palomo/go-discord-logger
```

There is a basic example using a discord bot
```go
package main

import (
	"github.com/Lucas-Palomo/go-discord-logger/pkg/bot"
	"github.com/Lucas-Palomo/go-discord-logger/pkg/message"
	"os"
)

func main() {
	// Create a logger
	logger := bot.NewLogger(
		bot.NewBot(
			"YOUR DISCORD BOT TOKEN",
			"THE CHANNEL ID",
		),
	)

	_, err := os.Open("/dev/x")

	if err != nil {
		// Each log event uses a message
		// The message is a structure that contains some information for a better log 
		logger.Error(message.NewMessage("Failed to open /dev/x", err))
	}
}

```


There is a basic example using a discord webhook

```go
package main

import (
  "github.com/Lucas-Palomo/go-discord-logger/pkg/message"
  "github.com/Lucas-Palomo/go-discord-logger/pkg/webhook"
  "os"
)

func main() {
  // Create a logger
  logger := webhook.NewLogger(
    webhook.NewWebhook("https://discord.com/api/webhooks/{webhook.id}/{webhook.token}"),
  )

  _, err := os.Open("/dev/x")

  if err != nil {
    // Each log event uses a message
    // The message is a structure that contains some information for a better log 
    logger.Error(message.NewMessage("Failed to open /dev/x", err))
  }
}

```
