package _5

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleA(t *testing.T) {
	lines := do(t, "sample.txt", false)

	fmt.Println(lines.print(10))
	assert.Equal(t, 5, lines.overlap())
}

func TestSampleB(t *testing.T) {
	lines := do(t, "sample.txt", true)

	fmt.Println(lines.print(10))
	assert.Equal(t, 12, lines.overlap())
}

func TestA(t *testing.T) {
	lines := do(t, "input.txt", false)

	fmt.Printf("overlap: %d\n", lines.overlap())
}

func TestB(t *testing.T) {
	lines := do(t, "input.txt", true)

	fmt.Printf("overlap: %d\n", lines.overlap())
}

func do(t *testing.T, fName string, diagonals bool) BruteForce {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var lines BruteForce

	for {
		var x1, y1, x2, y2 int
		var n int
		n, err = fmt.Fscanln(f, &x1, &y1, &x2, &y2)
		if n == 0 || err != nil {
			break
		}

		nextLine := &LineSegment{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}

		lines.addLine(nextLine, diagonals)
	}

	return lines
}
