package sendchamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpointSendTemplate = "/whatsapp/message/send"

	// message types
	WhatsappTypeTemplate = "template"
)

type Whatsapp struct {
	client *Client
}

type customData struct {
	Body map[string]string `json:"body"`
}

type sendTemplatePayload struct {
	Recipient    string     `json:"recipient"`
	Sender       string     `json:"sender"`
	Type         string     `json:"type"`
	TemplateCode string     `json:"template_code"`
	CustomData   customData `json:"custom_data"`
}

type sendTemplateResponse struct {
	Status  string                   `json:"status"`
	Code    string                   `json:"code"`
	Message string                   `json:"message"`
	Data    sendTemplateResponseData `json:"data"`
}

type sendTemplateResponseData struct {
	ProviderReference string `json:"provider_reference"`
	ProviderMessage   string `json:"provider_message"`
	ProviderStatus    string `json:"provider_status"`
}

func (w *Whatsapp) SendTemplate(recipient, sender, mType, templateCode string, data map[string]string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendTemplate)

	byte, err := json.Marshal(sendTemplatePayload{
		Recipient:    recipient,
		Sender:       sender,
		Type:         mType,
		TemplateCode: templateCode,
		CustomData: customData{
			Body: data,
		},
	})

	if err != nil {
		return sendTemplateResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))
	if err != nil {
		return sendTemplateResponse{}, err
	}

	addHeaders(req, w.client)
	res, err := w.client.httpClient.Do(req)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	defer res.Body.Close()
	r := sendTemplateResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}
