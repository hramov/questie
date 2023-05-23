package internal

type Questier interface {
	InitSteps()
	Show() BotAnswer
	Next(answer string) BotAnswer
	Reset()
}

type Step struct {
	Name                string
	WelcomeMessage      string
	WelcomeMessageImage string
	PossibleAnswers     []string
	ExpectedAnswer      string
	NextStep            *Step
	RightAnswerMessage  string
	WrongAnswerMessage  string
}

type BotAnswer struct {
	Text     string
	Image    string
	Keyboard []string
}

type BotMessage struct {
	ChatId    int64
	Message   string
	MessageId int
}
