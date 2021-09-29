package main

import (
	"fmt"
	"time"

	"github.com/FerestGo/tg-balancer/pkg/balancer"
	"github.com/FerestGo/tg-balancer/pkg/config"
	"github.com/FerestGo/tg-balancer/pkg/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const TOKEN_PATTERN = `^t\..*\S$`

func main() {
	start := time.Now()
	duration := time.Since(start)

	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("Config init error: %s", err.Error())
	}

	var r router.Router
	r.Get()
	balancer.InitExternal()
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		fmt.Printf("Telegram bot error: %s", err.Error())
	}
	fmt.Printf("Bot started %s %s \n", bot.Self.UserName, duration)

	_, err = bot.RemoveWebhook()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5
	updates, err := bot.GetUpdatesChan(u)
	reply := ""

	for update := range updates {
		if update.Message == nil {
			continue
		}
		start = time.Now()
		duration = time.Since(start)
		if r.CheckRegexp(TOKEN_PATTERN, update.Message.Text) == false {
			fmt.Printf("%d User: %d [%s] %s %s - '%s'", update.Message.MessageID, update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName, update.Message.From.LastName, update.Message.Text)
			reply = r.Handle(update.Message.Text, update.Message.From.ID)
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
			reply = balancer.InitAnalysis(update.Message.Text, update.Message.From.ID)
		}
		fmt.Printf(" | %s \n", duration)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ParseMode = "Markdown"
		msg.DisableWebPagePreview = true
		bot.Send(msg)

	}
}
