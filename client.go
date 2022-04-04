package sendchamp

import "net/http"

type Client struct {
	baseURL    string
	httpClient *http.Client
	publicKey  string
	mode       string
}

// return a new sms service instance
func (c *Client) NewSms() *Sms {
	return &Sms{
		client: c,
	}
}

// return a new voice instance
func (c *Client) NewVoice() *Voice {
	return &Voice{
		client: c,
	}
}

// return a new verification instance
func (c *Client) NewVerification() *Verification {
	return &Verification{
		client: c,
	}
}
