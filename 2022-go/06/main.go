package main

import (
	"fmt"
	"os"
)

const (
	startOfPacket  = 4
	startOfMessage = 14
)

func main() {
	input, err := os.ReadFile("./2022-go/06/input.txt")
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, startOfMessage, startOfMessage)
	for i, v := range input {
		buffer[i%startOfMessage] = v
		if i < startOfMessage-1 { // buffer initialization. Would be more efficient to do this outside the loop...
			continue
		}
		if hasRepeat(buffer) {
			continue
		}
		fmt.Printf("finished at index %d", i+1)
		break
	}
}

func hasRepeat(input []byte) bool {
	buf := make(map[byte]bool)
	for _, v := range input {
		if buf[v] {
			return true
		}
		buf[v] = true
	}
	return false
}
