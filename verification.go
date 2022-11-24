package sendchamp

import (
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
	Status  uint              `json:"status"`
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
	Status  uint                 `json:"status"`
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
	reqData := v.client.NewRequest(http.MethodPost, url)
	resp, err := v.client.SendRequest(reqData, payload)
	if err != nil {
		return sendOtpResponse{}, nil
	}
	r := sendOtpResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return sendOtpResponse{}, nil
	}
	return r, nil
}

func (v *Verification) ConfirmOTP(code, reference string) (confirmOtpResponse, error) {
	url := fmt.Sprint(v.client.baseURL, endpointConfirmOtp)

	payload := confirmOtpPayload{
		VerificationCode:      code,
		VerificationReference: reference,
	}
	reqData := v.client.NewRequest(http.MethodPost, url)
	resp, err := v.client.SendRequest(reqData, payload)
	if err != nil {
		return confirmOtpResponse{}, err
	}
	r := confirmOtpResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return confirmOtpResponse{}, err
	}

	return r, nil
}
