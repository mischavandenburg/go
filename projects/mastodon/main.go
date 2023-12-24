package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// A struct to hold the response data from Mastodon
type Response struct {
	ID        string `json:"id"`
	URI       string `json:"uri"`
	URL       string `json:"url"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// A function to make a POST request to Mastodon with the given status and access token
func postToMastodon(status, token string) (*Response, error) {
	// The base URL of the Mastodon instance
	baseURL := "https://toot.community"

	// The endpoint for creating a new status
	endpoint := "/api/v1/statuses"

	// The full URL of the request
	url := baseURL + endpoint

	// The payload of the request as a map
	payload := map[string]string{
		"status": status,
	}

	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	// Set the content type and authorization headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Close the response body when the function returns
	defer resp.Body.Close()

	// Read the response body as bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	// Create a new Response struct to hold the response data
	var result Response

	// Unmarshal the response body into the Response struct
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// Return the Response struct and nil error
	return &result, nil
}

func main() {
	// The status to post to Mastodon
	status := "Hello, world! This is posted automatically by a Go program I'm working on."

	// The access token to authenticate with Mastodon stored in an env variable
	token := os.Getenv("MASTODON_TOKEN")

	// Call the postToMastodon function and print the result or error
	result, err := postToMastodon(status, token)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success:", result)
	}
}
