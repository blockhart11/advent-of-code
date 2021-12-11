package _11

func cyclesUntilSynced(oGrid [][]int) int {
	var cycles int
	for {
		cycleOnce(oGrid)
		cycles++
		if allFlashed(oGrid) {
			return cycles
		}
	}
}

func cycleN(oGrid [][]int, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		flashes := cycleOnce(oGrid)
		//fmt.Printf("Cycle %d: %d flashes\n", i, flashes)
		sum += flashes
	}
	return sum
}

func cycleOnce(oGrid [][]int) int {
	var flashCount int
	incrementGrid(oGrid)
	for i, row := range oGrid {
		for j, _ := range row {
			flashCount += countFlashes(oGrid, i, j, false)
		}
	}
	resetFlashed(oGrid)
	return flashCount
}

func incrementGrid(oGrid [][]int) {
	for i, row := range oGrid {
		for j, _ := range row {
			oGrid[i][j]++
		}
	}
}

func countFlashes(oGrid [][]int, row, col int, flashed bool) int {
	// out of bounds?
	if row < 0 || col < 0 || row >= len(oGrid) || col >= len(oGrid[0]) {
		return 0
	}
	// If coming from a flash, increment
	if flashed && oGrid[row][col] < 10 {
		oGrid[row][col]++
	}
	// FLASH
	if oGrid[row][col] == 10 {
		oGrid[row][col]++
		var result int
		result += countFlashes(oGrid, row-1, col-1, true)
		result += countFlashes(oGrid, row-1, col, true)
		result += countFlashes(oGrid, row-1, col+1, true)
		result += countFlashes(oGrid, row, col-1, true)
		result += countFlashes(oGrid, row, col+1, true)
		result += countFlashes(oGrid, row+1, col-1, true)
		result += countFlashes(oGrid, row+1, col, true)
		result += countFlashes(oGrid, row+1, col+1, true)
		return result + 1
	}
	// else do nothing
	return 0
}

func resetFlashed(oGrid [][]int) {
	for i, row := range oGrid {
		for j, _ := range row {
			if oGrid[i][j] > 9 {
				oGrid[i][j] = 0
			}
		}
	}
}

func allFlashed(oGrid [][]int) bool {
	for i, row := range oGrid {
		for j, _ := range row {
			if oGrid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}