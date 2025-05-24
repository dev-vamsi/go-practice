package main

import (
	"fmt"
	"strings"
)

type TextObject struct {
	text string
}

func (t *TextObject) toLowerCase() string {
	return strings.ToLower(t.text)
}

func (t *TextObject) toUpperCase() string {
	return strings.ToUpper(t.text)
}

func (t *TextObject) printWordCount() {
	wordCount := map[string]int{}
	for _, val := range strings.Split(t.text, " ") {
		wordCount[val] = wordCount[val] + 1
	}
	fmt.Println(wordCount)
}

func (t *TextObject) printCharacterCount() {
	charCount := map[rune]int{}
	var mostRepeatedChar rune
	for _, val := range strings.Split(t.text, "") {
		_var := []rune(val)[0]
		charCount[_var] = charCount[_var] + 1
		if charCount[_var] > charCount[mostRepeatedChar] {
			mostRepeatedChar = _var
		}
	}
	fmt.Println(charCount)
	fmt.Println(string(mostRepeatedChar))
}

func (t *TextObject) getReversedString() string {
	len := len(t.text)
	res := make([]rune, len)
	for idx, val := range strings.Split(t.text, "") {
		res[len-idx-1] = []rune(val)[0]
	}
	return string(res)
}

func (t *TextObject) isPalindrome() bool {
	return t.text == t.getReversedString()
}
