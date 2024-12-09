package src

import (
	"fmt"
	"sync"
)

func SolveEx5(args []string) string {
	message := ""
	var list [10]string
	wg := sync.WaitGroup{}
	for i, arg := range args {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !goodString(arg) {
				return
			}
			list[i] = fmt.Sprintf("%d", binaryToDecimal(arg))
		}()
	}
	wg.Wait()
	return message + fmt.Sprint(list[:len(args)])
}

func binaryToDecimal(binary string) int {
	decimal := 0
	for i, digit := range binary {
		if digit == '1' {
			decimal += 1 << (len(binary) - i - 1)
		}
	}
	return decimal
}

func goodString(s string) bool {
	for _, c := range s {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}
