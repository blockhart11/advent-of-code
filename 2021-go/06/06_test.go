package _6

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 26, do(t, "sample.txt", 18))
	assert.Equal(t, 5934, do(t, "sample.txt", 80))
	assert.Equal(t, 26984457539, do(t, "sample.txt", 256))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", 80))
	fmt.Println(do(t, "input.txt", 256))
}

func do(t *testing.T, fName string, days int) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var (
		c1, c2, c3, c4, c5 int
		d1, d2, d3, d4, d5 int
	)

	for {
		var n, line int
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}
		switch line {
		case 1:
			c1++
		case 2:
			c2++
		case 3:
			c3++
		case 4:
			c4++
		case 5:
			c5++
		default:
			t.Errorf("unexpected input")
		}
	}
	// cheating a bit...
	//repls1 := growFish(1, 7, days)
	//repls2 := growFish(5, 7, days)
	//d1, d2 = repls1*c1, repls1*c2
	//d3, d4, d5 = repls2*c3, repls2*c4, repls2*c5
	d1 = growFish(1, 7, days) * c1
	d2 = growFish(2, 7, days) * c2
	d3 = growFish(3, 7, days) * c3
	d4 = growFish(4, 7, days) * c4
	d5 = growFish(5, 7, days) * c5

	return d1 + d2 + d3 + d4 + d5
}
