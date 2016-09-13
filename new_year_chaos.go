package main

import (
	"fmt"
	"hackerrank/utils"
)

type countWatcher map[int]int

func main() {
	numCases := 0
	fmt.Scanln(&numCases)

	cases := make([][]int, numCases)

	for i := 0; i < numCases; i++ {
		caseSize := 0
		fmt.Scanln(&caseSize)

		cases[i] = make([]int, caseSize)
		utils.ScanIntoArray(cases[i])
	}

	for i := 0; i < numCases; i++ {
		numSwaps := findNumSwaps(cases[i])

		if numSwaps >= 0 {
			fmt.Println(numSwaps)
		} else {
			fmt.Println("Too chaotic")
		}
	}
}

func findNumSwaps(array []int) int {
	arrayLen := len(array)

	switchCount := make(countWatcher, arrayLen)
	sumSwaps := 0

	for sorted := false; sorted == false; {
		numSwaps := 0

		for x := 0; x < arrayLen-1; x++ {
			if array[x] > array[x+1] {
				numSwaps++

				xValue := array[x]
				x1Value := array[x+1]
				array[x] = x1Value
				array[x+1] = xValue

				if switchCount.incrementAndCheck(xValue) {
					return -1
				}
			}
		}

		if numSwaps < 1 {
			sorted = true
		}

		sumSwaps += numSwaps
	}

	return sumSwaps
}

func (c countWatcher) incrementAndCheck(index int) bool {
	c[index] = c[index] + 1

	return c[index] > 2
}
