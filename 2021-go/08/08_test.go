package _8

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 26, do(t, "sample.txt", false))
	assert.Equal(t, 61229, do(t, "sample.txt", true))
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

	var displays []display

	for {
		var n int
		var d display
		var bar string
		n, err = fmt.Fscanln(
			f,
			&d.p0,
			&d.p1,
			&d.p2,
			&d.p3,
			&d.p4,
			&d.p5,
			&d.p6,
			&d.p7,
			&d.p8,
			&d.p9,
			&bar,
			&d.d0,
			&d.d1,
			&d.d2,
			&d.d3)
		if n == 0 || err != nil {
			break
		}
		displays = append(displays, d)
	}

	if !complicated {
		return countEasyDigits(displays)
	} else {
		return decodeAll(displays)
	}
}