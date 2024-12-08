package quiz

import (
	"math/rand"
	"slices"
)

const MAX_CHOICES = 4
const NUMBER_OF_QUESTIONS = 10

type choice struct {
	answer    string
	isCorrect bool
}

type question struct {
	word    string
	choices [MAX_CHOICES]choice
}

func questions(db *db) []question {
	questions := make([]question, NUMBER_OF_QUESTIONS)

	for i, index := range randomQuestionIndexes() {
		choices := [MAX_CHOICES]choice{}

		for j, wi := range choiceIndexes([]uint8{index}) {
			if j == 0 { // first choice is always correct
				correctChoice := choice{
					answer:    db.words[index].English,
					isCorrect: true,
				}
				choices[j] = correctChoice
				continue
			}

			wrongChoice := choice{
				answer:    db.words[wi].English,
				isCorrect: false,
			}
			choices[j] = wrongChoice
		}

		rand.Shuffle(MAX_CHOICES, func(i, j int) {
			choices[i], choices[j] = choices[j], choices[i]
		})

		questions[i] = question{
			word:    db.words[index].Estonian.allForms(),
			choices: choices,
		}
	}

	return questions
}

func choiceIndexes(alreadyPicked []uint8) []uint8 {
	if len(alreadyPicked) == MAX_CHOICES {
		return alreadyPicked
	}

	randomIndex := uint8(rand.Intn(NUMBER_OF_QUESTIONS))
	for slices.Contains(alreadyPicked, randomIndex) {
		randomIndex = uint8(rand.Intn(NUMBER_OF_QUESTIONS))
	}

	return choiceIndexes(append(alreadyPicked, randomIndex))
}

func randomQuestionIndexes() []uint8 {
	randomIndexes := make([]uint8, NUMBER_OF_QUESTIONS)
	for i := range NUMBER_OF_QUESTIONS {
		randomIndexes[i] = uint8(i)
	}
	rand.Shuffle(NUMBER_OF_QUESTIONS, func(i, j int) {
		randomIndexes[i], randomIndexes[j] = randomIndexes[j], randomIndexes[i]
	})

	return randomIndexes
}
