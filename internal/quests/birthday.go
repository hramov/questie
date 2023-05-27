package quests

import "github.com/hramov/questie/internal"

type BirthdayQuest struct {
	internal.Dialog
}

func (bq *BirthdayQuest) InitSteps() {
	endStep := &internal.Step{
		Name:                "end",
		WelcomeMessage:      "С днем рождения, дорогая моя Настенька. Я тебя очень сильно люблю!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/8.jpg",
		PossibleAnswers:     []string{"И я тебя"},
		ExpectedAnswer:      "И я тебя",
		NextStep:            nil,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "",
	}

	sixthQuestionStep := &internal.Step{
		Name:                "sixth_question",
		WelcomeMessage:      "Один из наиболее важных элементов колеса, представляющий собой упругую резино-металло-тканевую оболочку, установленную на обод диска",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/7.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "123",
		NextStep:            endStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	fifthQuestionStep := &internal.Step{
		Name:                "fifth_question",
		WelcomeMessage:      "Просто хочу напомнить, что скоро мы поедем на поезде...",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/6.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "123",
		NextStep:            sixthQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	forthQuestionStep := &internal.Step{
		Name:                "forth_question",
		WelcomeMessage:      "Чтобы получить следующую подсказку нужно знать, чем занимался этот крутой мужик",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/5.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "123",
		NextStep:            fifthQuestionStep,
		RightAnswerMessage:  "Точно-точно",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	thirdQuestionStep := &internal.Step{
		Name:                "third_question",
		WelcomeMessage:      "Следующий вопрос пришел в нашу редакцию от кота Базилио",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/4.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "123",
		NextStep:            forthQuestionStep,
		RightAnswerMessage:  "Да, все верно!",
		WrongAnswerMessage:  "Нет... еще это никнейм Папича",
	}

	secondQuestionStep := &internal.Step{
		Name:                "second_question",
		WelcomeMessage:      "Ищи вопрос в Машине",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/3.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "Конденсатор",
		NextStep:            thirdQuestionStep,
		RightAnswerMessage:  "Есть!",
		WrongAnswerMessage:  "Может быть, не та \"машина\"...",
	}

	firstQuestionStep := &internal.Step{
		Name:                "first_question",
		WelcomeMessage:      "Для начала вспомни дату нашего первого поцелуя. Ответом на это задание, будет являться сумма всех цифр этой даты!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/2.jpg",
		PossibleAnswers:     []string{"19", "21", "23", "25"},
		ExpectedAnswer:      "21",
		NextStep:            secondQuestionStep,
		RightAnswerMessage:  "Супер!",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	startStep := &internal.Step{
		Name:                "start",
		WelcomeMessage:      "Привет :) Сегодня замечательный день, и я надеюсь у тебя хорошее настроение :) Как ты уже догадалась, просто так подарок я тебе не отдам :) Придется поискать и пройти квест! Но ты же любишь приключения! Отправляйся на поиски, тебя ждет увлекательная игра!!!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/1.jpg",
		PossibleAnswers:     []string{"Вперед!"},
		ExpectedAnswer:      "Вперед!",
		NextStep:            firstQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "",
	}

	bq.StartStep = startStep
	bq.CurrentStep = bq.StartStep
}
