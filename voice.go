package sendchamp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	endpointSendVoice = "/voice/send"

	// Type constants
	VoiceTypeOutgoing = "outgoing" // only this works for voice currently
	// VoiceTypeIncoming = "incoming"
	// VoiceTypeTemplate = "template"
)

var (
	// voice errors
	ErrorVoiceRepeat = errors.New("repeat must be at least 1")
)

type Voice struct {
	client *Client
}

type sendVoicePayload struct {
	CustomerMobileNumber []string `json:"customer_mobile_number"`
	Message              string   `json:"message"`
	Type                 string   `json:"type"`
	Repeat               uint     `json:"repeat"`
}

type sendVoiceResponse struct {
	Status  uint                `json:"status"`
	Code    string                `json:"code"`
	Message string                `json:"message"`
	Data    sendVoiceResponseData `json:"data"`
}

type sendVoiceResponseData struct {
	ID             string `json:"id"`
	PhoneNumber    string `json:"phone_number"`
	Reference      string `json:"reference"`
	Status         string `json:"status"`
	CurrentBalance string `json:"currentBalance"`
}

// Send a voice to message to a customer mobile number.
// Use sendchamp.VoiceTypeOutgoing for voiceType - as it is currently
// the only supported type. repeat argument must be greater than zero.
func (v *Voice) Send(customerMobileNumbers []string, message, voiceType string, repeat uint) (sendVoiceResponse, error) {
	url := fmt.Sprint(v.client.baseURL, endpointSendVoice)
	// validate repeat
	if repeat <= 0 {
		return sendVoiceResponse{}, ErrorVoiceRepeat
	}
	payload := sendVoicePayload{
		CustomerMobileNumber: customerMobileNumbers,
		Message:              message,
		Type:                 voiceType,
		Repeat:               repeat,
	}
	reqData := v.client.NewRequest(http.MethodPost, url)
	resp, err := v.client.SendRequest(reqData, payload)
	if err != nil {
		return sendVoiceResponse{}, err
	}
	r := sendVoiceResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendVoiceResponse{}, err
	}
	return r, nil
}
