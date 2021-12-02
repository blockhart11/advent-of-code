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

	var inputA []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Error("can't convert file to []int")
		}
		inputA = append(inputA, next)
	}

	if err := scanner.Err(); err != nil {
		t.Error("scanner dun goofed :shrug:")
	}

	fmt.Printf("A: There are %d increases in depth!\n", A(inputA))
	
	var inputB []int
	for i, _ := range inputA {
		if i >= len(inputA) - 2 {
			break
		}
		inputB = append(inputB, inputA[i] + inputA[i+1] + inputA[i+2])
	}

	fmt.Printf("B: There are %d increases in depth!", A(inputB))
}