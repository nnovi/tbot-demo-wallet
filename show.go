package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func ShowCommand(msg *tgbotapi.Message, _ []string, out chan tgbotapi.Chattable) error {
	userWallet := getWallet(msg.Chat)

	result := strings.Builder{}

	rubPrice, err := getCoinPrice("USDT", "RUB")
	if err != nil {
		result.WriteString(fmt.Sprintf("error getting RUB/USDT rate\n"))
	}

	var totalLocal float64
	for coin, amt := range userWallet {
		if rubPrice > 0 {
			coinPrice, err := getCoinPrice(coin, "USDT")
			if err == nil {
				result.WriteString(fmt.Sprintf("%s: %.5f [%.2f руб.]\n", coin, amt, amt*coinPrice*rubPrice))
				totalLocal += amt * coinPrice * rubPrice
			} else {
				result.WriteString(fmt.Sprintf("%s: %.5f [%v]\n", coin, amt, err))
			}
		} else {
			result.WriteString(fmt.Sprintf("%s: %.5f\n", coin, amt))
		}
	}
	if totalLocal > 0 {
		result.WriteString("=============================\n")
		result.WriteString(fmt.Sprintf("TOTAL: %.2f руб.", totalLocal))
	}

	out <- CreateReply(msg, result.String())

	return nil
}
