package sendchamp

import (
	"net/http"
	"os"
)

const (
	URLLive           = "https://api.sendchamp.com/api/v1"
	URLTest           = "https://sandbox-api.sendchamp.com/api/v1"
	URLLocalSimulator = "http://localhost:2920/api/v1"

	ModeLive           = "live"
	ModeTest           = "test"
	ModeLocalSimulator = "local-simulator"
)

type Keys struct {
	PublicKey string
}

func NewClient(key *Keys, mode string) *Client {
	var publicKey string = os.Getenv("SENDCHAMP_PUBLIC_KEY")
	if key != nil {
		publicKey = key.PublicKey
	}

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
