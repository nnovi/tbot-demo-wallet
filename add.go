package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func AddCommand(msg *tgbotapi.Message, args []string, out chan tgbotapi.Chattable) error {
	if len(args) != 2 {
		return fmt.Errorf("USAGE: ADD <COIN> <AMOUNT>")
	}

	coin := strings.ToUpper(strings.Trim(args[0], " "))
	if len(coin) == 0 {
		return fmt.Errorf("COIN required")
	}
	_, err := getCoinPrice(coin, "USDT")
	if err != nil {
		return fmt.Errorf("invalid COIN '%s'", coin)
	}

	amt, err := strconv.ParseFloat(args[1], 64)
	if err != nil || amt < 0 {
		return fmt.Errorf("invalid amount")
	}

	userWallet := getWallet(msg.Chat)
	balance, err := userWallet.Add(coin, amt)
	if err != nil {
		return err
	}

	out <- CreateReply(msg, fmt.Sprintf("ADD OK: %s = %.5f", coin, balance))
	return nil
}
