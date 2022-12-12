package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/golang-collections/collections"
)

func main() {
	//input, err := os.ReadFile("./2022-go/12/sample.txt")
	input, err := os.ReadFile("./2022-go/12/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	var start, end collections.Point
	hill := make([][]*node, len(lines))
	for row, line := range lines {
		hill[row] = make([]*node, len(line))
		for col, letter := range line {
			switch letter {
			case 'S':
				start.Y = row
				start.X = col
				hill[row][col] = &node{'a', math.MaxInt}
			case 'E':
				end.Y = row
				end.X = col
				hill[row][col] = &node{'z', math.MaxInt}
			default:
				hill[row][col] = &node{letter, math.MaxInt}
			}
		}
	}

	// part 1
	shortest := shortestPath(hill, start, end, 0)
	fmt.Println(shortest)

	// part 2
	cost := math.MaxInt
	for row, rows := range hill {
		for col, _ := range rows {
			if hill[row][col].height == 'a' {
				pathCost := shortestPath(hill, collections.Point{
					X: col,
					Y: row,
				},
					end,
					0)
				if pathCost < cost {
					cost = pathCost
				}
			}
		}
	}
	fmt.Println(cost)
}

type node struct {
	height        rune
	costFromStart int
}

func shortestPath(hill [][]*node, start, end collections.Point, accum int) int {
	// for convenience
	switch {
	case start == end:
		return accum
	case start.X == -1, start.Y == -1, start.X >= len(hill[0]), start.Y >= len(hill):
		// out of bounds
		return math.MaxInt
	case hill[start.Y][start.X].costFromStart <= accum:
		// another path got here faster
		return math.MaxInt
	}
	n := hill[start.Y][start.X]
	shortest := math.MaxInt
	n.costFromStart = accum
	// check left
	if start.X > 0 && hill[start.Y][start.X-1].height <= n.height+1 {
		shortest = shortestPath(
			hill,
			collections.Point{X: start.X - 1, Y: start.Y},
			end,
			accum+1,
		)
	}
	// check up
	if start.Y > 0 && hill[start.Y-1][start.X].height <= n.height+1 {
		up := shortestPath(
			hill,
			collections.Point{X: start.X, Y: start.Y - 1},
			end,
			accum+1,
		)
		if up < shortest {
			shortest = up
		}
	}
	// check right
	if start.X < len(hill[0])-1 && hill[start.Y][start.X+1].height <= n.height+1 {
		right := shortestPath(
			hill,
			collections.Point{X: start.X + 1, Y: start.Y},
			end,
			accum+1,
		)
		if right < shortest {
			shortest = right
		}
	}
	// check down
	if start.Y < len(hill)-1 && hill[start.Y+1][start.X].height <= n.height+1 {
		down := shortestPath(
			hill,
			collections.Point{X: start.X, Y: start.Y + 1},
			end,
			accum+1,
		)
		if down < shortest {
			shortest = down
		}
	}

	return shortest
}
