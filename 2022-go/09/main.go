package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections"
)

func main() {
	//input, err := os.ReadFile("./2022-go/09/sample.txt")
	//input, err := os.ReadFile("./2022-go/09/sample2.txt")
	input, err := os.ReadFile("./2022-go/09/input.txt")
	if err != nil {
		panic(err)
	}

	const size = 1000
	// initialize the rope thingy. Try 1000x1000 as the size, whatev
	bridge := make([][]bool, size)
	for i, _ := range bridge {
		bridge[i] = make([]bool, size)
	}
	// Part 1
	//start := collections.Point{X: size / 2, Y: size / 2}
	//head, tail := start, start
	// Part 2
	ropeKnots := make([]*collections.Point, 10, 10)
	for i, _ := range ropeKnots {
		ropeKnots[i] = &collections.Point{X: size / 2, Y: size / 2}
	}

	// process all moves
	for _, line := range strings.Split(string(input), "\n") {
		cmd := strings.Split(line, " ")
		amount, _ := strconv.Atoi(cmd[1])
		for i := 0; i < amount; i++ {
			moveHead(ropeKnots[0], cmd[0])
			// Part 1
			//	moveTailTowardHead(&tail, head)
			//	bridge[tail.X][tail.Y] = true

			// Part 2
			for j := 1; j < len(ropeKnots); j++ {
				moveTailTowardHead(ropeKnots[j], *ropeKnots[j-1])
			}
			bridge[ropeKnots[9].X][ropeKnots[9].Y] = true
		}

	}

	// count up tail visits
	var count int
	for _, v := range bridge {
		for _, visited := range v {
			if visited {
				count++
			}
		}
	}
	fmt.Println(count)
}

func moveHead(from *collections.Point, direction string) {
	switch direction {
	case "L":
		from.X--
	case "U":
		from.Y--
	case "R":
		from.X++
	case "D":
		from.Y++
	default:
		panic("uhhh which direction is this")
	}
}

func moveTailTowardHead(from *collections.Point, target collections.Point) {
	towards := collections.Point{
		X: target.X - from.X,
		Y: target.Y - from.Y,
	}
	switch {
	case towards.X == 0:
		// same vertical plane
		switch towards.Y {
		case 2: // head is directly below
			from.Y++
		case -2: // head is directly above
			from.Y--
		}
	case towards.X == 1:
		switch towards.Y {
		case 2:
			// head is diagonally below and right
			from.X++
			from.Y++
		case -2:
			// head is diagonally above and right
			from.X++
			from.Y--
		}
	case towards.X == 2:
		switch towards.Y {
		case 1, 2:
			// head is diagonally below and right
			from.X++
			from.Y++
		case -1, -2:
			// head is diagonally above and right
			from.X++
			from.Y--
		case 0:
			// head is directly right
			from.X++
		}
	case towards.X == -1:
		switch towards.Y {
		case 2:
			// head is diagonally below and left
			from.X--
			from.Y++
		case -2:
			// head is diagonally above and left
			from.X--
			from.Y--
		}
	case towards.X == -2:
		switch towards.Y {
		case 1, 2:
			// head is diagonally below and left
			from.X--
			from.Y++
		case -1, -2:
			// head is diagonally above and left
			from.X--
			from.Y--
		case 0:
			// head is directly left
			from.X--
		}
	}

}
