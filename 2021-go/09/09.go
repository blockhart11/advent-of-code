package _9

import "sort"

type point struct {
	row int
	col int
	height int
}

func risk(floor [][]int) (int, []point) {
	var result struct {
		risk int
		lowPoints []point
	}
	for i, row := range floor {
		for j, currPoint := range row {
			// check up
			if i > 0 {
				if currPoint >= floor[i-1][j] {
					continue
				}
			}
			// check left
			if j > 0 {
				if currPoint >= floor[i][j - 1] {
					continue
				}
			}
			// check right
			if j < len(row) - 1 {
				if currPoint >= floor[i][j + 1] {
					continue
				}
			}
			// check down
			if i < len(floor) - 1 {
				if currPoint >= floor[i + 1][j] {
					continue
				}
			}
			result.lowPoints = append(result.lowPoints, point{i, j, currPoint})
			result.risk += currPoint + 1
		}
	}
	return result.risk, result.lowPoints
}

func lowPoints(floor [][]int) []point {
	_, lows := risk(floor)
	return lows
}

type basin struct {
	lowPoint point
	members []point
}

func sumThreeLargestBasins(floor [][]int) int {
	var basins []basin
	lowPoints := lowPoints(floor)
	for _, low := range lowPoints {
		basin := basin{
			lowPoint: low,
			members: []point{low},
		}

		// fill basin
		fillBasin(&basin, floor, low)

		basins = append(basins, basin)
	}
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i].members) > len(basins[j].members)
	})
	return len(basins[0].members)*len(basins[1].members)*len(basins[2].members)
}

func fillBasin(b *basin, floor [][]int, start point) {
	// look up
	if start.row > 0 {
		testPoint := point{start.row - 1, start.col, floor[start.row - 1][start.col]}
		if !basinContains(b, testPoint) && start.height < testPoint.height && testPoint.height < 9 {
			b.members = append(b.members, testPoint)
			fillBasin(b, floor, testPoint)
		}
	}
	// look left
	if start.col > 0 {
		testPoint := point{start.row, start.col - 1, floor[start.row][start.col - 1]}
		if !basinContains(b, testPoint) && start.height < testPoint.height && testPoint.height < 9 {
			b.members = append(b.members, testPoint)
			fillBasin(b, floor, testPoint)
		}
	}
	// look right
	if start.col < len(floor[0]) - 1 {
		testPoint := point{start.row, start.col + 1, floor[start.row][start.col + 1]}
		if !basinContains(b, testPoint) && start.height < testPoint.height && testPoint.height < 9 {
			b.members = append(b.members, testPoint)
			fillBasin(b, floor, testPoint)
		}
	}
	// look down
	if start.row < len(floor) - 1 {
		testPoint := point{start.row + 1, start.col, floor[start.row + 1][start.col]}
		if !basinContains(b, testPoint) && start.height < testPoint.height && testPoint.height < 9 {
			b.members = append(b.members, testPoint)
			fillBasin(b, floor, testPoint)
		}
	}
}

func basinContains(b *basin, p point) bool {
	for _, basinPoint := range b.members {
		if basinPoint.row == p.row && basinPoint.col == p.col {
			return true
		}
	}
	return false
}