package sendchamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpointSendSms = "/sms/send"

	// sendchamp routes
	RouteDND           = "dnd"
	RouteNonDND        = "non_dnd"
	RouteInternational = "international"
)

type Sms struct {
	client *Client
}

type sendSmsPayload struct {
	SenderName string   `json:"sender_name"`
	To         []string `json:"to"`
	Message    string   `json:"message"`
	Route      string   `json:"route"`
}

type SendSmsResponse struct {
	Status  string              `json:"status"`
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Data    SendSmsResponseData `json:"data"`
}

type SendSmsResponseData struct {
	Id            string `json:"id"`
	PhoneNumber   string `json:"phone_number"`
	Reference     string `json:"reference"`
	Amount        string `json:"amount"`
	ServiceCharge string `json:"service_charge"`
	Status        string `json:"status"`

	// it is possibly null
	DeliveredAt string `json:"delivered_at"`
	TotalSms    int    `json:"total_sms"`
	BusinessUid string `json:"business_uid"`
	Uid         string `json:"uid"`

	// timestamps
	SentAt    string `json:"sent_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Send sms to a phone number
func (s *Sms) Send(senderName string, to []string, message, route string) (SendSmsResponse, error) {
	url := fmt.Sprint(s.client.baseURL, endpointSendSms)

	// populate request body
	byte, err := json.Marshal(sendSmsPayload{
		SenderName: senderName,
		To:         to,
		Message:    message,
		Route:      route,
	})

	if err != nil {
		return SendSmsResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))

	if err != nil {
		return SendSmsResponse{}, err
	}

	// add necessary request headers
	addHeaders(req, s.client)

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return SendSmsResponse{}, err
	}

	defer res.Body.Close()

	r := SendSmsResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return SendSmsResponse{}, err
	}

	return r, nil
}
