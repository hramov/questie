package dialog

import "fmt"

type Questier interface {
	InitSteps()
	Show() BotAnswer
	Next(answer string) BotAnswer
	Reset()
}

type Step struct {
	name               string
	welcomeMessage     string
	possibleAnswers    []string
	expectedAnswer     string
	nextStep           *Step
	rightAnswerMessage string
	wrongAnswerMessage string
}

type BotAnswer struct {
	Text     string
	Image    string
	Keyboard []string
}

type Dialog struct {
	lastMessageId    int64
	startStep        *Step
	currentStep      *Step
	currentStepIndex int
}

func (d *Dialog) Show() BotAnswer {
	return BotAnswer{
		Text:     d.currentStep.welcomeMessage,
		Image:    "",
		Keyboard: d.currentStep.possibleAnswers,
	}
}

func (d *Dialog) Next(answer string) BotAnswer {
	if d.currentStep.expectedAnswer == answer {
		message := d.currentStep.rightAnswerMessage
		if d.currentStep.nextStep != nil {
			d.currentStepIndex++
			d.currentStep = d.currentStep.nextStep
			return BotAnswer{
				Text:     fmt.Sprintf("%s\n%s", message, d.currentStep.welcomeMessage),
				Image:    "",
				Keyboard: d.currentStep.possibleAnswers,
			}
		}
		return BotAnswer{
			Text: message,
		}
	}
	return BotAnswer{
		Text:     fmt.Sprintf("%s\n%s", d.currentStep.wrongAnswerMessage, d.currentStep.welcomeMessage),
		Image:    "",
		Keyboard: d.currentStep.possibleAnswers,
	}
}

func (d *Dialog) Reset() {
	d.currentStep = d.startStep
	d.currentStepIndex = 0
}
