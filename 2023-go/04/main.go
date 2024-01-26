package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
)

// Expected output
//  sample.txt: 13

func main() {
	input, err := os.ReadFile(os.Args[1])
	//input, err := os.ReadFile("./2023-go/04/sample.txt")
	if err != nil {
		panic(err)
	}

	// Declare data structures
	lines := strings.Split(string(input), "\n")
	var totalPoints int
	cardsCount := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cardsCount[i] = 1
	}

	// Parse the input
	for i, line := range lines {
		_ = fmt.Sprint(i, line) // no compiler error

		// Solve the problem
		substrs := strings.Split(line, "|")
		numbersYouHave := strings.Split(strings.TrimSpace(substrs[1]), " ")
		substrs = strings.Split(substrs[0], ":")
		winningNumbers := strings.Split(strings.TrimSpace(substrs[1]), " ")
		var points int

		offset := 1
		for _, num := range numbersYouHave {
			if num == "" {
				continue
			}
			if utils.Contains(winningNumbers, num) {
				// part 1
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				// part 2
				cardsCount[i+offset] += cardsCount[i]
				offset++
			}
		}
		fmt.Printf("card %d worth %d points\n", i, points)
		totalPoints += points
	}

	// Print the solution
	fmt.Println(totalPoints)

	// Part 2
	var part2Sum int
	for i, val := range cardsCount {
		fmt.Printf("Card %d: %d\n", i, val)
		part2Sum += val
	}
	fmt.Println(part2Sum)
}
