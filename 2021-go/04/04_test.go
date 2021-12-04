package _4

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleA(t *testing.T) {
	unused, lastCall := doA(t, "sample.txt")

	assert.Equal(t, 188, unused)
	assert.Equal(t, 24, lastCall)
	assert.Equal(t, 4512, unused*lastCall)
}

func TestSampleB(t *testing.T) {
	unused, lastCall := doB(t, "sample.txt")

	assert.Equal(t, 148, unused)
	assert.Equal(t, 13, lastCall)
	assert.Equal(t, 1924, unused*lastCall)
}

func TestA(t *testing.T) {
	unused, lastCall := doA(t, "input.txt")

	fmt.Printf("unused: %d\n", unused)
	fmt.Printf("lastCall: %d\n", lastCall)
	fmt.Printf("result: %d\n", unused*lastCall)
}

func TestB(t *testing.T) {
	unused, lastCall := doB(t, "input.txt")

	fmt.Printf("unused: %d\n", unused)
	fmt.Printf("lastCall: %d\n", lastCall)
	fmt.Printf("result: %d\n", unused*lastCall)
}

func doA(t *testing.T, fName string) (int, int) {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	calls, boards := inputToBingo(f)

	for _, v := range calls {
		callSquare(v, boards)
		winner := checkForWinner(boards)
		if winner != nil {
			return sumUnmarked(winner), v
		}
	}

	return -1, -1
}

func doB(t *testing.T, fName string) (int, int) {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	calls, boards := inputToBingo(f)
	var lastWinner *board
	var lastCall int

	for _, v := range calls {
		callSquare(v, boards)

		if lastWinner != nil {
			// has it won yet?
			if wins(lastWinner) {
				lastCall = v
				break
			}
		}

		losers := nonWinners(boards)
		if len(losers) == 1 {
			lastWinner = losers[0]
		}
	}

	return sumUnmarked(lastWinner), lastCall
}
