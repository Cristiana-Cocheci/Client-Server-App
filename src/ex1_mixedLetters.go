package src

import (
	"strconv"
	"strings"
)

func processPosition(index int, words []string, result chan string) {
	output := strconv.Itoa(index + 1)
	for _, word := range words {
		output += string(word[index])
	}
	result <- output
}

func MixedLetters(wordList []string) string {
	if len(wordList) == 0 {
		// http.Error(w, "No words provided", http.StatusBadRequest)
		return ""
	}

	length := len(wordList[0])
	for _, word := range wordList {
		if len(word) != length {
			// http.Error(w, "All words must be of the same length", http.StatusBadRequest)
			return ""
		}
	}

	result := make(chan string, length)

	for i := 0; i < length; i++ {
		go processPosition(i, wordList, result)
	}

	output := make([]string, length)
	for i := 0; i < length; i++ {
		word := <-result
		index, _ := strconv.ParseInt(word[:1], 10, 64)
		output[index-1] = word[1:]
	}

	return (strings.Join(output, ", "))
}
