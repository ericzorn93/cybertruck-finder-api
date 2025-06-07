package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ericzorn93/cybertruck-finder-api/internal/logger"
	"github.com/ericzorn93/cybertruck-finder-api/internal/models"
)

// FindCyberTrucks fetches Tesla Cybertruck inventory data from the Tesla API.
// It uses a custom HTTP client that forces HTTP/1.1 to avoid issues with HTTP/2.
func FindCyberTrucks() (models.CyberTruckInventoryResponse, error) {
	// The URL for the Tesla inventory API.
	url := "https://www.tesla.com/inventory/api/v4/inventory-results?query=%7B%22query%22%3A%7B%22model%22%3A%22ct%22%2C%22condition%22%3A%22new%22%2C%22options%22%3A%7B%22TRIM%22%3A%5B%22CTAWD%22%5D%2C%22WHEELS%22%3A%5B%22CYBER_WHEELS%22%5D%2C%22INTERIOR%22%3A%5B%22GREY%22%5D%7D%2C%22arrangeby%22%3Anull%2C%22order%22%3Anull%2C%22market%22%3A%22US%22%2C%22language%22%3A%22en%22%2C%22super_region%22%3A%22north%20america%22%2C%22PaymentType%22%3A%22lease%22%2C%22paymentRange%22%3A1800%2C%22lng%22%3A-73.99475629999999%2C%22lat%22%3A41.0229118%2C%22zip%22%3A%2207675%22%2C%22range%22%3A100%2C%22region%22%3A%22NJ%22%7D%2C%22offset%22%3A0%2C%22count%22%3A24%2C%22outsideOffset%22%3A0%2C%22outsideSearch%22%3Afalse%2C%22isFalconDeliverySelectionEnabled%22%3Atrue%2C%22version%22%3A%22v2%22%7D"

	// Create a custom HTTP transport.
	// Setting ForceAttemptHTTP2 to false explicitly disables HTTP/2 for this transport.
	// This applies to both HTTP and HTTPS connections.
	tr := &http.Transport{
		ForceAttemptHTTP2: false,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Set to true only for testing/debugging
		},
	}

	// Create a new GET request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Logger.Error("Error creating request", slog.Any("error", err))
		return models.CyberTruckInventoryResponse{}, err
	}

	// Add standard headers to make the request appear more like a browser's.
	req.Header.Set("User-Agent", "PostmanRuntime/7.44.0")  // Set a common User-Agent.
	req.Header.Set("Accept", "application/json")           // Request JSON content specifically.
	req.Header.Set("Accept", "*/*")                        // Accept all content types.
	req.Header.Set("Referer", "https://www.tesla.com/")    // Set a referer to mimic browser behavior.
	req.Header.Set("Connection", "keep-alive")             // Use keep-alive to maintain the connection.
	req.Header.Set("Cache-Control", "no-cache")            // Disable caching to ensure fresh data.
	req.Header.Set("Accept-Encoding", "gzip, deflate, br") // Accept compressed responses.

	logger.Logger.Info("Fetching Tesla Cybertruck inventory data (forcing HTTP/1.1)")

	// Create an HTTP client using our custom transport.
	client := &http.Client{
		Transport: tr,
		Timeout:   20 * time.Second, // Set a reasonable timeout for the request.
	}

	// Execute the request.
	resp, err := client.Do(req)
	if err != nil {
		// Handle network or client-side errors.
		logger.Logger.Error("Error making HTTP request", slog.Any("error", err))
		return models.CyberTruckInventoryResponse{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to prevent resource leaks.

	// Check the HTTP status code from the server.
	if resp.StatusCode != http.StatusOK {
		logger.Logger.Error("Received non-OK status code", slog.Int("statusCode", resp.StatusCode), slog.String("status", resp.Status))
		return models.CyberTruckInventoryResponse{}, fmt.Errorf("received non-OK status code: %d %s", resp.StatusCode, resp.Status)
	}

	// Read the successful response body.
	var cyberTruckInventoryResponse models.CyberTruckInventoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&cyberTruckInventoryResponse); err != nil {
		logger.Logger.Error("Error decoding JSON response", slog.Any("error", err))
		return models.CyberTruckInventoryResponse{}, fmt.Errorf("error decoding JSON response: %v", err)
	}

	// Print the successful response.
	logger.Logger.Info("Successfully received Cybertruck inventory data")

	return cyberTruckInventoryResponse, nil
}
