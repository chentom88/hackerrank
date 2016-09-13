package main

import (
	"fmt"
	"hackerrank/utils"
)

func main() {
	input := make([]int, 4)
	utils.ScanIntoArray(input)

	numerator := input[0] - input[2]
	denominator := input[3] - input[1]

	if denominator == 0 {
		if numerator == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

	if numerator/denominator < 0 {
		fmt.Println("NO")
		return
	}

	if numerator%denominator != 0 {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
}
