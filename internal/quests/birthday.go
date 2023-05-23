package quests

import "github.com/hramov/questie/internal"

type BirthdayQuest struct {
	internal.Dialog
}

func (bq *BirthdayQuest) InitSteps() {
	endStep := &internal.Step{
		Name:               "end",
		WelcomeMessage:     "You won!",
		PossibleAnswers:    []string{},
		ExpectedAnswer:     "Thank you",
		NextStep:           nil,
		RightAnswerMessage: "You are welcome",
		WrongAnswerMessage: "Bye",
	}

	firstQuestionStep := &internal.Step{
		Name:               "first_question",
		WelcomeMessage:     "What is 2 + 2?",
		PossibleAnswers:    []string{"2", "3", "4"},
		ExpectedAnswer:     "4",
		NextStep:           endStep,
		RightAnswerMessage: "Yeah!",
		WrongAnswerMessage: "Nope",
	}

	startStep := &internal.Step{
		Name:                "start",
		WelcomeMessage:      "Welcome to my new quest",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/Image.jpg",
		PossibleAnswers:     []string{"Да!"},
		ExpectedAnswer:      "Да!",
		NextStep:            firstQuestionStep,
		RightAnswerMessage:  "Get ready for questions",
		WrongAnswerMessage:  "",
	}

	bq.StartStep = startStep
	bq.CurrentStep = bq.StartStep
}
