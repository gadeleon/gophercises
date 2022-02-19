package main

import (
	"fmt"

	"github.com/gadeleon/gophercises/gophercise1/quiz"
)

func main() {
	q, err := quiz.NewQuiz("gophercise1/assets", "questions.csv")
	if err != nil {
		print(err)
	}
	fmt.Println(q)
}
