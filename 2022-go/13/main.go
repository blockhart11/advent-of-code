package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//input, err := os.ReadFile("./2022-go/13/sample.txt")
	input, err := os.ReadFile("./2022-go/13/input.txt")
	if err != nil {
		panic(err)
	}
	pairs := strings.Split(string(input), "\n\n")

	var allPackets []*element

	var rightOrders []int
	for i, pair := range pairs {
		packets := strings.Split(pair, "\n")
		index := 0
		left := parsePacket(packets[0], &index)
		index = 0
		right := parsePacket(packets[1], &index)
		if compare(left, right) == 1 {
			rightOrders = append(rightOrders, i+1)
		}
		allPackets = append(allPackets, &left, &right)
	}
	fmt.Println(rightOrders)
	var count int
	for _, v := range rightOrders {
		count += v
	}
	fmt.Println(count)

	// Part 2
	dividerA := &element{
		value: -1,
		list: []element{{
			value: -1,
			list:  []element{{value: 2}},
		}},
	}
	dividerB := &element{
		value: -1,
		list: []element{{
			value: -1,
			list:  []element{{value: 6}},
		}},
	}
	allPackets = append(allPackets, dividerA, dividerB)
	sort.Slice(allPackets, func(i int, j int) bool {
		return compare(*allPackets[i], *allPackets[j]) > 0
	})

	var idxA, idxB int
	for i, v := range allPackets {
		if v == dividerA {
			idxA = i + 1
		} else if v == dividerB {
			idxB = i + 1
		}
	}
	fmt.Printf("%d, %d", idxA, idxB)

}

type element struct {
	value int
	list  []element
}

func (e element) isList() bool {
	return e.value == -1
}

func parsePacket(input string, index *int) element {
	out := element{value: -1}

	for *index < len(input) {
		switch input[*index] {
		case '[':
			*index++
			out.list = parseElements(input, index)
		case '1':
			// special case if 10
			if input[*index+1] == '0' {
				out.value = 10
				*index += 2
				return out
			} else {
				out.value = 1
				*index++
				return out
			}
		case '0', '2', '3', '4', '5', '6', '7', '8', '9':
			v, err := strconv.Atoi(string(input[*index]))
			if err != nil {
				panic(err)
			}
			out.value = v
			*index++
			return out
		case ',':
			*index++
		case ']':
			return out
		default:
			panic("how did I get here")
		}
	}

	return out
}

func parseElements(input string, index *int) []element {
	var out []element
	for *index < len(input) {
		switch input[*index] {
		case '[':
			*index++
			elem := element{value: -1, list: parseElements(input, index)}
			out = append(out, elem)
		case '1':
			// special case if 10
			if input[*index+1] == '0' {
				out = append(out, element{value: 10})
				*index += 2
			} else {
				out = append(out, element{value: 1})
				*index++
			}
		case '0', '2', '3', '4', '5', '6', '7', '8', '9':
			v, err := strconv.Atoi(string(input[*index]))
			if err != nil {
				panic(err)
			}
			out = append(out, element{value: v})
			*index++
		case ',':
			*index++
			continue
		case ']':
			*index++
			return out
		}
	}
	panic("list never terminated")
}

func compare(left, right element) int {
	switch {
	case left.isList() && right.isList():
		for i := range left.list {
			if i >= len(right.list) {
				// right ran out of items
				return -1
			} else {
				switch compare(left.list[i], right.list[i]) {
				case 1:
					// left is first
					return 1
				case -1:
					// right is first
					return -1
				case 0:
					// tie. move on
					continue
				}
			}
		}
		if len(right.list) > len(left.list) {
			// right still has items
			return 1
		} else {
			return 0
		}
	case left.isList() && !right.isList():
		right.list = []element{{value: right.value}}
		right.value = -1
		return compare(left, right)
	case !left.isList() && right.isList():
		left.list = []element{{value: left.value}}
		left.value = -1
		return compare(left, right)
	default:
		switch {
		case left.value < right.value:
			return 1
		case right.value < left.value:
			return -1
		default:
			return 0
		}
	}
}
