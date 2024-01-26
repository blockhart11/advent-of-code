package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
)

// Expected output
//  sample.txt: 35

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Declare data structures
	result := make([]any, 0)

	// Parse the input
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		_ = fmt.Sprint(i, line) // no compiler error

		// Solve the problem
	}

	// Print the solution
	fmt.Println(utils.Fold(result, nil, func(current, next any) any {
		return nil
	}))
}
