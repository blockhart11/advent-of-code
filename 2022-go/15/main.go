package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
	"github.com/golang-collections/collections"
)

const (
	// sample input
	//evalRow = 10
	//gridSize = 20 // sample input

	// real input
	evalRow  = 2000000
	gridSize = 4000000
)

func main() {
	//input, err := os.ReadFile("./2022-go/15/sample.txt")
	input, err := os.ReadFile("./2022-go/15/input.txt")
	if err != nil {
		panic(err)
	}

	// populate the map with sensors and beacons
	var sensors []sensor
	var beacons []beacon // might as well track these too
	for _, line := range strings.Split(string(input), "\n") {
		coords := strings.Split(line, " ")
		b := beacon{
			X: utils.AtoiX(coords[2]),
			Y: utils.AtoiX(coords[3]),
		}
		beacons = append(beacons, b)
		s := sensor{
			loc: collections.Point{
				X: utils.AtoiX(coords[0]),
				Y: utils.AtoiX(coords[1]),
			},
			closest: &b,
		}
		s.distance = manhattanDistanceToBeacon(s)
		sensors = append(sensors, s)
	}

	// Part 1 - evaluate row 2000000
	whereCantBeaconsGo := make(map[int]bool)
	for _, s := range sensors {
		rowDist := int(math.Abs(float64(s.loc.Y - evalRow)))
		horizontalDiff := s.distance - rowDist // how far to mark horizontally
		for j := s.loc.X - horizontalDiff; j <= s.loc.X+horizontalDiff; j++ {
			whereCantBeaconsGo[j] = true
		}
	}
	fmt.Println(len(whereCantBeaconsGo))

	// Part 2 - find the beacon
	// it must be adjacent to a boundary, so just test the boundary around every sensor
	for i, s := range sensors {
		fmt.Printf("Checking boundary of sensor %d...", i)
		var testPoint collections.Point
		// upper left
		testPoint.X = s.loc.X - s.distance - 1
		testPoint.Y = s.loc.Y
		for testPoint.X <= s.loc.X {
			if testPoint.X < 0 { // skip to be in bounds
				testPoint.Y += testPoint.X
				testPoint.X = 0
			}
			if testPoint.Y < 0 { // donezo
				break
			}
			if !withinAnyRange(testPoint, sensors) {
				fmt.Printf("distress signal at %d, %d. Frequency is %d", testPoint.X, testPoint.Y, testPoint.X*4000000+testPoint.Y)
				return
			}
			testPoint.X++
			testPoint.Y--
		}
		// lower left border
		testPoint.X = s.loc.X - s.distance - 1
		testPoint.Y = s.loc.Y
		for testPoint.X <= s.loc.X {
			if testPoint.X < 0 { // skip to be in bounds
				testPoint.Y -= testPoint.X
				testPoint.X = 0
			}
			if testPoint.Y > gridSize { // donezo
				break
			}
			if !withinAnyRange(testPoint, sensors) {
				fmt.Printf("distress signal at %d, %d. Frequency is %d", testPoint.X, testPoint.Y, testPoint.X*4000000+testPoint.Y)
				return
			}
			testPoint.X++
			testPoint.Y++
		}
		// upper right border
		testPoint.X = s.loc.X + s.distance + 1
		testPoint.Y = s.loc.Y
		for testPoint.X > s.loc.X {
			if testPoint.X > gridSize { // skip to be in bounds
				testPoint.Y -= testPoint.X - gridSize
				testPoint.X = gridSize
			}
			if testPoint.Y < 0 { // donezo
				break
			}
			if !withinAnyRange(testPoint, sensors) {
				fmt.Printf("distress signal at %d, %d. Frequency is %d", testPoint.X, testPoint.Y, testPoint.X*4000000+testPoint.Y)
				return
			}
			testPoint.X--
			testPoint.Y--
		}
		// lower right border
		testPoint.X = s.loc.X + s.distance + 1
		testPoint.Y = s.loc.Y
		for testPoint.X > s.loc.X {
			if testPoint.X > gridSize { // skip to be in bounds
				testPoint.Y += testPoint.X - gridSize
				testPoint.X = gridSize
			}
			if testPoint.Y > gridSize { // donezo
				break
			}
			if !withinAnyRange(testPoint, sensors) {
				fmt.Printf("distress signal at %d, %d. Frequency is %d", testPoint.X, testPoint.Y, testPoint.X*4000000+testPoint.Y)
				return
			}
			testPoint.X--
			testPoint.Y++
		}
		fmt.Println("nope")
	}
}

type sensor struct {
	loc      collections.Point
	closest  *beacon
	distance int
}

type beacon collections.Point

func manhattanDistance(from collections.Point, to collections.Point) int {
	return int(math.Abs(float64(from.X-to.X)) + math.Abs(float64(from.Y-to.Y)))
}

func manhattanDistanceToBeacon(s sensor) int {
	return manhattanDistance(s.loc, collections.Point(*s.closest))
}

func withinAnyRange(point collections.Point, sensors []sensor) bool {
	for _, s2 := range sensors { // test if within any other sensor
		if manhattanDistance(s2.loc, point) <= s2.distance {
			return true
		}
	}
	return false
}
