package _1

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHoHoHo(t *testing.T) {
	e := "Ho ho ho!"
	assert.Equal(t, e, HoHoHo())
}

func TestAB(t *testing.T) {
	fmt.Printf("\n\n*****  DAY ONE  *****\n\n")

	f, err := os.Open("input.txt")
	if err != nil {
		t.Error("can't open f")
	}
	defer f.Close()

	var inA []int

	for {
		var n, value int
		n, err = fmt.Fscanln(f, &value)
		if n == 0 || err != nil {
			t.Error("file scanner borked")
			break
		}
		inA = append(inA, value)
	}

	fmt.Printf("A: There are %d increases in depth!\n", A(inA))

	var inB []int
	for i := range inA[:len(inA)-2] {
		inB = append(inB, inA[i] + inA[i+1] + inA[i+2])
	}

	fmt.Printf("B: There are %d increases in depth!\n", A(inB))

	// V2.0
	fmt.Printf("B, but better: There are (still) %d increases in depth!\n", B(inA))
}