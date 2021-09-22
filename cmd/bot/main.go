package main

import (
	"fmt"
	"time"

	"github.com/FerestGo/tg-balancer/pkg/config"
	"github.com/FerestGo/tg-balancer/pkg/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	start := time.Now()
	duration := time.Since(start)

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
	fmt.Printf("Bot started %s %s \n", bot.Self.UserName, duration)

	_, err = bot.RemoveWebhook()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if r.CheckRegexp(`^t\..*\S$`, update.Message.Text) == false {
			fmt.Printf("%d User: %d [%s] %s %s - '%s'", update.Message.MessageID, update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName, update.Message.From.LastName, update.Message.Text)
		} else {
			fmt.Printf("%d User: %d [%s] %s %s - '%s'", update.Message.MessageID, update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName, update.Message.From.LastName, "Secret token")
			deleteMessageConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
			}
			_, err := bot.DeleteMessage(deleteMessageConfig)
			if err != nil {
				fmt.Errorf("Delete secret token error: %s", err)
			}
		}
		start = time.Now()
		duration = time.Since(start)
		reply := r.Handle(update.Message.Text, update.Message.From.ID)
		fmt.Printf(" | %s \n", duration)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)

	}
}
