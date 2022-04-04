package sendchamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// otp channels
	OTPChannelSMS   = "sms"
	OTPChannelEmail = "email"

	// otp token types
	OTPTokenTypeNumeric      = "numeric"
	OTPTokenTypeAlphaNumeric = "alphanumeric"

	endpointSendOtp    = "/verification/create"
	endpointConfirmOtp = "/verification/confirm"
)

type Verification struct {
	client *Client
}

type SendOTPPayload struct {
	Channel              string      `json:"channel"`
	Sender               string      `json:"sender"`
	TokenType            string      `json:"token_type"`
	TokenLength          string      `json:"token_length"`
	ExpirationTime       int         `json:"expiration_time"`
	CustomerMobileNumber string      `json:"customer_mobile_number"`
	CustomerEmailAddress string      `json:"customer_email_address"`
	MetaData             interface{} `json:"meta_data"`
}

type confirmOtpPayload struct {
	VerificationCode      string `json:"verification_code"`
	VerificationReference string `json:"verification_reference"`
}

type sendOtpResponse struct {
	Status  string              `json:"status"`
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Data    sendOtpResponseData `json:"data"`
}

type sendOtpResponseData struct {
	BusinessUID string                     `json:"business_uid"`
	Reference   string                     `json:"reference"`
	Channel     sendOtpResponseDataChannel `json:"channel"`
	Status      string                     `json:"status"`
}

type sendOtpResponseDataChannel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type confirmOtpResponse struct {
	Status  string                 `json:"status"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    confirmOtpResponseData `json:"data"`
}

type confirmOtpResponseData struct {
	ID                 int                        `json:"id"`
	BusinessUID        string                     `json:"business_uid"`
	TransactionUID     string                     `json:"transaction_uid"`
	ChannelID          int                        `json:"channel_id"`
	Token              string                     `json:"token"`
	TokenType          string                     `json:"token_type"`
	TokenLength        string                     `json:"token_length"`
	TokenDuration      string                     `json:"token_duration"`
	ConfirmationCharge interface{}                `json:"confirmation_charge"`
	Status             string                     `json:"status"`
	Phone              string                     `json:"phone"`
	Email              string                     `json:"email"`
	VerifiedAt         string                     `json:"verified_at"`
	CreatedAt          string                     `json:"created_at"`
	UpdatedAt          string                     `json:"updated_at"`
	DeletedAt          string                     `json:"deleted_at"`
	Channel            sendOtpResponseDataChannel `json:"channel"`
}

// SendOTP - Send otp to a mobile number
func (v *Verification) SendOTP(payload SendOTPPayload) (sendOtpResponse, error) {
	url := fmt.Sprint(v.client.baseURL, endpointSendOtp)
	byte, err := json.Marshal(payload)
	if err != nil {
		return sendOtpResponse{}, nil
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))
	if err != nil {
		return sendOtpResponse{}, nil
	}

	addHeaders(req, v.client)
	res, err := v.client.httpClient.Do(req)
	if err != nil {
		return sendOtpResponse{}, nil
	}

	defer res.Body.Close()
	r := sendOtpResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return sendOtpResponse{}, nil
	}

	return r, nil
}

func (v *Verification) ConfirmOTP(code, reference string) (confirmOtpResponse, error) {
	url := fmt.Sprint(v.client.baseURL, endpointConfirmOtp)

	byte, err := json.Marshal(confirmOtpPayload{
		VerificationCode:      code,
		VerificationReference: reference,
	})
	if err != nil {
		return confirmOtpResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byte))
	if err != nil {
		return confirmOtpResponse{}, err
	}

	addHeaders(req, v.client)
	res, err := v.client.httpClient.Do(req)
	if err != nil {
		return confirmOtpResponse{}, err
	}

	defer res.Body.Close()
	r := confirmOtpResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return confirmOtpResponse{}, err
	}

	return r, nil
}
