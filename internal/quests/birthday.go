package quests

import "github.com/hramov/questie/internal"

type BirthdayQuest struct {
	internal.Dialog
}

func (bq *BirthdayQuest) InitSteps() {
	endStep := &internal.Step{
		Name:               "end",
		WelcomeMessage:     "С днем рождения, дорогая моя Настенька. Я тебя очень сильно люблю!",
		PossibleAnswers:    []string{"И я тебя"},
		ExpectedAnswer:     "И я тебя",
		NextStep:           nil,
		RightAnswerMessage: "",
		WrongAnswerMessage: "",
	}

	seventhQuestionStep := &internal.Step{
		Name:                "first_question",
		WelcomeMessage:      "Ответ на это задание тебе подскажет Машина!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/6.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "21",
		NextStep:            endStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	sixthQuestionStep := &internal.Step{
		Name:                "first_question",
		WelcomeMessage:      "Ответ на это задание тебе подскажет Машина!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/6.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "21",
		NextStep:            seventhQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	fifthQuestionStep := &internal.Step{
		Name:                "first_question",
		WelcomeMessage:      "Ответ на это задание тебе подскажет Машина!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/6.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "21",
		NextStep:            sixthQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	forthQuestionStep := &internal.Step{
		Name:                "first_question",
		WelcomeMessage:      "Ответ на это задание тебе подскажет Машина!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/5.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "21",
		NextStep:            fifthQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Что-то путаешь, дорогая :-) Попробуем еще раз",
	}

	thirdQuestionStep := &internal.Step{
		Name:                "third_question",
		WelcomeMessage:      "Вопрос на знание моих интересов. Закончи цитату: \"Сын мой, в день, когда ты родился, сами леса Лордерона прошептали твое имя - ?\"",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/4.jpg",
		PossibleAnswers:     []string{"Кенарий", "Артас", "Менетил", "Тирион"},
		ExpectedAnswer:      "Артас",
		NextStep:            forthQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Нет... еще это никнейм Папича",
	}

	secondQuestionStep := &internal.Step{
		Name:                "second_question",
		WelcomeMessage:      "Ищи вопрос в Машине",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/3.jpg",
		PossibleAnswers:     []string{},
		ExpectedAnswer:      "Конденсатор",
		NextStep:            thirdQuestionStep,
		RightAnswerMessage:  "",
		WrongAnswerMessage:  "Может быть, не та \"машина\"...",
	}

	firstQuestionStep := &internal.Step{
		Name:                "first_question",
		WelcomeMessage:      "Для начала вспомни дату нашего первого поцелуя. Ответом на это задание, будет являться сумма всех цифр этой даты!",
		WelcomeMessageImage: internal.AssetsFolder + "birthday/2.jpg",
		PossibleAnswers:     []string{"19", "21", "23", "25"},
		ExpectedAnswer:      "21",
		NextStep:            secondQuestionStep,
		RightAnswerMessage:  "",
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
