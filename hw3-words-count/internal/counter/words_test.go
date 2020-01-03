package counter

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
			"first",
			"Since 2016, thousands of Gophers around the world have helped the Go project by sharing your thoughts via our annual Go Developer Survey. Your feedback has played an enormous role in driving changes to our language, ecosystem, and community, including the gopls language server, new error-handling mechanics, the module mirror, and so much more from the latest Go 1.13 release. And of course, we publicly share each year's results, so we can all benefit from the community's insights.",
			[]string{
				"times: 6 occurs words: [the]",
				"times: 3 occurs words: [and go]",
				"times: 2 occurs words: [from language of our so we your]",
				"times: 1 occurs words: [113 2016 all an annual around benefit by can changes community communitys course developer driving each ecosystem enormous errorhandling feedback gophers gopls has have helped in including insights latest mechanics mirror module more much new played project publicly release results role server share sharing since survey thoughts thousands to via world years]",
			},
		},
		{
			"second",
			"One FOUR six two FIve  Two three foUr four - 'three' six -three- sIX FoUr five five five five SIX (six) six",
			[]string{
				"times: 6 occurs words: [six]",
				"times: 5 occurs words: [five]",
				"times: 4 occurs words: [four]",
				"times: 3 occurs words: [three]",
				"times: 2 occurs words: [two]",
				"times: 1 occurs words: [one]",
			},
		},
		{
			"third",
			"cat and dog one dog two cats and one man",
			[]string{
				"times: 2 occurs words: [and dog one]",
				"times: 1 occurs words: [cat cats man two]",
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
