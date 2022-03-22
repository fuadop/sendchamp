package sendchamp_test

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/fuadop/sendchamp"
)

var key string = os.Getenv("PUBLIC_KEY")

// todo: switch to test mode when it's back available
var client sendchamp.Client = *sendchamp.NewClient(key, sendchamp.ModeLive)

func TestSendSms(t *testing.T) {
	res, err := client.NewSms().Send("ASF", []string{"2348153207998"}, "TestSendSingleSms", sendchamp.RouteNonDND)
	if err != nil {
		t.Fail()
	}

	if res.Status != "success" {
		t.Fail()
	}
}

func TestCreateSenderID(t *testing.T) {
	id := rand.Int() / 1e12
	res, err := client.NewSms().CreateSenderID(strconv.Itoa(id), "Your test from golang passed", sendchamp.UseCaseMarketing)
	if err != nil {
		t.Fail()
	}

	if res.Status != "success" {
		t.Fail()
	}
}
