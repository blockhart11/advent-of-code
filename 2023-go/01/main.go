package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	values := make([]int, 0)
	for _, line := range strings.Split(string(input), "\n") {
		// do the thing
		first, last := -1, -1
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			switch {
			case err != nil:
				continue
			case first == -1:
				first = digit
			}
			last = digit
		}
		if first == -1 || last == -1 {
			panic("I messed up")
		}
		values = append(values, (first*10)+last)
	}
	sum := utils.Fold(values, 0, func(current, next int) int {
		return current + next
	})
	fmt.Println(sum)
}
