package utils

import "strconv"

func MustAtoi(input string) int {
	out, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return out
}

func Itoa(input int) string {
	return strconv.Itoa(input)
}
