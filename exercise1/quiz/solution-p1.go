package quiz

import (
	"encoding/csv"
	"io"
)

// Load into CSV

func LoadCSV(r io.Reader) *Reader {
	csv := csv.NewReader(r)

	return csv
}

// Read/Get next question
func NextQuestion(csv io.Reader) []string {
	return csv.Read()
}


// Get limit of questions

// Ask question

// Read input, log if correct

// Get next question

// Evaluate answer

// Reach Limit

// Report grade