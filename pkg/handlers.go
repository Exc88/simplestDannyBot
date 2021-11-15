package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Idk this command :(")

	switch message.Command() {
	case commandStart:
		msg.Text = "You enabled to start"
		_,err := b.bot.Send(msg)
		return err
	default:
		_,err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message)  {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
