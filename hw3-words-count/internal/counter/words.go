package counter

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type WordsCounter struct {
	CountedWordsMap map[string]int
	NumbersOfWords  map[int][]string
	TopTenNumbers   []int
}

func (w *WordsCounter) Top10(text string) []string {
	w.countWords(text)
	w.summariseCountedWords()
	w.getTopTenWords()

	return w.formTheResult()
}

func (w *WordsCounter) countWords(stringToCount string) {
	splittedString := strings.Split(stringToCount, " ")
	for _, wordAfterSplit := range splittedString {
		preparedWord := prepareWord(wordAfterSplit)
		if preparedWord == "" {
			continue
		}
		w.CountedWordsMap[preparedWord]++
	}
}

func prepareWord(word string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	word = regex.ReplaceAllString(word, "")

	return strings.ToLower(strings.Trim(word, ""))
}

func (w *WordsCounter) summariseCountedWords() {
	for word, count := range w.CountedWordsMap {
		w.NumbersOfWords[count] = append(w.NumbersOfWords[count], word)
	}
}

func (w *WordsCounter) getTopTenWords() {
	var sliceOfWordsNumbers []int
	for count := range w.NumbersOfWords {
		sliceOfWordsNumbers = append(sliceOfWordsNumbers, count)
	}
	sort.Slice(sliceOfWordsNumbers, func(i, j int) bool { return sliceOfWordsNumbers[i] > sliceOfWordsNumbers[j] })
	for i, val := range sliceOfWordsNumbers {
		if i <= 9 {
			w.TopTenNumbers = append(w.TopTenNumbers, val)
		}
	}
}

func (w *WordsCounter) formTheResult() []string {
	var resultSlice []string
	for _, number := range w.TopTenNumbers {
		sort.Strings(w.NumbersOfWords[number])
		resultSlice = append(resultSlice, fmt.Sprintf("times: %d occurs words: %v", number, w.NumbersOfWords[number]))
	}
	return resultSlice
}
