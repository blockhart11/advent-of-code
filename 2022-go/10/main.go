package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//input, err := os.ReadFile("./2022-go/10/sample.txt")
	input, err := os.ReadFile("./2022-go/10/input.txt")
	if err != nil {
		panic(err)
	}

	// initialization
	cycle, x := 1, 1
	var signal, addend int
	image := make([]bool, 240, 240)

	// process all commands
	for _, line := range strings.Split(string(input), "\n") {
		checkImage(image, cycle-1, x)
		checkCycle(&signal, cycle, x)
		if line != "noop" {
			addend, err = strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				panic("oh crap")
			}
			cycle++
			checkImage(image, cycle-1, x)
			checkCycle(&signal, cycle, x)
			x += addend
		}
		cycle++
	}
	fmt.Println(signal)

	// Part 2: print the image
	for i, v := range image {
		if v {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		if i%40 == 39 {
			fmt.Printf("\n")
		}
	}
}

func checkCycle(signal *int, cycle, x int) {
	switch cycle {
	case 20, 60, 100, 140, 180, 220:
		*signal += cycle * x
	}
}

func checkImage(image []bool, cycle, x int) {
	if x-1 <= cycle%40 && cycle%40 <= x+1 {
		image[cycle] = true
	}
}
