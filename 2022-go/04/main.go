package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./2022-go/04/input.txt")
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
		if (a <= c && b >= d) || // 2nd half contained inside first
			(a >= c && b <= d) { // 1st half contained inside second
			partOneCount++
		}
		if (a >= c && a <= d) || // 1st half lower bound inside 2nd half
			(b >= c && b <= d) || // 1st half upper bound inside 2nd half
			(c >= a && c <= b) || // 2nd half lower bound inside 1st half
			(d >= a && d <= b) { // 2nd half upper bound inside 1st half
			partTwoCount++
		}
	}
	fmt.Println(partOneCount)
	fmt.Println(partTwoCount)
}
