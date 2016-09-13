package main

import (
	"fmt"
	"hackerrank/utils"
)

func main() {
	var numCases int
	fmt.Scanln(&numCases)

	cases := make([][][]int, numCases)

	for i := 0; i < numCases; i++ {
		cases[i] = make([][]int, 2)
		cases[i][0] = make([]int, 2)

		utils.ScanIntoArray(cases[i][0])

		cases[i][1] = make([]int, cases[i][0][0])
		utils.ScanIntoArray(cases[i][1])
	}

	for i := 0; i < numCases; i++ {
		numStudents := cases[i][0][0]
		threshold := cases[i][0][1]

		numAttending := 0
		holdingClass := false
		for x := 0; x < numStudents; x++ {
			if cases[i][1][x] <= 0 {
				numAttending++
			}

			if numAttending >= threshold {
				holdingClass = true
				break
			}
		}

		if holdingClass {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
