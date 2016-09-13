package main

import (
	"fmt"
	"hackerrank/utils"
)

func main() {
	stackSums := make([]int, 3)
	stackSizes := make([]int, 3)
	stackIndices := make([]int, 3)
	utils.ScanIntoArray(stackSizes)

	stacks := make([][]int, 3)
	for i := 0; i < 3; i++ {
		stacks[i] = make([]int, stackSizes[i])
		utils.ScanIntoArray(stacks[i])

		for x := 0; x < stackSizes[i]; x++ {
			stackSums[i] += stacks[i][x]
		}
	}

	equalsFound := false
	for {
		if stackSums[0] < 1 || stackSums[1] < 1 || stackSums[2] < 1 {
			break
		}

		if stackSums[0] == stackSums[1] && stackSums[0] == stackSums[2] {
			equalsFound = true
			break
		}

		tallerStack := 0
		for i := 1; i < 3; i++ {
			if stackSums[i] > stackSums[tallerStack] {
				tallerStack = i
			}
		}

		indexToRemove := stackIndices[tallerStack]
		stackSums[tallerStack] = stackSums[tallerStack] - stacks[tallerStack][indexToRemove]
		stackIndices[tallerStack] = stackIndices[tallerStack] + 1
	}

	if equalsFound {
		fmt.Println(stackSums[0])
	} else {
		fmt.Println(0)
	}
}
