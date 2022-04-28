package sendchamp_test

import (
	"fmt"
	"testing"
)

func TestSendTemplate(t *testing.T) {
	recipient := "2348153207998"
	sender := "2348120678278"
	mType := "template"
	templateCode := "912671fe-5f20-4b59-92ee-a33a62ea6a19"
	data := map[string]string{
		"1": "Test",
		"2": "1234",
		"3": "10",
	}

	res, err := client.NewWhatsapp().SendTemplate(recipient, sender, mType, templateCode, data)
	if err != nil {
		t.Error(err)
	}

	if res.Status != "success" {
		t.Error("res.Status: ", res.Status)
	}

	fmt.Printf("%+v", res)
}
