package dialog

type QuestDialog struct {
	Dialog
}

func (qd *QuestDialog) InitSteps() {

	// TODO add more questions

	endStep := &Step{
		name:               "end",
		welcomeMessage:     "You won!",
		possibleAnswers:    []string{},
		expectedAnswer:     "Thank you",
		nextStep:           nil,
		rightAnswerMessage: "You are welcome",
		wrongAnswerMessage: "Bye",
	}

	firstQuestionStep := &Step{
		name:               "first_question",
		welcomeMessage:     "What is 2 + 2?",
		possibleAnswers:    []string{"2", "3", "4"},
		expectedAnswer:     "4",
		nextStep:           endStep,
		rightAnswerMessage: "Yeah!",
		wrongAnswerMessage: "Nope",
	}

	startStep := &Step{
		name:               "start",
		welcomeMessage:     "Welcome to my new quest",
		possibleAnswers:    []string{"Да!"},
		expectedAnswer:     "Да!",
		nextStep:           firstQuestionStep,
		rightAnswerMessage: "Get ready for questions",
		wrongAnswerMessage: "",
	}

	qd.startStep = startStep
	qd.currentStep = qd.startStep
}
