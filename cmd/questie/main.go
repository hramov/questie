package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hramov/questie/internal"
	"github.com/hramov/questie/internal/dialog"
	"github.com/hramov/questie/pkg/telegram"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type ClientMap = map[int64]dialog.Questier

type Clients struct {
	mu   sync.Mutex
	data ClientMap
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot := telegram.NewBot(os.Getenv("TG_TOKEN"))
	go bot.Start()

	clients := &Clients{
		mu:   sync.Mutex{},
		data: make(ClientMap),
	}

	log.Println("Bot started!")

	for update := range bot.GetUpdates() {
		chatID := update.Message.Chat.ID
		text := update.Message.Text

		if text == "/start" {
			clients.mu.Lock()
			currentQuest, exists := clients.data[chatID]

			if exists {
				currentQuest.Reset()
				delete(clients.data, chatID)
			}

			clients.mu.Unlock()

			keyboard := bot.GenerateStartKeyboard(internal.Quests)
			msg := tgbotapi.NewMessage(chatID, "Выбери квест!")
			msg.ReplyMarkup = keyboard
			bot.SendMessage(msg)
		} else {
			switch text {
			case "День рождения":
				questDialog := &dialog.QuestDialog{}
				questDialog.InitSteps()

				clients.mu.Lock()
				clients.data[chatID] = questDialog
				clients.mu.Unlock()

				sendHelloMessage(questDialog, chatID, bot)
				break
			default:
				clients.mu.Lock()
				currentQuest, exists := clients.data[chatID]
				clients.mu.Unlock()

				if !exists {
					sendNoActiveQuestError(chatID, bot)
					continue
				}

				result := currentQuest.Next(text)
				keyboard := bot.GenerateKeyboard(result)
				msg := tgbotapi.NewMessage(chatID, result.Text)
				msg.ReplyMarkup = keyboard
				bot.SendMessage(msg)
			}
		}
	}
}

func sendHelloMessage(dialog dialog.Questier, chatID int64, bot *telegram.Bot) {
	dialog.Reset()
	result := dialog.Show()
	keyboard := bot.GenerateKeyboard(result)
	msg := tgbotapi.NewMessage(chatID, result.Text)
	msg.ReplyMarkup = keyboard
	bot.SendMessage(msg)
}

func sendNoActiveQuestError(chatID int64, bot *telegram.Bot) {
	keyboard := bot.GenerateStartKeyboard(internal.Quests)
	msg := tgbotapi.NewMessage(chatID, "В данный момент не запущено ни одного квеста. Выберите квест ниже")
	msg.ReplyMarkup = keyboard
	bot.SendMessage(msg)
}
