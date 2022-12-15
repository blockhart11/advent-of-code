package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections"
)

func main() {
	//input, err := os.ReadFile("./2022-go/14/sample.txt")
	input, err := os.ReadFile("./2022-go/14/input.txt")
	if err != nil {
		panic(err)
	}

	const (
		//abyssDepth = 9 + 1 + 2 // sample
		width      = 700         // who cares
		abyssDepth = 164 + 1 + 2 // input part 2
	)

	// make cave
	cave := make([][]rune, abyssDepth)
	for i := range cave {
		cave[i] = make([]rune, width)
	}

	// setup rocks
	var maxDepth int
	for _, line := range strings.Split(string(input), "\n") {
		points := strings.Split(line, " -> ")
		var from collections.Point
		for i, point := range points {
			coord := strings.Split(point, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			if y > maxDepth {
				maxDepth = y
			}
			if i == 0 {
				// first point
				from.X = x
				from.Y = y
				continue
			}
			to := collections.Point{
				X: x,
				Y: y,
			}
			switch {
			case from.X > to.X:
				// left
				for col := from.X; col >= to.X; col-- {
					cave[from.Y][col] = '#'
				}
			case from.X < to.X:
				// right
				for col := from.X; col <= to.X; col++ {
					cave[from.Y][col] = '#'
				}
			case from.Y > to.Y:
				// up
				for row := from.Y; row >= to.Y; row-- {
					cave[row][from.X] = '#'
				}
			case from.Y < to.Y:
				// down
				for row := from.Y; row <= to.Y; row++ {
					cave[row][from.X] = '#'
				}
			default:
				panic("uhhhhh")
			}
			from = to
		}
	}

	// Part 2 - make the floor
	for col := 0; col < len(cave[0]); col++ {
		cave[abyssDepth-1][col] = '#'
	}

	fmt.Println(maxDepth)

	// drop sand
	for i := 1; i >= 0; i++ {
		grain := collections.Point{X: 500, Y: 0}
		for {
			// attempt to move grain
			if grain.Y == len(cave)-1 {
				// abyss
				fmt.Printf("reached the abyss on grain %d\n", i)
				return
			}
			if cave[grain.Y+1][grain.X] != '#' {
				grain.Y++
			} else if cave[grain.Y+1][grain.X-1] != '#' {
				grain.Y++
				grain.X--
			} else if cave[grain.Y+1][grain.X+1] != '#' {
				grain.Y++
				grain.X++
			} else if grain.X == 500 && grain.Y == 0 {
				fmt.Printf("grain %d cannot move from start\n", i)
				return
			} else {
				// stuck
				//fmt.Printf("grain rests at %d, %d\n", grain.X, grain.Y)
				cave[grain.Y][grain.X] = '#'
				break
			}
		}
	}
}
