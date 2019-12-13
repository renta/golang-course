package internal

import (
	"strings"
	"unicode"
)

func isEscapedNumberOrBackslash(isEscaped bool, symbol rune) bool {
	return isEscaped == true && (unicode.IsNumber(symbol) || symbol == 92)
}

func StringRepack(s string) string {
	var result []string

	var letterToPush string
	var isEscape bool

	for _, symbol := range s {
		if isEscape == false && symbol == 92 { // int32 code for \ symbol
			isEscape = true
			continue
		}
		if unicode.IsLetter(symbol) || isEscapedNumberOrBackslash(isEscape, symbol) {
			letterToPush = string(symbol)
			result = append(result, letterToPush)
			isEscape = false
			continue
		}
		if unicode.IsNumber(symbol) && isEscape == false {
			for i := 0; i < int(symbol-'1'); i++ {
				result = append(result, letterToPush)
			}
		}
	}

	return strings.Join(result, "")
}
