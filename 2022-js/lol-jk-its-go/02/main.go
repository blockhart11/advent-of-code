package main

import (
	"fmt"
	"os"
)

const (
	win      = 6
	lose     = 0
	draw     = 3
	rock     = 1
	paper    = 2
	scissors = 3
)

var gameMap = map[string]int{
	"AX": draw + rock,
	"AY": win + paper,
	"AZ": lose + scissors,
	"BX": lose + rock,
	"BY": draw + paper,
	"BZ": win + scissors,
	"CX": win + rock,
	"CY": lose + paper,
	"CZ": draw + scissors,
}
var gameMap2 = map[string]int{
	"AX": lose + scissors,
	"AY": draw + rock,
	"AZ": win + paper,
	"BX": lose + rock,
	"BY": draw + paper,
	"BZ": win + scissors,
	"CX": lose + paper,
	"CY": draw + scissors,
	"CZ": win + rock,
}

func main() {
	f, err := os.Open("./2022-js/lol-jk-its-go/02/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var points int
	for {
		var line string
		_, err = fmt.Fscanln(f, &line)
		if err != nil {
			break
		}
		//points += gameMap[line] // for part 1
		points += gameMap2[line]
	}
	fmt.Println(points)
}
