package sendchamp_test

import (
	"os"
	"testing"

	"github.com/fuadop/sendchamp"
)

func TestSendSms(t *testing.T) {
	key, ok := os.LookupEnv("PUBLIC_KEY")
	if !ok {
		key = "PUBLIC_KEY"
	}

	client := sendchamp.NewClient(key, sendchamp.ModeLive)
	res, err := client.NewSms().Send("ASF", []string{"2348153207998"}, "TestSendSingleSms", sendchamp.RouteNonDND)
	if err != nil {
		t.Fail()
	}

	if res.Status != "success" {
		t.Fail()
	}
}
