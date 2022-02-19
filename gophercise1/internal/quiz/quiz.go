package quiz

type Question struct {
	Question string
	Answer   string
}

type Quiz struct {
	Source    string
	Questions []Question
	Score     int
}

// Open CSV File
// Create Quiz with Questions
// Go through Quiz.Questions, and ask Question. Score ++ if correct
