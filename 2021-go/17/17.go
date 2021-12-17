package _17

import "fmt"

type targetArea struct {
	min, max vector2
}

type vector2 struct {
	x, y int
}

func countValidLaunches(target targetArea) int {
	// Set some initial boundaries
	ceil := vector2{target.max.x + 1, -(target.min.y + 1)}
	floor := vector2{findFloorX(target.min.x), target.min.y - 1}

	var result int
	// pass 1: brute force all potential options
	for x := floor.x; x <= ceil.x; x++ {
		for y := floor.y; y <= ceil.y; y++ {
			if launchHits(vector2{x, y}, target) > 0 {
				fmt.Println("Hit: ", x, y)
				result++
			}
		}
	}
	return result
}

func launchHits(vel vector2, t targetArea) int {
	pos := vector2{0, 0}
	// no point in checking 0,0 - let's move
	moveAndAdjustVelocity(&pos, &vel)
	steps := 1
	for {
		if pos.x > t.max.x || pos.y < t.min.y {
			return -1
		} else if isInTarget(pos, t) {
			return steps
		} else {
			moveAndAdjustVelocity(&pos, &vel)
			steps++
		}
	}
}

func moveAndAdjustVelocity(pos *vector2, vel *vector2) {
	pos.x = pos.x + vel.x
	pos.y = pos.y + vel.y
	vel.y--
	if vel.x > 0 {
		vel.x--
	} else if vel.x < 0 {
		vel.x++
	}
}

func isInTarget(pos vector2, t targetArea) bool {
	return pos.x >= t.min.x && pos.y >= t.min.y && pos.x <= t.max.x && pos.y <= t.max.y
}

func findFloorX(minX int) int {
	var i int
	for {
		if stopsAt(i) >= minX {
			return i
		}
		i++
	}
}

func stopsAt(initialXVelocity int) int {
	var result int
	for initialXVelocity > 0 {
		result += initialXVelocity
		initialXVelocity--
	}
	return result
}
