package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./2022-js/lol-jk-its-go/04/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var partOneCount, partTwoCount int
	for {
		var a, b, c, d int
		_, err = fmt.Fscanf(f, "%d %d %d %d", &a, &b, &c, &d)
		if err != nil {
			break
		}
		if (a == c || b == d) ||
			(a < c && b > d) ||
			(a > c && b < d) {
			partOneCount++
		}
		if (a >= c && a <= d) ||
			(b >= c && b <= d) ||
			(c >= a && c <= b) ||
			(d >= a && d <= b) {
			partTwoCount++
		}
	}
	fmt.Println(partOneCount)
	fmt.Println(partTwoCount)
}
