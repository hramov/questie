package dialog

import "testing"

func TestQuestDialog_Next(t *testing.T) {
	questDialog := &QuestDialog{}
	questDialog.InitSteps()

	result := questDialog.Show()

	if result.Text != "Welcome to my new quest" {
		t.Errorf("wrong result: %s", result)
	}

	result = questDialog.Next("yes")

	if result.Text != "Get ready for questions" {
		t.Errorf("wrong result: %s", result)
	}

	result = questDialog.Show()
	if result.Text != "What is 2 + 2?" {
		t.Errorf("wrong result: %s", result)
	}

	result = questDialog.Next("3")
	if result.Text != "Nope\nWhat is 2 + 2?" {
		t.Errorf("wrong result: %s", result)
	}

}
