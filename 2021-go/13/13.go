package _13

import "fmt"

type point struct {
	x, y int
}

type fold struct {
	// if it ain't x, it's y
	xAxis bool
	value int
}

type paper struct {
	points []point
	folds []fold
}

func foldOnceAndCount(p *paper) int {
	firstFold := p.folds[0]
	if firstFold.xAxis {
		foldVertical(firstFold.value, p)
	} else {
		foldHorizontal(firstFold.value, p)
	}
	return p.countDots()
}

func (p *paper) foldAll() {
	for _, fold := range p.folds {
		if fold.xAxis {
			foldVertical(fold.value, p)
		} else {
			foldHorizontal(fold.value, p)
		}
	}
}

func (p *paper) print() {
	max := p.max()
	for i := 0; i <= max.y; i++ {
		for j := 0; j <= max.x; j++ {
			if p.contains(j, i) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (p *paper) max() point {
	var maxX, maxY int
	for _, myPoint := range p.points {
		if myPoint.x > maxX {
			maxX = myPoint.x
		}
		if myPoint.y > maxY {
			maxY = myPoint.y
		}
	}
	return point{
		x: maxX,
		y: maxY,
	}
}

func foldHorizontal(y int, p *paper) {
	for i, myPoint := range p.points {
		if myPoint.y > y {
			newY := 2*y - myPoint.y
			if p.contains(myPoint.x, newY) {
				p.points[i].y, p.points[i].x = -1, -1
			} else {
				p.points[i].y = newY
			}
		}
	}
}

func foldVertical(x int, p *paper) {
	for i, myPoint := range p.points {
		if myPoint.x > x {
			newX := 2*x - myPoint.x
			if p.contains(newX, myPoint.y) {
				p.points[i].y, p.points[i].x = -1, -1
			} else {
				p.points[i].x = newX
			}
		}
	}
}

func (p *paper) countDots() int {
	var result int
	for _, myPoint := range p.points {
		if myPoint.x == -1 || myPoint.y == -1 {
			continue
		}
		result += 1
	}
	return result
}

func (p *paper) contains(x, y int) bool {
	for _, v := range p.points {
		if v.x == x && v.y == y {
			return true
		}
	}
	return false
}