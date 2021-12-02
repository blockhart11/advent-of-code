package _1

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

func B(in []int) int {
	count := 0
	for i, _ := range in[:len(in)-3] {
		if in[i] < in[i+3] {
			count++
		}
	}
	return count
}