package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func delCommand(msg *tgbotapi.Message, args []string, out chan tgbotapi.Chattable) error {
	if len(args) != 1 {
		return fmt.Errorf("USAGE: DEL <COIN>")
	}

	coin := strings.ToUpper(strings.Trim(args[0], " "))
	if len(coin) == 0 {
		return fmt.Errorf("COIN required")
	}

	userWallet := getWallet(msg.Chat)
	err := userWallet.Delete(coin)
	if err != nil {
		return err
	}

	out <- createReply(msg, fmt.Sprintf("DEL OK: %s", coin))
	return nil
}
