package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./2022-js/lol-jk-its-go/01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var current, minOfMaxes int
	maxes := make([]int, 3, 3)
	for {
		var food int
		_, err = fmt.Fscanln(f, &food)
		if err != nil {
			break // EOF or whatever. Who cares.
		}
		if food == 0 {
			if current > minOfMaxes {
				// replace min with current
				for i, v := range maxes {
					if v == minOfMaxes {
						maxes[i] = current
						break
					}
				}
				// find new min. it's at LEAST current, so start there as the first guess
				minOfMaxes = current
				for _, v := range maxes {
					if v < minOfMaxes {
						minOfMaxes = v
					}
				}
			}
			current = 0
		} else {
			current += food
		}
	}
	fmt.Println(maxes[0] + maxes[1] + maxes[2])
}
