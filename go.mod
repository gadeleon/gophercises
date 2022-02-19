module github.com/gadeleon/gophercises

go 1.17

require github.com/gadeleon/gophercises/gophercise1/quiz v0.0.0

replace github.com/gadeleon/gophercises/gophercise1/quiz => ./gophercise1/internal/quiz
