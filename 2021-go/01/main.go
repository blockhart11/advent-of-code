package main

func HoHoHo() string {
	return "Ho ho ho!"
}

func A(in []int) int {
	count := 0
	prev := 1000000    // good enough for me!
	for _, v := range in {
		if v > prev {
			count++
		}
		prev = v
	}
	return count
}