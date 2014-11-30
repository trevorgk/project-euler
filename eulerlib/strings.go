package eulerlib

import (
	"unicode/utf8"
	"strings"
)

func RuneToInt(r rune) int {
	return int(r - '0')
}

func IsPalindrome(s string) bool {
	reverse := Reverse(s)
	return reverse == s
}

func Reverse(s string) string  {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func WordCount(verbiage string, includeWhitespace bool) int {
	if (includeWhitespace) {
		return utf8.RuneCountInString(verbiage)
	}
	sum, split := 0, strings.Split(verbiage, " ")

	for i := 0; i < len(split); i++ {
		sum += utf8.RuneCountInString(split[i])
	}

	return sum
}