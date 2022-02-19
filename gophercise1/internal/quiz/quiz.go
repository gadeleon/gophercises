package quiz

type Question struct {
	Question string
	Answer   string
}

type Quiz struct {
	Questions []Question
	result    bool
	Score     int
}
