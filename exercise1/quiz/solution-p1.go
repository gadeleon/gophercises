package main

import (
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

func PromptQuestion(line []string) string {
	q := line[0]
	a := line[1]
	fmt.Printf("\n%s = ", q)
	var inp string
	i, err := fmt.Scanln(&inp)
	if err != nil {
		fmt.Print("UHHHHH")
	}
	return fmt.Sprintf("You put: %s,  I want: %s\n", i, a)
}

// Get limit of questions

// Read input, log if correct

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
	a := PromptQuestion(q)
	fmt.Print(a)

}