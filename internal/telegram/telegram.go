package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/purwowd/bridge-whatsapp-telegram/internal/config"
	"github.com/purwowd/bridge-whatsapp-telegram/internal/twilio"
	"log"
)

func StartTelegramBot(cfg *config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			whatsappNumber := fmt.Sprintf("whatsapp:%s", cfg.TwilioWhatsAppNumber)
			twilio.SendWhatsAppMessage(cfg, whatsappNumber, update.Message.Text)
		}
	}
}

func SendToTelegram(cfg *config.Config, fromNumber, message string) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Println("Error creating Telegram bot:", err)
		return
	}

	msg := tgbotapi.NewMessage(cfg.TelegramChatID, fmt.Sprintf("From: %s\nMessage: %s", fromNumber, message))
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("Error sending message to Telegram:", err)
	}
}
