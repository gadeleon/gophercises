package quiz

import (
	"encoding/csv"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Question struct {
	Question string
	Answer   string
}

type Quiz struct {
	Questions []*Question
	Score     int
}

// Open CSV File
// Create Quiz with Questions
// Go through Quiz.Questions, and ask Question. Score ++ if correct

// func matchName(filename string) filepath.WalkFunc {
// 	return func(path string, info fs.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	if info.Name() == filename {
// 		filename = info.
// 	}
// 	return nil
// }

func NewQuiz(folder string, filename string) (*Quiz, error) {
	// find the file
	qs := make([]*Question, 0)
	walk := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == filename {
			// Make the filepath to open
			fp := folder + "/" + filename
			inf, err := os.Open(fp)
			if err != nil {
				return err
			}
			defer inf.Close()
			// open as csv
			csvReader := csv.NewReader(inf)
			for {
				rec, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					return err
				}
				qs = append(qs, &Question{rec[0], rec[1]})

			}
		}
		return nil
	})
	if walk != nil {
		return nil, walk
	}
	quiz := &Quiz{
		qs,
		0,
	}
	return quiz, nil

}
