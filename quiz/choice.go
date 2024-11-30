package quiz

const MAX_CHOICES = 4

type choice struct {
	word      string
	isCorrect bool
}

type question struct {
	word    string
	choices [MAX_CHOICES]choice
}
