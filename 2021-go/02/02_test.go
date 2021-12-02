package _2

import (
	"fmt"
	"os"
	"testing"
)

func TestMoveAll(t *testing.T) {
	fmt.Printf("\n\n*****  DAY TWO  ****\n\n")

	f, err := os.Open("input.txt")
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var c []command

	for {
		var cmd string
		var value int
		var n int
		n, err = fmt.Fscanln(f, &cmd, &value)
		if n == 0 || err != nil {
			break
		}
		cNew := command{
			cmd:   cmd,
			value: value,
		}
		c = append(c, cNew)
	}

	d, h := moveAll(c)
	fmt.Printf("A: Depth %d, Horizontal %d, Product %d\n", d, h, d*h)

	d, h = moveAllAim(c)
	fmt.Printf("B: Depth %d, Horizontal %d, Product %d\n", d, h, d*h)
}