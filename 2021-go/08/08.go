package _8

import (
	"bytes"
	"sort"
	"strings"
)

type display struct {
	p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 string
	d0, d1, d2, d3 string
}

type pattern struct {
	wires []byte
	value int
}

type betterDisplay struct {
	wireMap map[byte]byte
	patterns []pattern
	digits []string
}

func countEasyDigits(displays []display) int {
	var count int
	for _, d := range displays {
		if isUnique(d.d0) {
			count++
		}
		if isUnique(d.d1) {
			count++
		}
		if isUnique(d.d2) {
			count++
		}
		if isUnique(d.d3) {
			count++
		}
	}
	return count
}

func decodeAll(displays []display) int {
	var result int
	for _, display := range displays {
		result += decode(display)
	}
	return result
}

func decode(d display) int {
	bd := toBetterDisplay(d)
	result, mult := 0, 1000
	for _, digit := range bd.digits {
		for _, pattern := range bd.patterns {
			if len(digit) == len(pattern.wires) {
				match := true
				for _, wire := range pattern.wires {
					if !strings.Contains(digit, string(wire)) {
						match = false
						break
					}
				}
				if match {
					result += mult * pattern.value
					break
				}
			}
		}
		mult /= 10
	}
	return result
}

func toBetterDisplay(d display) betterDisplay {
	bd := betterDisplay{
		wireMap:  nil,
		patterns: make([]pattern, 10),
		digits:   make([]string, 4),
	}
	bd.patterns[0] = pattern{[]byte(d.p0), -1}
	bd.patterns[1] = pattern{[]byte(d.p1), -1}
	bd.patterns[2] = pattern{[]byte(d.p2), -1}
	bd.patterns[3] = pattern{[]byte(d.p3), -1}
	bd.patterns[4] = pattern{[]byte(d.p4), -1}
	bd.patterns[5] = pattern{[]byte(d.p5), -1}
	bd.patterns[6] = pattern{[]byte(d.p6), -1}
	bd.patterns[7] = pattern{[]byte(d.p7), -1}
	bd.patterns[8] = pattern{[]byte(d.p8), -1}
	bd.patterns[9] = pattern{[]byte(d.p9), -1}
	bd.digits[0] = d.d0
	bd.digits[1] = d.d1
	bd.digits[2] = d.d2
	bd.digits[3] = d.d3
	// sort by count of wires
	sort.Slice(bd.patterns, func(i, j int) bool {
		return len(bd.patterns[i].wires) < len(bd.patterns[j].wires)
	})
	for i, r := range bd.patterns {
		sort.Slice(r.wires, func(i, j int) bool {
			return r.wires[i] < r.wires[j]
		})
		switch len(r.wires) {
		case 2: // one digit
			bd.patterns[i].value = 1
		case 3: // seven digit
			bd.patterns[i].value = 7
		case 4: // four digit
			bd.patterns[i].value = 4
		case 5: // two, three, or five digit
		// check for two - it matches exactly 2 wires from 4
		if getIdxOfDigit(2, bd) == -1 {
			four := getIdxOfDigit(4, bd)
			var count int
			for _, wire := range bd.patterns[four].wires {
				if !bytes.ContainsRune(r.wires, rune(wire)) {
					count++
				}
			}
			// if exactly two wires overlap, this must be 2
			if count == 2 {
				bd.patterns[i].value = 2
			}
		}
		// check for three - it has both of the wires in 1
		if getIdxOfDigit(3, bd) == -1 {
			one := getIdxOfDigit(1, bd)
			isThree := true
			for _, wire := range bd.patterns[one].wires {
				if !bytes.ContainsRune(r.wires, rune(wire)) {
					// if the wire isn't in this pattern, this isn't three!
					isThree = false
					break
				}
			}
			if isThree {
				bd.patterns[i].value = 3
			}
		}
		// check for five - it'll be the last one
		case 6: // zero, six, or nine digit
		// check for zero - it will only have 1 of the 2 bars in 4 that are not in 1
		if getIdxOfDigit(0, bd) == -1 {
			four := getIdxOfDigit(4, bd)
			one := getIdxOfDigit(1, bd)
			// get the wires in four that aren't also in one
			var fourButNotOneWires []byte
			for _, wire := range bd.patterns[four].wires {
				if !bytes.ContainsRune(bd.patterns[one].wires, rune(wire)) {
					fourButNotOneWires = append(fourButNotOneWires, wire)
				}
			}
			// of these wires (there should be two), count how many are in the current number
			var count int
			for _, wire := range fourButNotOneWires {
				if bytes.ContainsRune(r.wires, rune(wire)) {
					count++
				}
			}
			// if exactly one wire matches, this is the zero digit!
			if count == 1 {
				bd.patterns[i].value = 0
			}
		}
		// check for six - it's the only one without both wires also in one
		if getIdxOfDigit(6, bd) == -1 {
			one := getIdxOfDigit(1, bd)
			for _, wire := range bd.patterns[one].wires {
				if !bytes.ContainsRune(r.wires, rune(wire)) {
					// if the wire isn't in this pattern, this is six!
					bd.patterns[i].value = 6
				}
			}
		}
		// check for nine - it contains all of the same wires that four does
		if getIdxOfDigit(9, bd) == -1 {
			four := getIdxOfDigit(4, bd)
			isNine := true
			for _, wire := range bd.patterns[four].wires {
				if !bytes.ContainsRune(r.wires, rune(wire)) {
					// if the wire isn't in this pattern, this isn't nine!
					isNine = false
					break
				}
			}
			if isNine {
				bd.patterns[i].value = 9
			}
		}
		case 7: // eight digit
			bd.patterns[i].value = 8
		}
	}
	// iterate again and populate the final one, digit 5
	for i, r := range bd.patterns {
		if r.value == -1 {
			bd.patterns[i].value = 5
			break
		}
	}
	return bd
}

// getIdxOfDigit returns the index in the display where the digit is found
func getIdxOfDigit(num int, bd betterDisplay) int {
	for i, v := range bd.patterns {
		if v.value == num {
			return i
		}
	}
	return -1
}

func isUnique(in string) bool {
	l := len(in)
	if l == 2 || l == 3 || l == 4 || l == 7 {
		return true
	}
	return false
}