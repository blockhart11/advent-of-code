package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./2022-js/lol-jk-its-go/03/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var answerOne, answerTwo int
	for {
		var line1, line2, line3 string
		_, err = fmt.Fscanln(f, &line1)
		if err != nil {
			break
		}
		_, err = fmt.Fscanln(f, &line2)
		if err != nil {
			break
		}
		_, err = fmt.Fscanln(f, &line3)
		if err != nil {
			break
		}
		answerOne += partOne(line1)
		answerOne += partOne(line2)
		answerOne += partOne(line3)
		answerTwo += partTwo(line1, line2, line3)
	}
	fmt.Printf("Part 1: %d\n", answerOne)
	fmt.Printf("Part 2: %d\n", answerTwo)
}

func priority(item rune) int {
	switch {
	case item >= rune('A') && item <= rune('Z'):
		return int(item - rune('A') + 27)
	case item >= rune('a') && item <= rune('z'):
		return int(item - rune('a') + 1)
	}
	return -1
}

func partOne(rucksack string) int {
	left := rucksack[:len(rucksack)/2]
	right := rucksack[len(rucksack)/2:]

	catalog := make(map[rune]bool)
	for _, v := range left {
		catalog[v] = true
	}
	for _, v := range right {
		if catalog[v] {
			return priority(v)
		}
	}
	panic("please like and subscribe")
}

func partTwo(line1, line2, line3 string) int {
	catalog := make(map[rune]int)
	for _, v := range line1 {
		catalog[v] = 1
	}
	for _, v := range line2 {
		if catalog[v] == 1 {
			catalog[v] = 2
		}
	}
	for _, v := range line3 {
		if catalog[v] == 2 {
			return priority(v)
		}
	}
	panic("yikes")
}
