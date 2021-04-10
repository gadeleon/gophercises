package main

import (
	"os"
	"reflect"
	"testing"
)



//func TestGetFileNameList(t *testing.T) {
//	// TODO: at some point mock the directory
//	// TODO: Fix error case (by letting user specify path)
//	expected := []string{"gopher.json","story01.json"}
//	actual, _ := getFileNameList("./")
//
//	if !Equal(expected,actual)  {
//		t.Errorf("Got %v, Wanted %v", actual, expected)
//	}
//	// Test On Directory Doesn't Exist /bottomdollar
//	//_, err := getFileNameList()
//}


//func TestParseStory(t *testing.T) {
//	s, _ := parseStory(JSONblob)
//	title := s.Instances["new-york"].Title
//	if title != "Visiting New York" {
//		t.Error("Got:", title, "expected:", "Visiting New York")
//	}
//	o := s.Instances["debate"].Options
//	ol := len(o)
//	// Debate has 3 options
//	expected := 3
//
//	if ol != expected {
//		t.Errorf("Got %d, Expected %d", ol, expected)
//	}
//
//}
//// Get the correct route if we pass a specific endpoint
//// If we get /index, we get the list.
//// If we get /new-york we go to that part of the arc
//func TestRouting(t *testing.T) {
//
//}
//
//// Parse multiple stories accepts a story
//// Reads slice  of Stories
//// Struct Index that contains a slice of Stories
//// test parse stories index and then test stories
//func TestNewParseStories(t *testing.T) {
//	s, _ := parseStories(JSONblob)
//	index := len(s.Stories)
//	expected := 3
//	if index != expected {
//		t.Errorf("Wanted %d got %d",expected, index)
//	}
//}

func Test_ingestFile(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want StoryJson
	}{
		{"Story01", args{"story01.json"}, StoryJson{ StoryArc {  "story",nil,nil} }  },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ingestFile(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ingestFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFileNameList(t *testing.T) {
	curdir,_ := os.Getwd()
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"expected", args{curdir}, []string{"gopher.json","story01.json"}, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFileNameList(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileNameList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFileNameList() got = %v, want %v", got, tt.want)
			}
		})
	}
}