package main

import (
	"github.com/renta/golang-course/hw2-repeat-symbols/internal"
	"testing"
)

func Test_stringRepack(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{"", "0123456789", ""},
		{"","a4bc2d5e", "aaaabccddddde"},
		{"", "abcd","abcd"},
		{"","45",""},
		{"",`qwe\4\5`, "qwe45"},
		{"", `qwe\45`,"qwe44444"},
		{"", `qwe5\4`,"qweeeee4"},
		{"", `qwe\\5`,`qwe\\\\\`},
		{"", `q\we`,""},
		{"", "ab/c","ab/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := internal.StringRepack(tt.input); got != tt.want {
				t.Errorf("stringRepack() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}