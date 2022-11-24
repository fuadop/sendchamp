package sendchamp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpointSendWhatsapp = "/whatsapp/message/send"

	// message types
	WhatsappTypeTemplate = "template"
	WhatsappTypeText     = "text"
	WhatsappTypeAudio    = "audio"
	WhatsappTypeVideo    = "video"
	WhatsappTypeLocation = "location"
	WhatsappTypeSticker  = "sticker"
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

type sendTextPayload struct {
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

type sendAudioPayload struct {
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Link      string `json:"link"`
}

type sendVideoPayload struct {
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
	Type      string `json:"type"`
	Link      string `json:"link"`
}

type sendLocationPayload struct {
	Recipient string  `json:"recipient"`
	Sender    string  `json:"sender"`
	Type      string  `json:"type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}
type sendTemplateResponse struct {
	Status  uint                   `json:"status"`
	Code    string                   `json:"code"`
	Message string                   `json:"message"`
	Data    sendTemplateResponseData `json:"data"`
}

type sendTemplateResponseData struct {
	ProviderReference string `json:"provider_reference"`
	ProviderMessage   string `json:"provider_message"`
	ProviderStatus    string `json:"provider_status"`
}

// SendTemplate - Send a whatsapp message using
// template created on dashboard.
func (w *Whatsapp) SendTemplate(sender, recipient, templateCode string, data map[string]string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendWhatsapp)
	payload := sendTemplatePayload{
		Recipient:    recipient,
		Sender:       sender,
		Type:         WhatsappTypeTemplate,
		TemplateCode: templateCode,
		CustomData: customData{
			Body: data,
		},
	}
	reqData := w.client.NewRequest(http.MethodPost, url)
	resp, err := w.client.SendRequest(reqData, payload)
	if err != nil {
		return sendTemplateResponse{}, err
	}
	r := sendTemplateResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}

// SendText - Send a whatsapp text.
func (w *Whatsapp) SendText(sender, recipient, message string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendWhatsapp)
	payload := sendTextPayload{
		Recipient: recipient,
		Sender:    sender,
		Type:      WhatsappTypeText,
		Message:   message,
	}
	reqData := w.client.NewRequest(http.MethodPost, url)
	resp, err := w.client.SendRequest(reqData, payload)
	if err != nil {
		return sendTemplateResponse{}, err
	}
	r := sendTemplateResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}

// SendAudio - Send a whatsapp audio message.
func (w *Whatsapp) SendAudio(sender, recipient, message, link string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendWhatsapp)
	payload := sendAudioPayload{
		Recipient: recipient,
		Sender:    sender,
		Type:      WhatsappTypeAudio,
		Message:   message,
		Link:      link,
	}
	reqData := w.client.NewRequest(http.MethodPost, url)
	resp, err := w.client.SendRequest(reqData, payload)
	if err != nil {
		return sendTemplateResponse{}, err
	}
	r := sendTemplateResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}

// SendVideo - Send a whatsapp video message.
func (w *Whatsapp) SendVideo(sender, recipient, link string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendWhatsapp)
	payload := sendVideoPayload{
		Recipient: recipient,
		Sender:    sender,
		Type:      WhatsappTypeVideo,
		Link:      link,
	}
	reqData := w.client.NewRequest(http.MethodPost, url)
	resp, err := w.client.SendRequest(reqData, payload)
	if err != nil {
		return sendTemplateResponse{}, err
	}
	r := sendTemplateResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}

// SendSticker - Send a whatsapp sticker message.
func (w *Whatsapp) SendSticker(sender, recipient, link string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendWhatsapp)
	payload := sendVideoPayload{
		Recipient: recipient,
		Sender:    sender,
		Type:      WhatsappTypeSticker,
		Link:      link,
	}
	reqData := w.client.NewRequest(http.MethodPost, url)
	resp, err := w.client.SendRequest(reqData, payload)
	if err != nil {
		return sendTemplateResponse{}, err
	}
	r := sendTemplateResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}

// SendLocation - Send a location via whatsapp.
func (w *Whatsapp) SendLocation(sender, recipient string, longitude, latitude float64, name, address string) (sendTemplateResponse, error) {
	url := fmt.Sprint(w.client.baseURL, endpointSendWhatsapp)
	payload := sendLocationPayload{
		Recipient: recipient,
		Sender:    sender,
		Type:      WhatsappTypeLocation,
		Latitude:  latitude,
		Longitude: longitude,
		Name:      name,
		Address:   address,
	}
	reqData := w.client.NewRequest(http.MethodPost, url)
	resp, err := w.client.SendRequest(reqData, payload)
	if err != nil {
		return sendTemplateResponse{}, err
	}
	r := sendTemplateResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendTemplateResponse{}, err
	}

	return r, nil
}
