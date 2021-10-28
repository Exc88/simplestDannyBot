package main

import "C"
import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)


func main() {
	bot, err := tgbotapi.NewBotAPI("1953280162:AAFJvO1IzF4X4HKNhFU7A8wyVrk8TnCkMu0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,"txt"))
		}

		comp := update.Message.Text
		if comp == "ремонт" {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "С каким устройством у вас проблема?\n" +
				"Напишите:\n -Компьютер\n -Ноутбук\n -Телефон\n -Планшет"))
		}
	}

}
