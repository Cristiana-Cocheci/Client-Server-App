package src

import (
	"fmt"
	"strings"
	"unicode"
)

func processWord6(word string) bool {
	l := len(word)
	if !unicode.IsUpper(rune(word[0])) || !unicode.IsUpper(rune(word[l-1])) {
		return false
	}
	cnt := 0
	for i := 1; i < l-1; i++ {
		if !unicode.IsUpper(rune(word[i])) {
			cnt++
		}
	}
	return cnt%2 == 0
}

func process6(row string, resultChan chan float32) {
	var cnt float32 = 0
	words := strings.Split(row, "; ")
	for _, word := range words {
		if processWord6(word) {
			cnt++
		}
	}
	resultChan <- cnt / float32(len(words))
}

func mapReduce6(args []string) string {
	resultChan := make(chan float32, len(args))
	for _, row := range args {
		go process6(row, resultChan)
	}
	var res float32 = 0
	for i := 0; i < len(args); i++ {
		res += <-resultChan
	}
	res = res / float32(len(args))
	return fmt.Sprintf("Average number of words : %f\n", res)
}
