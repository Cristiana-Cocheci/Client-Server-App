package src

import (
	"fmt"
	"strconv"
	"strings"
)

func processWord15(n int) bool {
	cnt := 0
	for n != 0 {
		cnt++
		n = n & (n - 1)
	}

	return cnt == 3
}

func process15(row string, resultChan chan float32) {
	var cnt float32 = 0
	words := strings.Split(row, "; ")
	for _, word := range words {
		n, _ := strconv.Atoi(word)
		if processWord15(n) {
			cnt++
		}
	}
	resultChan <- cnt
}

func mapReduce15(args []string) string {
	resultChan := make(chan float32, len(args))
	for _, row := range args {
		go process15(row, resultChan)
	}
	var res float32 = 0
	for i := 0; i < len(args); i++ {
		res += <-resultChan
	}
	res = res / float32(len(args))
	return fmt.Sprintf("Average number of words with 3 1s in binary representation: %f\n", res)
}
