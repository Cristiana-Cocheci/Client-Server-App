package src

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func SolveEx6(args []string) string {
	k := rand.IntN(26)
	var dir string
	if rand.Float64() > 0.5 {
		dir = "LEFT"
	} else {
		dir = "RIGHT"
	}
	message := fmt.Sprintf("Key: %d, Direction: %s => ", k, dir)

	var resList [10]string
	wg := sync.WaitGroup{}
	for i, arg := range args {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resList[i] = caesar(arg, dir, k)
		}()
	}
	wg.Wait()

	return message + fmt.Sprint(resList[:len(args)])
}

func caesar(s string, dir string, k int) string {
	var factor rune
	switch dir {
	case "LEFT":
		factor = -1
	case "RIGHT":
		factor = 1
	}
	news := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			news += string((c-'A'+26+factor*rune(k))%26 + 'A')
		} else if c >= 'a' && c <= 'z' {
			news += string((c-'a'+26+factor*rune(k))%26 + 'a')
		} else {
			news += string(c)
		}
	}
	return news
}
