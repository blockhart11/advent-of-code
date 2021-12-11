package _11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 204, do(t, "sample.txt", 10, false))
	assert.Equal(t, 1656, do(t, "sample.txt",100, false))
	assert.Equal(t, 195, do(t, "sample.txt",-1, true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", 100, false))
	fmt.Println(do(t, "input.txt", -1, true))
}

func do(t *testing.T, fName string, n int, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var lines [][]int

	for {
		var n int
		var line string
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}

		var lineInt []int
		splitLine := strings.Split(line, "")
		for _, s := range splitLine {
			i, _ := strconv.Atoi(s)
			lineInt = append(lineInt, i)
		}

		lines = append(lines, lineInt)
	}

	if !complicated {
		return cycleN(lines, n)
	} else {
		return cyclesUntilSynced(lines)
	}
}