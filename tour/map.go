package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word_map := make(map[string]int)
	for _, word := range strings.Fields(s) {
		_, contained := word_map[word]
		if contained {
			word_map[word] += 1
		} else {
			word_map[word] = 1
		}
	}
	return word_map
}

func main() {
	wc.Test(WordCount)
}
