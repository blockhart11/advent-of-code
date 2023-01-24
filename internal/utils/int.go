package utils

import "math"

func MaxInt(ints ...int) int {
	out := math.MinInt
	for _, i := range ints {
		if i > out {
			out = i
		}
	}
	return out
}

func MinInt(ints ...int) int {
	out := math.MaxInt
	for _, i := range ints {
		if i < out {
			out = i
		}
	}
	return out
}
