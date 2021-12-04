package _3

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleA(t *testing.T) {
	gamma, epsilon := doA(t, "sample.txt")

	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)
	assert.Equal(t, 198, gamma*epsilon)
}

func TestSampleB(t *testing.T) {
	oxygen, co2 := doB(t, "sample.txt")

	assert.Equal(t, 23, oxygen)
	assert.Equal(t, 10, co2)
	assert.Equal(t, 230, oxygen*co2)
}

func TestA(t *testing.T) {
	gamma, epsilon := doA(t, "input.txt")

	fmt.Printf("gamma: %d\n", gamma)
	fmt.Printf("epsilon: %d\n", epsilon)
	fmt.Printf("result: %d\n", gamma*epsilon)
}

func TestB(t *testing.T) {
	oxygen, co2 := doB(t, "input.txt")

	fmt.Printf("oxygen: %d\n", oxygen)
	fmt.Printf("co2: %d\n", co2)
	fmt.Printf("result: %d\n", oxygen*co2)
}

func doA(t *testing.T, fName string) (int, int) {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var input []string

	for {
		var line string
		var n int
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}

		input = append(input, line)
	}

	return bin2dec(mcb(input)), bin2dec(lcb(input))
}

func doB(t *testing.T, fName string) (int, int) {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var input []string

	for {
		var line string
		var n int
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}

		input = append(input, line)
	}

	oxygen, co2 := oxygen(input), co2(input)

	return bin2dec(oxygen), bin2dec(co2)
}