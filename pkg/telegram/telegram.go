package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hramov/questie/internal"
)

type Bot struct {
	token  string
	update chan tgbotapi.Update
	bot    *tgbotapi.BotAPI
}

func NewBot(token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	messageChan := make(chan tgbotapi.Update, 10)
	return &Bot{
		token:  token,
		update: messageChan,
		bot:    bot,
	}
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		b.update <- update
	}
}

func (b *Bot) GetUpdates() chan tgbotapi.Update {
	return b.update
}

func (b *Bot) GenerateKeyboard(result internal.BotAnswer) *tgbotapi.ReplyKeyboardMarkup {
	keys := result.Keyboard

	if keys == nil || len(keys) == 0 {
		return nil
	}
	var buttons []tgbotapi.KeyboardButton

	for _, key := range keys {
		buttons = append(buttons, tgbotapi.KeyboardButton{
			Text: key,
		})
	}

	return &tgbotapi.ReplyKeyboardMarkup{
		Keyboard:              [][]tgbotapi.KeyboardButton{buttons},
		ResizeKeyboard:        true,
		OneTimeKeyboard:       true,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}

func (b *Bot) GenerateStartKeyboard(keys []string) *tgbotapi.ReplyKeyboardMarkup {

	if keys == nil || len(keys) == 0 {
		return nil
	}
	var buttons []tgbotapi.KeyboardButton

	for _, key := range keys {
		buttons = append(buttons, tgbotapi.KeyboardButton{
			Text: key,
		})
	}

	return &tgbotapi.ReplyKeyboardMarkup{
		Keyboard:              [][]tgbotapi.KeyboardButton{buttons},
		ResizeKeyboard:        true,
		OneTimeKeyboard:       true,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}

func (b *Bot) SendMessage(message tgbotapi.MessageConfig) {
	_, err := b.bot.Send(message)
	if err != nil {
		panic(fmt.Sprintf("cannot send message: %v", err.Error()))
	}
}

func (b *Bot) SendImage(message tgbotapi.PhotoConfig) {
	_, err := b.bot.Send(message)
	if err != nil {
		panic(fmt.Sprintf("cannot send message: %v", err.Error()))
	}
}
