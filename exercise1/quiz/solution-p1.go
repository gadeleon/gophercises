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

	//return fmt.Sprintf("You put: %s,  I want: %s\n", inp, a)
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
	fmt.Printf("Questions correct: %d, Total Questions: %d", c, t)
}
// Get next question

// Evaluate answer

// Reach Limit

// Report grade

func main(){
	inf, err := LoadCSV("problems.csv")
	if err != nil {
		fmt.Print(err)
	}
	q, err := NextQuestion(inf)
	if err != nil {
		fmt.Print(err)
	}
	a := PromptQuestion(1, q)
	// fmt.Print(a)
	myans := GetAnswer()
	//fmt.Print(myans)
	c := EvalAnswer(myans, a)
	fmt.Println(c)

	// Get limit of questions
	// change this loop to limit
	// total := options.limit
	crct := 0
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}


}