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
		return "Error: No words provided"
	}

	if len(wordList) != int(conf.ArrayLength) {
		return "Error: Incorrect number of arguments provided, configuration requires " + strconv.FormatInt(conf.ArrayLength, 10) + " arguments\n"
	}

	length := len(wordList[0])
	for _, word := range wordList {
		if len(word) != length {
			return "Error: All words must be of the same length"
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
