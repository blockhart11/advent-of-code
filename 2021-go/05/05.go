package _5

import "strconv"

type LineSegment struct {
	x1, y1, x2, y2 int
}

type BruteForce [1000][1000]int

func (b *BruteForce) addLine(line *LineSegment, diagonals bool) {
	var start, end int
	// horizontal
	if line.y1 == line.y2 {
		start = minInt(line.x1, line.x2)
		end = maxInt(line.x1, line.x2)
		for i := start; i <= end; i++ {
			b[line.y1][i]++
		}
	}

	// vertical
	if line.x1 == line.x2 {
		start = minInt(line.y1, line.y2)
		end = maxInt(line.y1, line.y2)
		for i := start; i <= end; i++ {
			b[i][line.x1]++
		}
	}

	// diagonal - part b only
	if diagonals && line.x1 != line.x2 && line.y1 != line.y2 {
		slope := 1
		var magnitude int
		var startY int
		if line.x1 < line.x2 {
			// left to right
			start = line.x1
			end = line.x2
			startY = line.y1
			if line.y1 > line.y2 {
				//descending
				slope = -1
			}
		} else {
			// right to left
			start = line.x2
			end = line.x1
			startY = line.y2
			if line.y2 > line.y1 {
				//descending
				slope = -1
			}
		}
		magnitude = end - start
		for i := 0; i <= magnitude; i++ {
			b[startY+(i*slope)][start+i]++
		}
	}
}

func (b *BruteForce) print(maxRows int) string {
	var result string
	for i, row := range b {
		if i >= maxRows {
			break
		}
		for j, col := range row {
			if j >= maxRows {
				break
			}
			if col == 0 {
				result += "."
			} else {
				result += strconv.Itoa(col)
			}
		}
		result += "\n"
	}
	return result
}

func (b *BruteForce) overlap() int {
	result := 0
	for _, row := range b {
		for _, col := range row {
			if col >= 2 {
				result++
			}
		}
	}
	return result
}

func minInt(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}