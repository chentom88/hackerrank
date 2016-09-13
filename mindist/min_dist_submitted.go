package main

import (
	"fmt"
)

type distInfo struct {
	LastIndex   int
	MinDistance int
}

func main() {
	var numValues int
	var values []int64
	var valInterfaces []interface{}

	fmt.Scanln(&numValues)

	values = make([]int64, numValues)
	valInterfaces = make([]interface{}, numValues)

	for i := 0; i < numValues; i++ {
		valInterfaces[i] = &values[i]
	}

	fmt.Scanln(valInterfaces...)

	fmt.Println(findMinDistance(numValues, values))
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
