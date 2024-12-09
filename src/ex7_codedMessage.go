package src

import (
	"strings"
	"sync"
)

func SolveEx7(word string) string {
	var res [100]string
	var number = 0
	wg := sync.WaitGroup{}
	cnt := 0
	for _, c := range word {
		if c >= '0' && c <= '9' {
			number = number*10 + int(c-'0')
		} else {
			wg.Add(1)
			go func(n int, i int) {
				defer wg.Done()
				res[i] = ""
				for j := 0; j < n; j++ {
					res[i] += string(c)
				}
			}(number, cnt)
			number = 0
			cnt++
		}
	}
	wg.Wait()
	return strings.Join(res[:cnt], "")
}
