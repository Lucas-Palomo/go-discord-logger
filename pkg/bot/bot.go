package bot

import (
	"github.com/Lucas-Palomo/go-discord-logger/internal"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	internal.Dispatcher
	token   string
	channel string
	session *discordgo.Session
}

func NewBot(token string, channel string) *Bot {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalln("Failed to create a discord session,", err)
	}
	return &Bot{
		token:   token,
		channel: channel,
		session: session,
	}
}

func (bot *Bot) Send(content string) {
	_, err := bot.session.ChannelMessageSend(bot.channel, content)
	if err != nil {
		log.Fatalln("Failed to send message", err)
	}
}
