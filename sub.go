package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func SubCommand(msg *tgbotapi.Message, args []string, out chan tgbotapi.Chattable) error {
	if len(args) != 2 {
		return fmt.Errorf("USAGE: SUB <COIN> <AMOUNT>")
	}

	coin := strings.ToUpper(strings.Trim(args[0], " "))
	if len(coin) == 0 {
		return fmt.Errorf("COIN required")
	}

	amt, err := strconv.ParseFloat(args[1], 64)
	if err != nil || amt < 0 {
		return fmt.Errorf("invalid amount")
	}

	userWallet := getWallet(msg.Chat)
	balance, err := userWallet.Sub(coin, amt)
	if err != nil {
		return err
	}

	out <- CreateReply(msg, fmt.Sprintf("SUB OK: %s = %.5f", coin, balance))
	return nil
}
