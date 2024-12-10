package src

import (
	"fmt"
	"strconv"
	"sync"
)

func inverse(num int) int {
	inv := 0
	for num > 0 {
		inv = inv*10 + num%10
		num /= 10
	}
	return inv
}

func solveEx3(numList []int) int {
	n := len(numList)
	msgs := make(chan int, len(numList))
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			msgs <- inverse(numList[i])
		}()
	}
	wg.Wait()

	sum := 0
	for i := 0; i < n; i++ {
		sum += <-msgs
	}
	return sum
}

func SolveEx3(args []string) string {
	numList := []int{}
	for _, arg := range args {
		num := 0
		fmt.Sscanf(arg, "%d", &num)
		numList = append(numList, num)
	}
	if len(numList) != int(conf.ArrayLength) {
		return "Error: Incorrect number of arguments provided, configuration requires " + strconv.FormatInt(conf.ArrayLength, 10) + " arguments\n"
	}
	return fmt.Sprint(solveEx3(numList))
}
