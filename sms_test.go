package sendchamp_test

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/fuadop/sendchamp"
)

var key = &sendchamp.Keys{
	PublicKey: os.Getenv("PUBLIC_KEY"),
}

// todo: switch to test mode when it's back available
var client sendchamp.Client = *sendchamp.NewClient(key, sendchamp.ModeLive)

func TestSendSms(t *testing.T) {
	res, err := client.NewSms().Send("ASF", []string{"2348153207998"}, "TestSendSingleSms", sendchamp.RouteNonDND)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestCreateSenderID(t *testing.T) {
	id := rand.Int() / 1e12
	res, err := client.NewSms().CreateSenderID(strconv.Itoa(id), "Your test from golang passed", sendchamp.UseCaseMarketing)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestGetDeliveryReport(t *testing.T) {
	// send an sms then pass uid to deliveryreport func
	sms := client.NewSms()
	res, err := sms.Send("ASF", []string{"2348153207998"}, "test_delivery_report", sendchamp.RouteNonDND)
	if err != nil {
		t.Error(err)
	}

	if res.Data.ID != nil {
		// assert to string
		resp, err := sms.GetDeliveryReport(res.Data.ID.(string))
		if err != nil {
			t.Error(err)
		}
		if resp.Status != 200 {
			t.Error("res.Status: ", resp.Status)
		}

		if res.Data.ID.(string) != resp.Data.ID {
			t.Error("Message IDs don't match")
		}
	}
}

// todo: This test fails (-)
func TestGetBulkDeliveryReport(t *testing.T) {
	// send an sms then pass uid to deliveryreport func
	sms := client.NewSms()
	res, err := sms.Send("ASF", []string{"2348153207998", "2349158808093"}, "test_delivery_report", sendchamp.RouteNonDND)
	if err != nil {
		t.Error(err)
	}

	resp, err := sms.GetBulkDeliveryReport(res.Data.UID)
	if err != nil {
		t.Error(err)
	}

	if resp.Status != 200 {
		t.Error("res.Status: ", resp.Status)
	}

	if res.Data.UID != resp.Data.UID {
		t.Error("Message IDs don't match")
	}
}
