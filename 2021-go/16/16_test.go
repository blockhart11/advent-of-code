package _16

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 16, do(t, "sample.txt", false))
	assert.Equal(t, 12, do(t, "sampleB.txt", false))
	assert.Equal(t, 23, do(t, "sampleC.txt", false))
	assert.Equal(t, 31, do(t, "sampleD.txt", false))
	assert.Equal(t, 3, do(t, "sampleE.txt", true))
	assert.Equal(t, 54, do(t, "sampleF.txt", true))
	assert.Equal(t, 7, do(t, "sampleG.txt", true))
	assert.Equal(t, 9, do(t, "sampleH.txt", true))
	assert.Equal(t, 1, do(t, "sampleI.txt", true))
	assert.Equal(t, 0, do(t, "sampleJ.txt", true))
	assert.Equal(t, 0, do(t, "sampleK.txt", true))
	assert.Equal(t, 1, do(t, "sampleL.txt", true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", false))
}

func do(t *testing.T, fName string, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var packetList string
	var n int
	n, err = fmt.Fscanln(f, &packetList)
	if n == 0 || err != nil {
		t.Errorf("can't read input")
	}

	count, eval := processInput(packetList)
	if !complicated {
		return count
	} else {
		return eval
	}
}
