package _7

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, result{2, 37}, do(t, "sample.txt", false))
	assert.Equal(t, result{5, 168}, do(t, "sample.txt", true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", false))
	fmt.Println(do(t, "input.txt", true))
}

func do(t *testing.T, fName string, scaled bool) result {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var lines []int

	for {
		var n, line int
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}
		lines = append(lines, line)
	}

	if !scaled {
		return leastCost(lines)
	} else {
		return leastCostScaled(lines)
	}
}