package sendchamp_test

import (
	"testing"

	"github.com/fuadop/sendchamp"
)

func TestSendVoice(t *testing.T) {
	customerMobileNumbers := []string{"2348153207998"}
	message := "Test from golang test suite."
	voiceType := sendchamp.VoiceTypeOutgoing // only supported type currently
	var repeat uint = 3                      // repeat the voice 3 times

	// expect this to pass
	res, err := client.NewVoice().Send(customerMobileNumbers, message, voiceType, repeat)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}

	_, err = client.NewVoice().Send(customerMobileNumbers, message, voiceType, 0)
	// expect this to fail due to validation error
	if err != sendchamp.ErrorVoiceRepeat {
		t.Errorf("Expected: %s, but got: %s", sendchamp.ErrorVoiceRepeat.Error(), err)
	}
}
