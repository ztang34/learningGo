package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := `Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	tokens := strings.Fields(text)
	count := make(map[string]int)

	for _, token := range tokens {
		count[strings.ToLower(token)]++
	}

	var sortedWords []string

	for key := range count {
		sortedWords = append(sortedWords, key)
	}

	sort.Strings(sortedWords)

	for _, word := range sortedWords {
		fmt.Printf("%s appeard %d in text\n", word, count[word])
	}

}
