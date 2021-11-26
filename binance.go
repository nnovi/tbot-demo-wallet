package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type binancePriceResponce struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

func getCoinPrice(fromCoin, toCoin string) (price float64, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s%s", fromCoin, toCoin))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	jsonResp := binancePriceResponce{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return
	}

	if jsonResp.Code != 0 {
		err = fmt.Errorf("COIN '%s' error", fromCoin)
		return
	}

	return jsonResp.Price, nil
}
