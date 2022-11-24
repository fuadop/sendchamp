package sendchamp_test

import (
	"testing"

	"github.com/fuadop/sendchamp"
)

type metadata struct {
	FirstName string // important - export fields and add json annotations
	LastName  string
}

func TestSendOTP(t *testing.T) {
	// test sms otp
	payload := sendchamp.SendOTPPayload{
		Channel:              sendchamp.OTPChannelSMS,
		Sender:               "ASF",
		TokenType:            sendchamp.OTPTokenTypeNumeric,
		TokenLength:          "4",
		ExpirationTime:       6,
		CustomerMobileNumber: "2348153207998",
		CustomerEmailAddress: "fuadolatunji@gmail.com",
		MetaData:             metadata{"Fuad", "Olatunji"},
	}

	res, err := client.NewVerification().SendOTP(payload)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Errorf("res.Status: %v", res.Status)
	}

	payload = sendchamp.SendOTPPayload{
		Channel:              sendchamp.OTPChannelEmail,
		Sender:               "ASF",
		TokenType:            sendchamp.OTPTokenTypeAlphaNumeric,
		TokenLength:          "4",
		ExpirationTime:       6,
		CustomerMobileNumber: "2348153207998",
		CustomerEmailAddress: "fuadolatunji@gmail.com",
		MetaData:             metadata{"Fuad", "Olatunji"},
	}

	res, err = client.NewVerification().SendOTP(payload)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Errorf("res.Status: %v", res.Status)
	}
}

func TestConfirmOTP(t *testing.T) {
	code, reference := "01799", "de858be1-6240-48fb-916c-4d07d8c9f79d"
	res, err := client.NewVerification().ConfirmOTP(code, reference)
	if err != nil {
		t.Error(err)
	}
	if res.Status != 200 {
		t.Errorf("res.Status: %v", res.Status)
	}
}
