package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const tgBotToken = "2135305458:AAGt8DbF_b6pG98H4cKF8AMCLjmKZViP3fQ"

type CommandHandler func(msg *tgbotapi.Message, args []string, out chan tgbotapi.Chattable) error

var commands = map[string]CommandHandler{
	"ADD":  AddCommand,
	"SUB":  SubCommand,
	"DEL":  DelCommand,
	"SHOW": ShowCommand,
}

func HelpCommand(msg *tgbotapi.Message, _ []string, out chan tgbotapi.Chattable) error {
	out <- CreateReply(msg, `ADD <COIN> <AMOUNT>
SUB <COIN> <AMOUNT>
DEL <COIN>
SHOW
HELP
.....
`)

	return nil
}

func CreateReply(sourceMessage *tgbotapi.Message, reply string) tgbotapi.Chattable {
	result := tgbotapi.NewMessage(sourceMessage.Chat.ID, reply)
	result.ReplyToMessageID = sourceMessage.MessageID
	return result
}

func main() {
	bot, err := tgbotapi.NewBotAPI(tgBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	outChannel := make(chan tgbotapi.Chattable)

	go func() {
		for msg := range outChannel {
			bot.Send(msg)
		}
	}()

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := strings.Split(update.Message.Text, " ")

		if handler, ok := commands[strings.ToUpper(msg[0])]; ok {
			go func() {
				err := handler(update.Message, msg[1:], outChannel)
				if err != nil {
					outChannel <- CreateReply(update.Message, fmt.Sprintf("%v", err))
				}
			}()
		} else {
			go func() {
				_ = HelpCommand(update.Message, msg[1:], outChannel)
			}()
		}
	}
}
