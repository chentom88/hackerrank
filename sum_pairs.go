package main

import (
	"fmt"
	"hackerrank/utils"
)

func main() {
	input := make([]int, 2)
	utils.ScanIntoArray(input)

	array := make([]int, input[0])
	utils.ScanIntoArray(array)

	sumPairs := 0
	for i := 0; i < input[0]; i++ {
		for j := i + 1; j < input[0]; j++ {
			if (array[i]+array[j])%input[1] == 0 {
				sumPairs++
			}
		}
	}

	fmt.Println(sumPairs)
}
