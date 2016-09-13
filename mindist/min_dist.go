package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type distInfo struct {
	LastIndex   int
	MinDistance int
}

func main() {
	seed, _ := strconv.Atoi(os.Args[1])
	arraySize, _ := strconv.Atoi(os.Args[2])

	var values []int64
	values = make([]int64, arraySize, arraySize)

	rand.Seed(int64(seed))

	populateVals(values)

	fmt.Println(findMinDistance(len(values), values))
}

func populateVals(arr []int64) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int63n(1000)
	}
}

func printSlice(arr []int64) {
	fmt.Println(arr)
}

func findMinDistance(numValues int, values []int64) int {
	maxDist := numValues + 1
	seenValues := make(map[int64]*distInfo)

	for i, val := range values {
		dist, ok := seenValues[val]

		if ok {
			tempDist := i - dist.LastIndex

			if tempDist < dist.MinDistance {
				dist.MinDistance = tempDist
			}

			dist.LastIndex = i
			continue
		}

		seenValues[val] = &distInfo{
			LastIndex:   i,
			MinDistance: maxDist,
		}
	}

	overallMinDist := maxDist
	for _, dist := range seenValues {
		if dist.MinDistance < overallMinDist {
			overallMinDist = dist.MinDistance
		}
	}

	if overallMinDist == maxDist {
		overallMinDist = -1
	}

	return overallMinDist
}
