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
