package _4

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type board [5][5]square

type square struct {
	value int
	marked bool
}

func wins(b *board) bool {
	// check rows
	for _, row := range b {
		if all(row) {
			return true
		}
	}

	// check cols
	var col [5]square
	for i := 0; i < 5; i++ {
		for j, row := range b {
			col[j] = row[i]
		}
		if all(col) {
			return true
		}
	}

	return false
}

// all returns true if all squares in the array are marked
func all(b [5]square) bool {
	for _, v := range b {
		if !v.marked {
			return false
		}
	}
	return true
}

func callSquare(call int, boards []*board) {
	for i, b := range boards {
		for j, r := range b {
			for k, c := range r {
				if c.value == call {
					boards[i][j][k].marked = true
				}
			}
		}
	}
}

func sumUnmarked(b *board) int {
	sum := 0
	for _, r := range b {
		for _, c := range r {
			if !c.marked {
				sum += c.value
			}
		}
	}
	return sum
}

func checkForWinner(boards []*board) *board {
	for i, b := range boards {
		if wins(b) {
			fmt.Printf("Winner: Board %d\n", i)
			return b
		}
	}
	return nil
}

func nonWinners(boards []*board) []*board{
	var result []*board
	for _, b := range boards {
		if !wins(b) {
			result = append(result, b)
		}
	}
	return result
}

func inputToBingo(f *os.File) ([]int, []*board) {
	var calls []int
	var boards []*board

	var n int
	var line string
	// first line is the calls
	n, err := fmt.Fscanln(f, &line)
	if n == 0 || err != nil {
		fmt.Errorf("file read error on line 1")
	}
	split := strings.Split(line, ",")
	for _, s := range split {
		next, err := strconv.Atoi(s)
		if err != nil {
			fmt.Errorf("can't parse input")
		}
		calls = append(calls, next)
	}

	fmt.Printf("Calls: %d\n\n", calls)

	// now parse the boards
	var b *board
	rowIdx := 0
	for {
		var boardLine [5]int
		n, err = fmt.Fscanln(f, &boardLine[0], &boardLine[1], &boardLine[2], &boardLine[3], &boardLine[4])
		if err == io.EOF {
			break
		}

		if err != nil {
			// make a new board
			b = &board{}
			rowIdx = 0
			boards = append(boards, b)
		} else {
			addRow(boardLine, rowIdx, b)
			rowIdx++
		}
	}

	return calls, boards
}

func addRow(r [5]int, rowIdx int, b *board) {
	row := [5]square{}
	for i, v := range r {
		row[i].value = v
		row[i].marked = false
	}
	b[rowIdx] = row
}