package sendchamp

import (
	"bytes"
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
	CustomerMobileNumber string `json:"customer_mobile_number"`
	Message              string `json:"message"`
	Type                 string `json:"type"`
	Repeat               uint   `json:"repeat"`
}

type sendVoiceResponse struct {
	Status  string                `json:"status"`
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
func (v *Voice) Send(customerMobileNumber, message, voiceType string, repeat uint) (sendVoiceResponse, error) {
	url := fmt.Sprint(v.client.baseURL, endpointSendVoice)
	// validate repeat
	if repeat <= 0 {
		return sendVoiceResponse{}, ErrorVoiceRepeat
	}

	byte, err := json.Marshal(sendVoicePayload{
		CustomerMobileNumber: customerMobileNumber,
		Message:              message,
		Type:                 voiceType,
		Repeat:               repeat,
	})
	if err != nil {
		return sendVoiceResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))
	if err != nil {
		return sendVoiceResponse{}, err
	}

	addHeaders(req, v.client)
	res, err := v.client.httpClient.Do(req)
	if err != nil {
		return sendVoiceResponse{}, err
	}
	defer res.Body.Close()

	r := sendVoiceResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return sendVoiceResponse{}, err
	}

	return r, nil
}
