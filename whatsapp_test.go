package sendchamp_test

import (
	"testing"
)

func TestSendTemplate(t *testing.T) {
	sender := "2348120678278"
	recipient := "2348153207998"
	templateCode := "912671fe-5f20-4b59-92ee-a33a62ea6a19"
	data := map[string]string{
		"1": "Test",
		"2": "1234",
		"3": "10",
	}

	res, err := client.NewWhatsapp().SendTemplate(sender, recipient, templateCode, data)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestSendText(t *testing.T) {
	sender := "2348120678278"
	recipient := "2348153207998"
	message := "Hello World"

	res, err := client.NewWhatsapp().SendText(sender, recipient, message)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestSendAudio(t *testing.T) {
	sender := "2348120678278"
	recipient := "2348153207998"
	message := "I am the best"
	link := "https://sample-videos.com/audio/mp3/crowd-cheering.mp3"

	res, err := client.NewWhatsapp().SendAudio(sender, recipient, message, link)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestSendVideo(t *testing.T) {
	sender := "2348120678278"
	recipient := "2348153207998"
	link := "https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4"

	res, err := client.NewWhatsapp().SendVideo(sender, recipient, link)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestSendSticker(t *testing.T) {
	sender := "2348120678278"
	recipient := "2348153207998"
	link := "https://studio.posit.us/api/samples/sticker.webp"

	res, err := client.NewWhatsapp().SendSticker(sender, recipient, link)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}

func TestSendLocation(t *testing.T) {
	sender := "2348120678278"
	recipient := "2348153207998"
	longitude := -46.662787
	latitude := -23.553610
	name := "Robbu Brazil"
	address := "Av. Angélica, 2530 - Bela Vista, São Paulo - SP, 01228-200"

	res, err := client.NewWhatsapp().SendLocation(sender, recipient, longitude, latitude, name, address)
	if err != nil {
		t.Error(err)
	}

	if res.Status != 200 {
		t.Error("res.Status: ", res.Status)
	}
}
