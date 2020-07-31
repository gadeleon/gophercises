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
//func NextQuestion(csv io.Reader)  []string {
//	q := csv.Read()
//	return q

//}

// Get Question from CSV

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
	fmt.Print(inf.Read())

}