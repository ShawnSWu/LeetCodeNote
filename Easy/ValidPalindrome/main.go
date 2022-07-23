package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	word := handleNonCharacters(s)

	mid := len(word) / 2

	var font string
	var back string

	if isEvenLength(word) {
		font = word[0:mid]
		back = word[mid:len(word)]
	} else {
		font = word[0:mid]
		back = word[mid+1 : len(word)]
	}

	backP := len(back) - 1

	for i := 0; i < len(font); i++ {
		if font[i] != back[backP-i] {
			return false
		}
	}

	return true
}

//word是否是偶數長度
func isEvenLength(s string) bool {
	return len(s)%2 == 0
}

func handleNonCharacters(s string) string {
	r := regexp.MustCompile("\\w")
	charList := r.FindAllString(s, -1)

	var word = ""
	for c := range charList {
		word += charList[c]
	}
	word = strings.ReplaceAll(word, "_", "")
	return strings.ToLower(word)
}
