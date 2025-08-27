package exoplanets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func QueryExoplanets(query string) (any, error) {
	form := url.Values{}
	form.Set("query", query)
	form.Set("format", "json")

	url := os.Getenv("EXOPLANET_TAP_SYNC_URL")

	resp, err := http.Post(
		url,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(form.Encode()),
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to post query: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("TAP error: %v", string(body))
	}

	var tr any
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, fmt.Errorf("failed to decode TAP response: %w", err)
	}

	return &tr, nil

}
