package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

type wallet map[string]float64

var walletsLock = sync.Mutex{}
var wallets = make(map[int64]wallet)

func getWallet(chat *tgbotapi.Chat) wallet {
	walletsLock.Lock()
	defer walletsLock.Unlock()

	w, ok := wallets[chat.ID]
	if !ok {
		w = make(wallet)
		wallets[chat.ID] = w
	}
	return w
}

func (w wallet) Add(coin string, amt float64) (float64, error) {
	walletsLock.Lock()
	defer walletsLock.Unlock()

	w[coin] = w[coin] + amt
	return w[coin], nil
}

func (w wallet) Sub(coin string, amt float64) (float64, error) {
	walletsLock.Lock()
	defer walletsLock.Unlock()

	if w[coin] < amt {
		return w[coin], fmt.Errorf("not enought balance for %s", coin)
	}
	w[coin] = w[coin] - amt
	if w[coin] == 0 {
		delete(w, coin)
	}
	return w[coin], nil
}

func (w wallet) Delete(coin string) error {
	walletsLock.Lock()
	defer walletsLock.Unlock()

	delete(w, coin)
	return nil
}

func (w wallet) Balance(coin string) float64 {
	return w[coin]
}
