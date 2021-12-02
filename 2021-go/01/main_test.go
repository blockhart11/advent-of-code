package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHoHoHo(t *testing.T) {
	e := "Ho ho ho!"
	assert.Equal(t, e, HoHoHo())
}

func TestAB(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Error("can't open file")
	}
	defer file.Close()

	var inA []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Error("can't convert file to []int")
		}
		inA = append(inA, next)
	}

	if err := scanner.Err(); err != nil {
		t.Error("scanner dun goofed :shrug:")
	}

	fmt.Printf("A: There are %d increases in depth!\n", A(inA))

	var inB []int
	for i, _ := range inA[:len(inA)-2] {
		inB = append(inB, inA[i] + inA[i+1] + inA[i+2])
	}

	fmt.Printf("B: There are %d increases in depth!\n", A(inB))

	// V2.0
	fmt.Printf("B, but better: There are (still) %d increases in depth!\n", B(inA))
}