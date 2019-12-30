package internal

import (
	"reflect"
	"testing"
)

func TestWordsCounter_Top10(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			"second",
			"One FOUR six two FIve  Two three foUr four - 'three' six -three- sIX FoUr five five five five SIX (six) six",
			[]string{
				"times: 6 occurs words: [six]",
				"times: 5 occurs words: [five]",
				"times: 4 occurs words: [four]",
				"times: 3 occurs words: [three]",
				"times: 2 occurs words: [two ]",
				"times: 1 occurs words: [one]",
			},
		},
		{
			"third",
			"cat and dog one dog two cats and one man",
			[]string{
				"times: 2 occurs words: [and dog one]",
				"times: 1 occurs words: [cat two cats man]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WordsCounter{
				CountedWordsMap: make(map[string]int),
				NumbersOfWords:  make(map[int][]string),
			}
			if got := w.Top10(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Top10() = %v, want %v", got, tt.want)
			}
		})
	}
}