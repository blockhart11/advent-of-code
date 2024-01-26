package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
	"github.com/golang-collections/collections"
)

// Expected output
//  sample.txt: 4361, 467835

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Parse the file
	lines := strings.Split(string(input), "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for j, char := range line {
			grid[i][j] = byte(char)
		}
	}
	var partNumbers []int
	type gear struct {
		count int
		val   int
	}
	gears := make(map[string]*gear)
	// Part 1
	for i, line := range grid {
		for j := 0; j < len(line); j++ {
			char := grid[i][j]
			if !isDigit(char) {
				continue
			}

			// parse the number
			var val int
			var foundSymbol bool
			var gearLoc *collections.Point
			for {
				// does this border a symbol?
				row, col := bordersSymbol(grid, i, j)
				if row >= 0 || col >= 0 {
					foundSymbol = true
					gearLoc = &collections.Point{
						X: col,
						Y: row,
					}
				}
				// collect value
				thisDigit, err := strconv.Atoi(string(char))
				if err != nil {
					panic(string(char) + "not a number")
				}
				val = (val * 10) + thisDigit
				j++
				if j >= len(line) {
					// end of the line
					break
				}
				if !isDigit(line[j]) {
					break
				}
				char = grid[i][j]
			}
			if foundSymbol {
				partNumbers = append(partNumbers, val)
				if gearLoc != nil {
					key := fmt.Sprintf(`%d,%d`, gearLoc.Y, gearLoc.X)
					g, ok := gears[key]
					switch {
					case !ok:
						gears[key] = &gear{
							count: 1,
							val:   val,
						}
					case g.count == 1:
						g.val *= val
						g.count++
					default:
						delete(gears, key)
					}
				}
			}
		}
	}
	fmt.Println(utils.Fold(partNumbers, 0, func(current, next int) int {
		return current + next
	}))
	gearSum := 0
	for _, g := range gears {
		if g.count == 2 {
			gearSum += g.val
		}
	}
	fmt.Println(gearSum)
}

func bordersSymbol(grid [][]byte, row, col int) (int, int) {
	// look up
	if row > 0 && isSymbol(grid[row-1][col]) {
		return row - 1, col
	}
	// look down
	if row < len(grid)-1 && isSymbol(grid[row+1][col]) {
		return row + 1, col
	}
	// look left
	if col > 0 && isSymbol(grid[row][col-1]) {
		return row, col - 1
	}
	// look right
	if col < len(grid[row])-1 && isSymbol(grid[row][col+1]) {
		return row, col + 1
	}
	// look up-left
	if row > 0 && col > 0 && isSymbol(grid[row-1][col-1]) {
		return row - 1, col - 1
	}
	// look up-right
	if row > 0 && col < len(grid[row])-1 && isSymbol(grid[row-1][col+1]) {
		return row - 1, col + 1
	}
	// look down-left
	if row < len(grid)-1 && col > 0 && isSymbol(grid[row+1][col-1]) {
		return row + 1, col - 1
	}
	// look down-right
	if row < len(grid)-1 && col < len(grid[row])-1 && isSymbol(grid[row+1][col+1]) {
		return row + 1, col + 1
	}
	return -1, -1
}

func isSymbol(input byte) bool {
	return input != '.' && !isDigit(input)
}

func isDigit(input byte) bool {
	return '0' <= input && input <= '9'
}
