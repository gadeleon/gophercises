package quiz

import (
	"reflect"
	"testing"
)

func TestNewQuiz(t *testing.T) {
	type args struct {
		folder   string
		filename string
	}
	tArgs := args{"../../../gophercise1/assets", "test_question.csv"}
	expected := &Quiz{
		[]*Question{{"1+1", "2"}},
		0,
	}
	tests := []struct {
		name    string
		args    args
		want    *Quiz
		wantErr bool
	}{
		{"Happy Path", tArgs, expected, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuiz(tt.args.folder, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuiz() = %v, want %v", got, tt.want)
			}
		})
	}
}
