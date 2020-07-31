package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// Load filename f into CSV reader

func LoadCSV(f string) (*csv.Reader, error) {
	inf, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	csvf := csv.NewReader(inf)

	return csvf, err
}

// Read/Get next line of csv
func NextQuestion(csv *csv.Reader)  ([]string, error) {
	q, err := csv.Read()
	if err != nil {
		return nil, err
	}
	return q, err

}

// Prompt Question, know the answer

func PromptQuestion(num int, line []string) string {
	q := line[0]
	a := line[1]
	fmt.Printf("\nProblem #%d What is %s = ",num, q)

	return a
}

func GetAnswer() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	inp := s.Text()

	return inp
}

func EvalAnswer(myans string, ans string) bool {
	if myans == ans {
		return true
	}
	return false
}



func Grade(c int, t int) {
	eval := fmt.Sprintf("Questions correct: %d, Total Questions: %d", c, t)
	fmt.Println(eval)
}

func main(){
	// open csv
	inf, err := LoadCSV("problems.csv")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	crct := 0
	// Uses Limit as a max for questions
	// Hardcode as 10 for proof of concept.
	for i := 1; i <= 10; i++ {
		q, err := NextQuestion(inf)
		if err != nil {
			fmt.Println(err)
		}
		a := PromptQuestion(i, q)
		myans := GetAnswer()
		c := EvalAnswer(myans, a)
		if c == true {
			crct += 1
		}
	}
	// Replace t with total from option flags
	Grade(crct, 10)



}