package interfaces

type BotMessage struct {
	ChatId    int64
	Message   string
	MessageId int
}

type BotHandler interface {
	Start()
	SendMessage(botMessage *BotMessage)
}
