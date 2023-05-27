package quests

import "github.com/hramov/questie/internal"

type ExampleQuest struct {
	internal.Dialog
}

func (eq *ExampleQuest) InitSteps() {
	endStep := &internal.Step{
		Name:                "end",
		WelcomeMessage:      "",
		WelcomeMessageImage: internal.AssetsFolder + "",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "",
		NextStep:            nil,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "",
	}

	startStep := &internal.Step{
		Name:                "start",
		WelcomeMessage:      "",
		WelcomeMessageImage: internal.AssetsFolder + "",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "",
		NextStep:            endStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "",
	}

	eq.StartStep = startStep
	eq.CurrentStep = eq.StartStep
}
