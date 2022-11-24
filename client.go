package sendchamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

// return a new whatsapp instance
func (c *Client) NewWhatsapp() *Whatsapp {
	return &Whatsapp{
		client: c,
	}
}

type Request struct {
	Method string
	URL    string
}

func (c *Client) NewRequest(method, url string) *Request {
	return &Request{
		Method: method,
		URL: url,
	}
}

func (c *Client) SendRequest(reqData *Request, payload interface{}) (b []byte, err error) {
	var (
		req  *http.Request
		body *bytes.Buffer
	)
	if payload != nil {
		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBytes)
		req, err = http.NewRequest(reqData.Method, reqData.URL, body)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(reqData.Method, reqData.URL, nil)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, c.publicKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return b, nil
}
