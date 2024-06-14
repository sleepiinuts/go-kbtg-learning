package funcs

import (
	"strings"
)

var trf = map[rune]string{
	'0': "zero",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
	' ': "",
}

func TrimSpaceAndConvInt(str string) string {
	var out string

	str = strings.ToLower(str)

	for _, r := range str {
		if _, ok := trf[r]; ok {
			out += trf[r]
			continue
		}
		out += string(r)
	}

	return out
}
