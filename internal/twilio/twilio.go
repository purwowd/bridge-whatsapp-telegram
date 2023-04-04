package twilio

import (
	"github.com/purwowd/bridge-whatsapp-telegram/internal/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendWhatsAppMessage(cfg *config.Config, recipient, message string) error {
	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(recipient)
	params.SetFrom(cfg.TwilioWhatsAppNumber)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}

func ReceiveWhatsAppMessage(cfg *config.Config, fromNumber string, webhookCallback func(string, string)) {
}
