package main

import (
	"fmt"
	"github.com/renta/golang-course/hw2-repeat-symbols/internal"
)

func main() {
	testStrings := map[string]string{
		"0123456789": "",
		"a4bc2d5e":   "aaaabccddddde",
		"abcd":       "abcd",
		"45":         "",
		`qwe\4\5`:    "qwe45",
		`qwe\45`:     "qwe44444",
		`qwe5\4`:     "qweeeee4",
		`qwe\\5`:     `qwe\\\\\`,
		`q\we`:       "",
		"ab/c":       "ab/c",
	}

	for input, desiredOutput := range testStrings {
		res, err := internal.StringRepack(input)
		fmt.Printf(
			"Input string: %s and the result should be: %s and the result is: %s does they are equal: %t error is: %v\n",
			input,
			desiredOutput,
			res,
			desiredOutput == res,
			err,
		)
	}
}
