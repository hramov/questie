package dialog

import "testing"

type TestQuest struct {
	Dialog
}

func (tq *TestQuest) InitSteps() {
	endStep := &Step{
		name:               "end",
		welcomeMessage:     "You won!",
		possibleAnswers:    []string{"Thank you"},
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
		possibleAnswers:    []string{"Yes!"},
		expectedAnswer:     "Yes!",
		nextStep:           firstQuestionStep,
		rightAnswerMessage: "Get ready for questions",
		wrongAnswerMessage: "",
	}

	tq.startStep = startStep
	tq.currentStep = tq.startStep
}

func TestQuestDialog_Next(t *testing.T) {
	quest := &TestQuest{}
	quest.InitSteps()

	result := quest.Show()
	if result.Text != "Welcome to my new quest" {
		t.Errorf("wrong result: %s / %s", "Welcome to my new quest", result.Text)
	}

	result = quest.Next("Yes!")
	if result.Text != "Get ready for questions\nWhat is 2 + 2?" {
		t.Errorf("wrong result: %s / %s", "Get ready for questions\nWhat is 2 + 2?", result.Text)
	}

	result = quest.Next("2")
	if result.Text != "Nope\nWhat is 2 + 2?" {
		t.Errorf("wrong result: %s / %s", "Nope\nWhat is 2 + 2?", result.Text)
	}

	result = quest.Next("4")
	if result.Text != "Yeah!\nYou won!" {
		t.Errorf("wrong result: %s / %s", "Yeah!\nYou won!", result.Text)
	}

	result = quest.Next("Thank you")
	if result.Text != "You are welcome" {
		t.Errorf("wrong result: %s / %s", "You are welcome", result.Text)
	}

	result = quest.Next("OK")
	if result.Text != "Bye" {
		t.Errorf("wrong result: %s / %s", "Bye", result.Text)
	}
}
