package main

import (
	"github.com/renta/golang-course/hw3-words-count/internal/counter"
	"log"
)

func main() {
	textForTest := []string{
		"Since 2016, thousands of Gophers around the world have helped the Go project by sharing your thoughts via our annual Go Developer Survey. Your feedback has played an enormous role in driving changes to our language, ecosystem, and community, including the gopls language server, new error-handling mechanics, the module mirror, and so much more from the latest Go 1.13 release. And of course, we publicly share each year's results, so we can all benefit from the community's insights.",
		"One FOUR six two FIve  Two three foUr four - 'three' six -three- sIX FoUr five five five five SIX (six) six",
		"cat and dog one dog two cats and one man",
	}

	for _, oneString := range textForTest {
		var wordsCounter = counter.WordsCounter{
			CountedWordsMap: make(map[string]int),
			NumbersOfWords:  make(map[int][]string),
		}
		res := wordsCounter.Top10(oneString)
		log.Println("------------------------------------")
		for _, resultString := range res {
			log.Println(resultString)
		}
		log.Println("------------------------------------")
	}
}
