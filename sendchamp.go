package sendchamp

import "net/http"

const (
	URLLive           = "https://api.sendchamp.com/api/v1"
	URLTest           = "https://sandbox-api.sendchamp.com/api/v1"
	URLLocalSimulator = "http://localhost:2920/api/v1"

	ModeLive           = "live"
	ModeTest           = "test"
	ModeLocalSimulator = "local-simulator"
)

func NewClient(publicKey, mode string) *Client {
	baseUrl := URLTest

	if mode == ModeLive {
		baseUrl = URLLive
	}

	if mode == ModeLocalSimulator {
		baseUrl = URLLocalSimulator
	}

	// set the base url depending on the mode
	return &Client{
		baseURL:    baseUrl,
		httpClient: http.DefaultClient,
		publicKey:  publicKey,
		mode:       mode,
	}
}
