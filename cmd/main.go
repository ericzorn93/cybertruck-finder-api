package main

import (
	"encoding/json"
	"net/http"

	"github.com/ericzorn93/cybertruck-finder-api/internal/logger"
	"github.com/ericzorn93/cybertruck-finder-api/internal/utils"
)

func main() {
	http.HandleFunc("GET /cybertrucks", func(w http.ResponseWriter, r *http.Request) {
		res, err := utils.FindCyberTrucks()
		if err != nil {
			http.Error(w, "Failed to fetch Cybertruck data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	logger.Logger.Info("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
