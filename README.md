# Discord Logger

This project is a simple logger that uses a Discord bot to send logs to a channel

## Features

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

There two requirements:
- A token from a Discord bot
- The channel ID to send the logs

## Usage

Add the library using the following command

```shell
go get github.com/Lucas-Palomo/go-discord-logger
```

There is a basic example
```go
package main

import (
	discordlogger "github.com/Lucas-Palomo/go-discord-logger/pkg"
	"os"
)

func main() {
	// Create a logger
	logger := discordlogger.NewLogger(
		"YOUR DISCORD BOT TOKEN",
		"THE CHANNEL ID",
	)
	
	_, err := os.Open("/dev/x")
	
	if err != nil {
		// Each log event uses a message
		// The message is a structure that contains some information for a better log 
		logger.Error(discordlogger.NewMessage("Failed to open /dev/x", err))
	}
}

```

