package main

import (
	"fmt"
	"github.com/FerestGo/tg-balancer/pkg/config"
	"github.com/FerestGo/tg-balancer/pkg/router"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("Config init error: %s", err.Error())
	}

	var r router.Router
	r.Get()
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		fmt.Printf("Telegram bot error: %s", err.Error())
	}
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.RemoveWebhook()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Printf("%d User: %d [%s] %s %s - '%s' \n", update.Message.MessageID, update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName, update.Message.From.LastName, update.Message.Text)
		reply := r.Handle(update.Message.Text, update.Message.From.ID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
		// fmt.Printf("%\n", message.MessageID)
		// message.Text = "3"
		// bot.Send(message)

	}
}
