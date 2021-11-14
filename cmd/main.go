package main

import (
	"github.com/Exc88/simplestDannyBot.git/pkg"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1953280162:AAEWa1cqSHlpHCTi-GEu7CjGyrIt2k7jAAs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start();err != nil{
		log.Fatal(err)
	}
}
