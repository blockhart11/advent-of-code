package _15

import "fmt"

func complicateThings(floormap [][]int) int {
	oldSize := len(floormap)
	newSize := oldSize * 5
	biggerFloormap := make([][]int, newSize)
	for i, _ := range biggerFloormap {
		biggerFloormap[i] = make([]int, newSize)
	}
	// fill out the new bigger map
	for i, v := range floormap {
		for j, w := range v {
			// col 0
			biggerFloormap[i][j] = w
			biggerFloormap[i+oldSize][j] = w + 1
			biggerFloormap[i+oldSize*2][j] = w + 2
			biggerFloormap[i+oldSize*3][j] = w + 3
			biggerFloormap[i+oldSize*4][j] = w + 4

			// col 1
			biggerFloormap[i][j+oldSize] = w + 1
			biggerFloormap[i+oldSize][j+oldSize] = w + 2
			biggerFloormap[i+oldSize*2][j+oldSize] = w + 3
			biggerFloormap[i+oldSize*3][j+oldSize] = w + 4
			biggerFloormap[i+oldSize*4][j+oldSize] = w + 5
			// col 2
			biggerFloormap[i][j+oldSize*2] = w + 2
			biggerFloormap[i+oldSize][j+oldSize*2] = w + 3
			biggerFloormap[i+oldSize*2][j+oldSize*2] = w + 4
			biggerFloormap[i+oldSize*3][j+oldSize*2] = w + 5
			biggerFloormap[i+oldSize*4][j+oldSize*2] = w + 6
			// col 3
			biggerFloormap[i][j+oldSize*3] = w + 3
			biggerFloormap[i+oldSize][j+oldSize*3] = w + 4
			biggerFloormap[i+oldSize*2][j+oldSize*3] = w + 5
			biggerFloormap[i+oldSize*3][j+oldSize*3] = w + 6
			biggerFloormap[i+oldSize*4][j+oldSize*3] = w + 7
			// col 4
			biggerFloormap[i][j+oldSize*4] = w + 4
			biggerFloormap[i+oldSize][j+oldSize*4] = w + 5
			biggerFloormap[i+oldSize*2][j+oldSize*4] = w + 6
			biggerFloormap[i+oldSize*3][j+oldSize*4] = w + 7
			biggerFloormap[i+oldSize*4][j+oldSize*4] = w + 8
		}
	}
	// normalize the big map
	for i, v := range biggerFloormap {
		for j, w := range v {
			if w > 9 {
				biggerFloormap[i][j] -= 9
			}
		}
	}

	//for _, v := range biggerFloormap {
	//	for _, w := range v {
	//		fmt.Printf("%d", w)
	//	}
	//	fmt.Printf("\n")
	//}
	return shortestPathCost(biggerFloormap)
}

func shortestPathCost(floorMap [][]int) int {
	// assume square
	size := len(floorMap)
	bestCostMap := make([][]int, size)
	for i, _ := range bestCostMap {
		bestCostMap[i] = make([]int, size)
	}

	costLimit := approximateShortestPath(floorMap, 0, size, 0, 0)
	fmt.Println("Limit:", costLimit)
	findShortestPath(floorMap, bestCostMap, size, 0, 0, costLimit)
	return bestCostMap[size-1][size-1]
}

func approximateShortestPath(floorMap [][]int, cost, size, xPos, yPos int) int {
	// are we done?
	if xPos == size-1 && yPos == size-1 {
		return cost
	}

	downCost, rightCost := 10, 10

	// can we go right?
	if xPos < size-1 {
		rightCost = floorMap[xPos+1][yPos]
	}
	// how about down?
	if yPos < size-1 {
		downCost = floorMap[xPos][yPos+1]
	}
	if downCost < rightCost {
		// just go down
		return approximateShortestPath(floorMap, cost+downCost, size, xPos, yPos+1)
	} else {
		return approximateShortestPath(floorMap, cost+rightCost, size, xPos+1, yPos)
	}

}

func findShortestPath(floorMap [][]int, bestCostMap [][]int, size, xPos, yPos, limit int) {

	currCost := bestCostMap[xPos][yPos]

	// are we done?
	if (xPos == size-1 && yPos == size-1) || currCost >= limit {
		return
	}

	var bestMoveCost, nextMoveCost int
	// try left
	if xPos > 0 {
		nextX, nextY := xPos-1, yPos
		bestMoveCost = bestCostMap[nextX][nextY]
		nextMoveCost = floorMap[nextX][nextY] + currCost
		if bestMoveCost == 0 || bestMoveCost > nextMoveCost {
			bestCostMap[nextX][nextY] = nextMoveCost
			findShortestPath(floorMap, bestCostMap, size, nextX, nextY, limit)
		}
	}
	// try right
	if xPos < size-1 {
		nextX, nextY := xPos+1, yPos
		bestMoveCost = bestCostMap[nextX][nextY]
		nextMoveCost = floorMap[nextX][nextY] + currCost
		if bestMoveCost == 0 || bestMoveCost > nextMoveCost {
			bestCostMap[nextX][nextY] = nextMoveCost
			findShortestPath(floorMap, bestCostMap, size, nextX, nextY, limit)
		}
	}
	// try up
	if yPos > 0 {
		nextX, nextY := xPos, yPos-1
		bestMoveCost = bestCostMap[nextX][nextY]
		nextMoveCost = floorMap[nextX][nextY] + currCost
		if bestMoveCost == 0 || bestMoveCost > nextMoveCost {
			bestCostMap[nextX][nextY] = nextMoveCost
			findShortestPath(floorMap, bestCostMap, size, nextX, nextY, limit)
		}
	}
	// try down
	if yPos < size-1 {
		nextX, nextY := xPos, yPos+1
		bestMoveCost = bestCostMap[nextX][nextY]
		nextMoveCost = floorMap[nextX][nextY] + currCost
		if bestMoveCost == 0 || bestMoveCost > nextMoveCost {
			bestCostMap[nextX][nextY] = nextMoveCost
			findShortestPath(floorMap, bestCostMap, size, nextX, nextY, limit)
		}
	}
}
