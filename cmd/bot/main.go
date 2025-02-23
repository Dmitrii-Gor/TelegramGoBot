package main

import (
	api "github.com/OvyFlash/telegram-bot-api"
	"log"
	"time"
)

func main() {
	bot, err := api.NewBotAPI("BOT_TOKEN")
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := api.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChannel := bot.GetUpdatesChan(updateConfig)

	// Optional: Clear initial updates
	time.Sleep(time.Millisecond * 500)
	updatesChannel.Clear()

	for update := range updatesChannel {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyParameters.MessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
