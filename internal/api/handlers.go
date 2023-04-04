package api

import (
	"encoding/json"
	"github.com/purwowd/bridge-whatsapp-telegram/internal/config"
	"github.com/purwowd/bridge-whatsapp-telegram/internal/telegram"
	"io/ioutil"
	"net/http"
)

func WebhookHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}

		var payload map[string]interface{}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fromNumber := payload["From"].(string)
		message := payload["Body"].(string)

		telegram.SendToTelegram(cfg, fromNumber, message)

		w.WriteHeader(http.StatusOK)
	}
}
