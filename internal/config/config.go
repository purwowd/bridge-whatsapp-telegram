package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken        string
	TelegramChatID       int64
	TwilioAccountSID     string
	TwilioAuthToken      string
	TwilioWhatsAppNumber string
	ListenAddress        string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		TelegramToken:        viper.GetString("telegram.token"),
		TelegramChatID:       viper.GetInt64("telegram.chat_id"),
		TwilioAccountSID:     viper.GetString("twilio.account_sid"),
		TwilioAuthToken:      viper.GetString("twilio.auth_token"),
		TwilioWhatsAppNumber: viper.GetString("twilio.whatsapp_number"),
		ListenAddress:        viper.GetString("listen_address"),
	}

	return cfg, nil
}
