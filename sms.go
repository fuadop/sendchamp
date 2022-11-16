package sendchamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpointSendSms            = "/sms/send"
	endpointSenderID           = "/sms/create-sender-id"
	endpointDeliveryReport     = "/sms/status/"
	endpointBulkDeliveryReport = "/sms/bulk-sms-status/"

	// sms routes
	RouteDND           = "dnd"
	RouteNonDND        = "non_dnd"
	RouteInternational = "international"

	// sender ID use cases
	UseCaseTransactional             = "transactional"
	UseCaseMarketing                 = "marketing"
	UseCaseTransactionalAndMarketing = "transaction_marketing"
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

type createSenderIdPayload struct {
	Name    string `json:"name"`
	Sample  string `json:"sample"`
	UseCase string `json:"use_case"`
}

type sendSmsResponse struct {
	Status  uint                `json:"status"`
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Data    sendSmsResponseData `json:"data"`
}

type sendSmsResponseData struct {
	// ID can be string or int: use type assertion for required case ID.(string)
	ID            interface{} `json:"id"`
	PhoneNumber   string      `json:"phone_number"`
	Reference     string      `json:"reference"`
	Amount        string      `json:"amount"`
	ServiceCharge string      `json:"service_charge"`
	Status        string      `json:"status"`

	// it is possibly null
	DeliveredAt string `json:"delivered_at"`
	TotalSms    int    `json:"total_sms"`
	BusinessUID string `json:"business_uid"`
	UID         string `json:"uid"`

	// timestamps
	SentAt    string `json:"sent_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type createSenderIDResponse struct {
	Status  string                     `json:"status"`
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Data    createSenderIDResponseData `json:"data"`
}

type createSenderIDResponseData struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	BusinessUID string `json:"business_uid"`
}

type GetDeliveryReport struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Send sms to one or many phone number
func (s *Sms) Send(senderName string, to []string, message, route string) (sendSmsResponse, error) {
	url := fmt.Sprint(s.client.baseURL, endpointSendSms)

	// populate request body
	byte, err := json.Marshal(sendSmsPayload{
		SenderName: senderName,
		To:         to,
		Message:    message,
		Route:      route,
	})

	if err != nil {
		return sendSmsResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))

	if err != nil {
		return sendSmsResponse{}, err
	}

	// add necessary request headers
	addHeaders(req, s.client)

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return sendSmsResponse{}, err
	}

	defer res.Body.Close()

	r := sendSmsResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return sendSmsResponse{}, err
	}

	return r, nil
}

// Create a sender ID
func (s *Sms) CreateSenderID(name, sample, useCase string) (createSenderIDResponse, error) {
	url := fmt.Sprint(s.client.baseURL, endpointSenderID)

	byte, err := json.Marshal(createSenderIdPayload{
		Name:    name,
		Sample:  sample,
		UseCase: useCase,
	})
	if err != nil {
		return createSenderIDResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))
	if err != nil {
		return createSenderIDResponse{}, err
	}

	addHeaders(req, s.client)

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return createSenderIDResponse{}, err
	}

	defer res.Body.Close()
	r := createSenderIDResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return createSenderIDResponse{}, err
	}

	return r, nil
}

// Get single sms delivery report
// This works for sms sent to a single number
// only.
// smsID = res.Data.ID from sms.Send method
func (s *Sms) GetDeliveryReport(smsID string) (sendSmsResponse, error) {
	url := fmt.Sprint(s.client.baseURL, endpointDeliveryReport, smsID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return sendSmsResponse{}, err
	}

	addHeaders(req, s.client)
	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return sendSmsResponse{}, err
	}

	defer res.Body.Close()
	r := sendSmsResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return sendSmsResponse{}, err
	}

	return r, nil
}

func (s *Sms) GetBulkDeliveryReport(bulkSmsUID string) (sendSmsResponse, error) {
	url := fmt.Sprint(s.client.baseURL, endpointDeliveryReport, bulkSmsUID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return sendSmsResponse{}, err
	}

	addHeaders(req, s.client)
	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return sendSmsResponse{}, err
	}

	defer res.Body.Close()
	r := sendSmsResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return sendSmsResponse{}, err
	}

	return r, nil
}
