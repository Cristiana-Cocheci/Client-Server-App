package src

import (
	"fmt"
	"sync"
)

func SolveEx8(args []string) string {
	numList := []int{}
	for _, arg := range args {
		num := 0
		fmt.Sscanf(arg, "%d", &num)
		numList = append(numList, num)
	}
	return fmt.Sprint(solveEx8(numList))
}

func solveEx8(numList []int) int {
	sumcif := 0
	wg := sync.WaitGroup{}
	wg.Add(len(numList))
	for _, n := range numList {
		go func() {
			defer wg.Done()
			if prim(n) {
				sumcif += nrCif(n)
			}
		}()
	}
	wg.Wait()
	return sumcif
}

func prim(n int) bool {
	if n < 2 || n%2 == 0 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nrCif(n int) int {
	c := 0
	for n > 0 {
		c++
		n /= 10
	}
	return c
}
