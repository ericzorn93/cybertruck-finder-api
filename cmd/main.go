package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/ericzorn93/cybertruck-finder-api/internal/logger"
	"github.com/ericzorn93/cybertruck-finder-api/internal/utils"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	godotenv.Load()

	logger.Logger.Debug(
		"Twilio Info",
		slog.String("account_sid", os.Getenv("TWILIO_ACCOUNT_SID")),
		slog.String("auth_token", os.Getenv("TWILIO_AUTH_TOKEN")),
		slog.String("phone_number", os.Getenv("TWILIO_PHONE_NUMBER")),
	)

	client := twilio.NewRestClient()

	http.HandleFunc("GET /cybertrucks", func(w http.ResponseWriter, r *http.Request) {
		res, err := utils.FindCyberTrucks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := &api.CreateMessageParams{}
		params.SetBody("Test message from Cybertruck Finder API")
		params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
		params.SetTo(os.Getenv("TWILIO_TO_PHONE_NUMBER"))
		_, err = client.Api.CreateMessage(params)
		if err != nil {
			logger.Logger.Error("Failed to send message: ", slog.Any("error", err))
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}
		logger.Logger.Info("Message sent successfully")

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

		// Send a message using Twilio
	})

	// Start the server
	logger.Logger.Info("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
