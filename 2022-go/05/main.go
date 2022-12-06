package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	input, err := os.Open("./2022-go/05/input.txt")
	cratesRaw, err := os.ReadFile("./2022-go/05/crates.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// Make stacks
	crates := make([]stack.Stack, 9, 9)
	cratesRawSplit := strings.Split(string(cratesRaw), "\n")
	for _, v := range cratesRawSplit {
		for i, crate := range v {
			if string(crate) == " " {
				continue
			}
			crates[i].Push(int(crate))
		}
	}

	fmt.Println(peekAll(crates))
	for {
		var move, from, to int
		_, err = fmt.Fscanf(input, "move %d from %d to %d", &move, &from, &to)
		if err != nil {
			break
		}
		// adjust for 0 indexing here because I'm lazy later
		from--
		to--

		// Part 1 -- Uncomment if you dare...
		//for i := 0; i < move; i++ {
		//	crates[to].Push(crates[from].Pop())
		//}

		// Part 2
		buf := make([]interface{}, move, move)
		for i := 0; i < move; i++ {
			// Might as well populate it backwards to make the pushin' easier later
			buf[move-i-1] = crates[from].Pop()
		}
		for _, v := range buf {
			crates[to].Push(v)
		}
	}
	fmt.Println(peekAll(crates))
}

func peekAll(stacks []stack.Stack) string {
	var out string
	for _, v := range stacks {
		out += string(v.Peek().(int))
	}
	return out
}
