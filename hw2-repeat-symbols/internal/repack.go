package internal

import (
	"errors"
	"strings"
	"unicode"
)

func isBackslashOrNumberWasBackslashed(isEscape bool, symbol rune) bool {
	return isEscape == true && isNumberOrBackslash(symbol) == false
}

func isValidSymbol(isEscape bool, symbol rune) bool {
	//                                 int32 for / symbol
	return unicode.IsLetter(symbol) || symbol == 47 || (isEscape == true && isNumberOrBackslash(symbol) == true)
}

func isNumberOrBackslash(symbol rune) bool {
	return unicode.IsNumber(symbol) || symbol == 92
}

func needToRepeatSymbol(isEscape bool, symbol rune) bool {
	return isEscape == false && unicode.IsNumber(symbol)
}

func StringRepack(s string) (string, error) {
	var result []string

	var letterToPush string
	var isEscape bool

	for _, symbol := range s {
		if isBackslashOrNumberWasBackslashed(isEscape, symbol) {
			return "", errors.New("wrong string format")
		}
		if isEscape == false && symbol == 92 { // int32 code for \ symbol
			isEscape = true
			continue
		}
		if isValidSymbol(isEscape, symbol) {
			letterToPush = string(symbol)
			result = append(result, letterToPush)
			isEscape = false
			continue
		}
		if needToRepeatSymbol(isEscape, symbol) {
			for i := 0; i < int(symbol-'1'); i++ {
				result = append(result, letterToPush)
			}
		}
	}

	return strings.Join(result, ""), nil
}
