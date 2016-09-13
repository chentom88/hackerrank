package main

import (
	"fmt"
	"hackerrank/utils"
	"sort"
)

type gameBoard struct {
	featurePositions []int
	features         map[int]int
	knownPaths       map[int]int
	shortestPath     int
}

var MAX_PATHS int = 1000000

func main() {
	numBoards := 0
	fmt.Scanln(&numBoards)

	boards := make([]*gameBoard, numBoards)
	for i := 0; i < numBoards; i++ {
		boards[i] = &gameBoard{
			features:     make(map[int]int),
			knownPaths:   make(map[int]int),
			shortestPath: MAX_PATHS,
		}

		numLadders := 0
		fmt.Scanln(&numLadders)

		ladder := make([]int, 2)
		for x := 0; x < numLadders; x++ {
			utils.ScanIntoArray(ladder)
			boards[i].featurePositions = append(boards[i].featurePositions, ladder[0])
			boards[i].features[ladder[0]] = ladder[1]
		}

		numSnakes := 0
		fmt.Scanln(&numSnakes)

		snake := make([]int, 2)
		for x := 0; x < numSnakes; x++ {
			utils.ScanIntoArray(snake)
			boards[i].featurePositions = append(boards[i].featurePositions, snake[0])
			boards[i].features[snake[0]] = snake[1]
		}

		boards[i].featurePositions = append(boards[i].featurePositions, 1)
		boards[i].featurePositions = append(boards[i].featurePositions, 100)
		sort.Ints(boards[i].featurePositions)
	}

	for i := 0; i < numBoards; i++ {
		boards[i].findShortest(1, 0, boards[i].featurePositions[1:], make(map[int]bool))

		if boards[i].shortestPath < MAX_PATHS {
			fmt.Println(boards[i].shortestPath)
		} else {
			fmt.Println(-1)
		}
	}
}

func (b *gameBoard) findShortest(curPos int, runSum int, nextNodes []int, visited map[int]bool) {
	if _, ok := visited[curPos]; ok {
		return
	}

	if runSum >= b.shortestPath {
		return
	}

	newVisited := addToVisited(visited, curPos)

	// traverse the snake or ladder to get at new location
	var newNext []int
	if curPos > 1 {
		curPos = b.features[curPos]

		if curPos == 100 {
			if runSum < b.shortestPath {
				b.shortestPath = runSum
			}

			return
		}

		nextFeatIndex := sort.Search(len(b.featurePositions), func(i int) bool { return b.featurePositions[i] > curPos })
		newNext = b.featurePositions[nextFeatIndex:]
	} else {
		newNext = nextNodes
	}

	numNext := len(newNext)
	for i := 0; i < numNext; i++ {
		nextPos := newNext[i]

		if _, ok := visited[nextPos]; ok {
			continue
		}

		nextCost := b.findPathNum(curPos, nextPos)
		if nextCost < 0 {
			continue
		}

		nextSum := runSum + nextCost
		if nextPos != 100 {
			b.findShortest(nextPos, nextSum, newNext[i+1:], newVisited)
		} else if nextSum < b.shortestPath {
			b.shortestPath = nextSum
		}
	}
}

func (b *gameBoard) findPathNum(start int, end int) int {
	key := (100 * start) + end

	value, ok := b.knownPaths[key]

	if ok {
		return value
	}

	value = 0
	for current := start; current < end; {
		current += 6
		value++

		// Assumes there will never be six consecutive features
		if current != end {
			for i := 0; i < 6; i++ {
				if i == 5 {
					// back at original spot, no way past!
					value = -1
					b.knownPaths[key] = value
					return value
				}

				if _, ok := b.features[current]; !ok {
					break
				} else {
					current -= 1
				}
			}
		}
	}

	b.knownPaths[key] = value
	return value
}

// I hate this, must refactor
func addToVisited(orig map[int]bool, newValue int) map[int]bool {
	lenOrig := len(orig)

	newMap := make(map[int]bool, lenOrig+1)

	for k, v := range orig {
		newMap[k] = v
	}

	newMap[newValue] = true

	return newMap
}
