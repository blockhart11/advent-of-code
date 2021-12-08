package _7

import (
	"math"
	"sort"
)

type result struct {
	pos int
	cost int
}

func leastCost(crabs []int) result {
	sort.Ints(crabs)
	// starting costs for each side assuming 0 index first
	var lCost, lCount, rCost, rCount, tracker int
	for _, v := range crabs {
		if v != 0 {
			rCount++
			rCost += v
		} else {
			lCount++
			tracker++
		}
	}
	currResult := result{
		pos:  0,
		cost: rCost,
	}

	for i := 0; i < crabs[len(crabs)-1]; i++ {
		cost := rCost+lCost
		//fmt.Printf("Cost at %d: %d\n", i, cost)
		if cost < currResult.cost{
			currResult.pos = i
			currResult.cost = cost
		}

		rCost -= rCount
		lCost += lCount
		for crabs[tracker] <= i + 1 {
			rCount--
			lCount++
			tracker++
			if tracker >= len(crabs) {
				break
			}
		}
	}

	return currResult
}

func leastCostScaled(crabs []int) result {
	sort.Ints(crabs)
	r := result{
		pos:  0,
		cost: cost(crabs, 0),
	}
	// lets try brute force?
	for i := 1; i <= crabs[len(crabs) - 1]; i++ {
		costNext := cost(crabs, i)
		if costNext < r.cost {
			r.pos = i
			r.cost = costNext
		}
		//fmt.Printf("Cost at %d: %d\n", i, costNext)
	}
	return r
}

func cost(crabs []int, pos int) int {
	var r int
	for _, v := range crabs {
		diff := math.Abs(float64(pos - v))
		r += int((diff * (diff + 1)) / 2)
	}
	return r
}