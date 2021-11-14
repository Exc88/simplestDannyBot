package main

import (
	"github.com/Exc88/simplestDannyBot.git/pkg"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1953280162:AAFJvO1IzF4X4HKNhFU7A8wyVrk8TnCkMu0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start();err != nil{
		log.Fatal(err)
	}
}
