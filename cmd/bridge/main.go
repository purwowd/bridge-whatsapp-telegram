package main

import (
	"github.com/purwowd/bridge-whatsapp-telegram/internal/api"
	"github.com/purwowd/bridge-whatsapp-telegram/internal/config"
	"github.com/purwowd/bridge-whatsapp-telegram/internal/telegram"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	go telegram.StartTelegramBot(cfg)

	http.HandleFunc("/webhook", api.WebhookHandler(cfg))
	log.Fatal(http.ListenAndServe(cfg.ListenAddress, nil))
}
