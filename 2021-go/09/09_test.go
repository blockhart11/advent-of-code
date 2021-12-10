package _9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 15, do(t, "sample.txt", 5, 10, false))
	assert.Equal(t, 1134, do(t, "sample.txt", 5, 10, true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", 100, 100, false))
	fmt.Println(do(t, "input.txt", 100, 100, true))
}

func do(t *testing.T, fName string, rowSize int, colSize int, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	points := make([][]int, rowSize)
	for i, _ := range points {
		points[i] = make([]int, colSize)
	}
	var row int

	for {
		var n int
		var line string
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}

		for i, digit := range strings.Split(line, "") {
			points[row][i], _ = strconv.Atoi(digit)
		}
		row++
	}

	if !complicated {
		result, _ := risk(points)
		return result
	} else {
		return sumThreeLargestBasins(points)
	}
}