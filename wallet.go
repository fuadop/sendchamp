package sendchamp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpointWalletBalance = "/wallet/wallet_balance"
)

type walletBalanceResponse struct {
	Status  string                    `json:"status"`
	Code    string                    `json:"code"`
	Message string                    `json:"message"`
	Data    walletBalanceResponseData `json:"data"`
}

type walletBalanceResponseData struct {
	UID              string                           `json:"uid"`
	BusinessUID      string                           `json:"busines_uid"`
	Type             string                           `json:"type"`
	AvailableBalance string                           `json:"available_balance"`
	Currency         string                           `json:"currency"`
	BusinessCurrency string                           `json:"business_currency"`
	Details          walletBalanceResponseDataDetails `json:"details"`
}

type walletBalanceResponseDataDetails struct {
	BaseAmount       string  `json:"base_amount"`
	BaseCurrency     string  `json:"base_currency"`
	ExchangeRage     string  `json:"exchange_rate"`
	BusinessAmount   float64 `json:"business_amount"`
	BusinessCurrency string  `json:"business_currency"`
}

func (c *Client) WalletBalance() (walletBalanceResponse, error) {
	url := fmt.Sprint(c.baseURL, endpointWalletBalance)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return walletBalanceResponse{}, err
	}

	// add necessary request headers
	addHeaders(req, c)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return walletBalanceResponse{}, err
	}

	defer res.Body.Close()

	r := walletBalanceResponse{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return walletBalanceResponse{}, err
	}

	return r, nil
}
