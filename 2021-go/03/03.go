package _3

import (
	"math"
	"strings"
)

func mcb(in []string) string {
	var ones []int
	for i := 0; i < len(in[0]); i++ {
		ones = append(ones, 0)
	}

	for _, val := range in {
		for col, digit := range val {
			if digit == '1' {
				ones[col] += 1
			}
		}
	}

	var result string

	for _, v := range ones {
		if v > len(in) / 2 {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result
}

func lcb(in []string) string {
	s := mcb(in)
	result := strings.ReplaceAll(s, "1", "-1")
	result = strings.ReplaceAll(result, "0", "1")
	result = strings.ReplaceAll(result, "-1", "0")

	return result
}

func oxygen(in []string) string {
	var keep []int
	for i := 0; i < len(in); i++ {
		keep = append(keep, 1)
	}

	for i := 0; i < len(in[0]); i++ {
		ones := 0
		zeros := 0
		for j, val := range in {
			if keep[j] == 1 {
				if string(val[i]) == "1" {
					ones++
				} else {
					zeros++
				}
			}
		}

		if ones >= zeros {
			//drop zeros
			for j, val := range in {
				if keep[j] == 1 && string(val[i]) == "0" {
					keep[j] = 0
				}
			}
		} else {
			//drop ones
			for j, val := range in {
				if keep[j] == 1 && string(val[i]) == "1" {
					keep[j] = 0
				}
			}
		}

		if countOnes(keep) <= 1 {
			for j, val := range keep {
				if val == 1 {
					return in[j]
				}
			}
		}
	}
	return ""
}

func co2(in []string) string {
	var keep []int
	for i := 0; i < len(in); i++ {
		keep = append(keep, 1)
	}

	for i := 0; i < len(in[0]); i++ {
		ones := 0
		zeros := 0
		for j, val := range in {
			if keep[j] == 1 {
				if string(val[i]) == "1" {
					ones++
				} else {
					zeros++
				}
			}
		}

		if zeros > ones {
			//drop zeros
			for j, val := range in {
				if keep[j] == 1 && string(val[i]) == "0" {
					keep[j] = 0
				}
			}
		} else {
			//drop ones
			for j, val := range in {
				if keep[j] == 1 && string(val[i]) == "1" {
					keep[j] = 0
				}
			}
		}

		if countOnes(keep) <= 1 {
			for j, val := range keep {
				if val == 1 {
					return in[j]
				}
			}
		}
	}
	return ""
}

func bin2dec(in string) int {
	l := len(in)
	result := 0
	for i, v := range in {
		if v == '1' {
			exp := float64(l - i - 1)
			result += int(math.Pow(2, exp))
		}
	}
	return result
}

func countOnes(i []int) int {
	r := 0
	for _, v := range i {
		if v == 1 {
			r++
		}
	}
	return r
}