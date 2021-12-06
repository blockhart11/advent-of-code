package _6

// part A would be growFish(_, 7, 80)
func growFish(timer int, cycleLength int, days int) int {
	return growHelper(timer, cycleLength, 1, days)
}

func growHelper(timer int, cycleLength int, today int, days int) int {
	if today > days {
		return 1
	}

	if timer <= 0 {
		return growHelper(8, cycleLength, today + 1, days) +
				growHelper(cycleLength-1, cycleLength, today + 1, days)
	} else {
		// skip to zero to save cycles
		return growHelper(0, cycleLength, today + timer, days)
	}
}