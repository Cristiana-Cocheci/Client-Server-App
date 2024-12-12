package src

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func SolveEx5(args []string) string {
	message := ""
	var list []string
	if len(args) != int(conf.ArrayLength) {
		return "Error: Incorrect number of arguments provided, configuration requires " + strconv.FormatInt(conf.ArrayLength, 10) + " arguments\n"
	}
	for i := 0; i < int(conf.ArrayLength); i++ {
		list = append(list, "")
	}
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
	return message + strings.Join(list[:len(args)], " ")
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
