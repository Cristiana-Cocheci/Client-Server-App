package src

import (
	"fmt"
	"sync"
)

func doubleFirstDigit(n int) int {
	cn := n
	p := 1
	uc := 0
	for n > 0 {
		uc = n % 10
		n /= 10
		p *= 10
	}
	return uc*p + cn
}

func solveEx12(numList []int) int {
	s := 0
	wg := sync.WaitGroup{}
	wg.Add(len(numList))
	for _, n := range numList {
		go func() {
			defer wg.Done()
			s += doubleFirstDigit(n)
		}()
	}
	wg.Wait()
	return s
}

func SolveEx12(args []string) string {
	numList := []int{}
	for _, arg := range args {
		num := 0
		fmt.Sscanf(arg, "%d", &num)
		numList = append(numList, num)
	}
	return fmt.Sprint(solveEx12(numList))
}
