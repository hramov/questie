package internal

import "fmt"

type Dialog struct {
	LastMessageId    int64
	StartStep        *Step
	CurrentStep      *Step
	CurrentStepIndex int
}

func (d *Dialog) Show() BotAnswer {
	return BotAnswer{
		Text:     d.CurrentStep.WelcomeMessage,
		Image:    d.CurrentStep.WelcomeMessageImage,
		Keyboard: d.CurrentStep.PossibleAnswers,
	}
}

func (d *Dialog) Next(answer string) BotAnswer {
	if d.CurrentStep.ExpectedAnswer == answer {
		message := d.CurrentStep.RightAnswerMessage
		if d.CurrentStep.NextStep != nil {
			d.CurrentStepIndex++
			d.CurrentStep = d.CurrentStep.NextStep
			return BotAnswer{
				Text:     fmt.Sprintf("%s\n%s", message, d.CurrentStep.WelcomeMessage),
				Image:    "",
				Keyboard: d.CurrentStep.PossibleAnswers,
			}
		}
		return BotAnswer{
			Text: message,
		}
	}
	if d.CurrentStep.NextStep != nil {
		return BotAnswer{
			Text:     fmt.Sprintf("%s\n%s", d.CurrentStep.WrongAnswerMessage, d.CurrentStep.WelcomeMessage),
			Image:    "",
			Keyboard: d.CurrentStep.PossibleAnswers,
		}
	}
	return BotAnswer{
		Text:     fmt.Sprintf("%s", d.CurrentStep.WrongAnswerMessage),
		Image:    "",
		Keyboard: d.CurrentStep.PossibleAnswers,
	}
}

func (d *Dialog) Reset() {
	d.CurrentStep = d.StartStep
	d.CurrentStepIndex = 0
}
