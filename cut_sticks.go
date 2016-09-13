package main

import (
	"fmt"
	"hackerrank/utils"
	"sort"
)

func main() {
	numInput := 0
	fmt.Scanln(&numInput)

	inputs := make([]int, numInput)
	utils.ScanIntoArray(inputs)

	sort.Ints(inputs)

	for curIndex := 0; curIndex < numInput; {
		fmt.Println(numInput - curIndex)

		delVal := inputs[curIndex]
		for i := curIndex; i < numInput; i++ {
			inputs[i] = inputs[i] - delVal

			if inputs[i] <= 0 {
				curIndex++
			}
		}
	}
}
