package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
)

// Expected output
//  sample.txt: 8, 2286

type color int

const (
	red color = iota
	green
	blue

	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Part 1
	possibleGameIDs := make([]int, 0)
Games:
	for _, game := range strings.Split(string(input), "\n") {
		words := strings.Split(game, " ")
		gameID, err := strconv.Atoi(words[1][:len(words[1])-1])
		if err != nil {
			panic("gameID:" + words[1][:len(words[1])-1] + "not a number")
		}
		for i := 2; i < len(words); i += 2 {
			amount, err := strconv.Atoi(words[i])
			if err != nil {
				panic(words[i] + "not a number")
			}
			if amount > maxColor(getColor(words[i+1])) {
				// Not a viable game. continue outer loop
				continue Games
			}
		}
		possibleGameIDs = append(possibleGameIDs, gameID)
	}
	fmt.Println(utils.Fold(possibleGameIDs, 0, func(current, next int) int {
		return current + next
	}))

	// Part 2
	gamePowers := make([]int, 0)
	for _, game := range strings.Split(string(input), "\n") {
		mostRed, mostGreen, mostBlue := 0, 0, 0
		words := strings.Split(game, " ")
		for i := 2; i < len(words); i += 2 {
			amount, err := strconv.Atoi(words[i])
			if err != nil {
				panic(words[i] + "not a number")
			}
			switch getColor(words[i+1]) {
			case red:
				mostRed = utils.MaxInt(amount, mostRed)
			case green:
				mostGreen = utils.MaxInt(amount, mostGreen)
			case blue:
				mostBlue = utils.MaxInt(amount, mostBlue)
			}
		}
		gamePowers = append(gamePowers, mostRed*mostGreen*mostBlue)
	}
	fmt.Println(utils.Fold(gamePowers, 0, func(current, next int) int {
		return current + next
	}))
}

func getColor(input string) color {
	switch strings.TrimRight(input, ",;") {
	case "red":
		return red
	case "green":
		return green
	case "blue":
		return blue
	}
	panic(input + " is not a color")
}

func maxColor(input color) int {
	switch input {
	case red:
		return maxRed
	case green:
		return maxGreen
	case blue:
		return maxBlue
	}
	panic(fmt.Sprintf("%d is not a color", input))
}
