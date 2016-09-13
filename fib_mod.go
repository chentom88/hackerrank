package main

import (
	"fmt"
	"hackerrank/utils"
	"math/big"
)

var memory map[int]*big.Int
var big2 = big.NewInt(2)

func main() {
	input := make([]int, 3)
	utils.ScanIntoArray(input)

	memory = make(map[int]*big.Int)
	memory[0] = big.NewInt(int64(input[0]))
	memory[1] = big.NewInt(int64(input[1]))

	term := input[2] - 1
	for i := 2; i < term; i++ {
		getTerm(i)
	}

	fmt.Println(getTerm(term))
}

func getTerm(term int) *big.Int {
	nextNum, ok := memory[term]

	if !ok {
		nextNum = big.NewInt(0)
		num1Squared := big.NewInt(0)

		num1Squared.Exp(getTerm(term-1), big2, nil)

		nextNum.Add(num1Squared, getTerm(term-2))
		memory[term] = nextNum
	}

	return nextNum
}
