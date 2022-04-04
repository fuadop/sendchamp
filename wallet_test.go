package sendchamp_test

import (
	"testing"
)

func TestWalletBalance(t *testing.T) {
	res, err := client.WalletBalance()
	if err != nil {
		t.Error(err)
	}

	if res.Status != "success" {
		t.Error("res.Status: ", res.Status)
	}
}
