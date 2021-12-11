package _10

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 26397, do(t, "sample.txt", false))
	assert.Equal(t, 288957, do(t, "sample.txt", true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", false))
	fmt.Println(do(t, "input.txt", true))
}

func do(t *testing.T, fName string, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var lines []string

	for {
		var n int
		var line string
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}

		lines = append(lines, line)
	}

	if !complicated {
		return totalScoreCorrupted(lines)
	} else {
		return middleScoreIncomplete(lines)
	}
}