package sendchamp

import "net/http"

const (
	URLLive = "https://api.sendchamp.com/api/v1"
	URLTest = "https://sandbox-api.sendchamp.com/api/v1"
)

func NewClient(publicKey, mode string) *Client {
	baseUrl := URLTest

	if mode == "live" {
		baseUrl = URLLive
	}

	// set the base url depending on the mode
	return &Client{
		baseURL:    baseUrl,
		httpClient: http.DefaultClient,
		publicKey:  publicKey,
		mode:       mode,
	}
}
