package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/blockhart11/advent-of-code/internal/utils"
)

const (
	day        = 16
	minutes    = 30
	startValve = "AA"

	// sample input
	//inputFile = "sample.txt"

	// real input
	inputFile = "input.txt"
)

var (
	bestSoFar int
	numValves int
	cave      map[string]valve
)

func main() {
	input, err := os.ReadFile("./2022-go/" + utils.Itoa(day) + "/" + inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	cave = make(map[string]valve)

	// create cave and valves
	for _, line := range lines {
		args := strings.Split(line, " ")
		v := valve{
			name:  args[0],
			rate:  utils.MustAtoi(args[1]),
			costs: make(map[string]int),
		}
		for i := 2; i < len(args); i++ {
			v.tunnels = append(v.tunnels, args[i])
		}
		cave[v.name] = v
		if v.rate > 0 {
			numValves++
		}
	}

	// We really only need nodes with a flow rate, and the cost to between them
	for k, v := range cave {
		if v.rate > 0 || v.name == startValve {
			for k2, v2 := range cave {
				if v2.rate > 0 && v.costs[k2] == 0 {
					if cost := cheapestCostTo(v, v2); cost > 0 {
						v.costs[k2] = cost
						// if not starting valve
						if v.name != startValve {
							v2.costs[k] = cost
						}
					}
				}
			}
		}
	}

	// now delete all the nodes without a flow rate
	for k, v := range cave {
		if v.rate == 0 && v.name != startValve {
			delete(cave, k)
		}
	}

	// part 1 - do the thing
	best := 0
	fmt.Println(mostFlowFrom(startValve, "", minutes, &best, 0))

	// part 2 - you and the elephant both do the thing
	fmt.Println(mostFlowFromWithElephant(startValve, "", minutes-4, &best, 0))
}

type valve struct {
	name    string
	rate    int
	tunnels []string
	costs   map[string]int // cost to get to each meaningful destination
}

func cheapestCostTo(from, to valve) int {
	switch {
	case from.name == to.name:
		return 0
	default:
		best := math.MaxInt
		return cheapestCostHelper(from, to, "", &best, 0)
	}
}

func cheapestCostHelper(from, to valve, visited string, best *int, accum int) int {
	switch {
	case from.name == to.name:
		if *best > accum {
			*best = accum
		}
		return accum
	case accum >= *best:
		return -1
	case strings.Contains(visited, from.name):
		return -1
	default:
		var results []int
		for _, v := range from.tunnels {
			if cost := cheapestCostHelper(cave[v], to, visited+" "+from.name, best, accum+1); cost > 0 {
				results = append(results, cost)
			}
		}
		return utils.MinInt(results...)
	}
}

func mostFlowFrom(start string, path string, left int, best *int, accum int) int {
	// sort options by value
	type option struct {
		to    string
		cost  int
		value int
	}
	var options []option
	for k, v := range cave[start].costs {
		if !strings.Contains(path, k) {
			if value := valueOfValve(cave[k].rate, v, left); value > 0 {
				options = append(options, option{to: k, cost: v, value: value})
			}
		}
	}
	// if there are no options, job's done
	if len(options) == 0 {
		if accum > *best {
			path += " " + start
			*best = accum
			fmt.Printf("New best path: %s, score %d\n", path, *best)
		}
		return accum
	}
	// another short circuit option
	var maxScoreLeft int
	for _, v := range options {
		maxScoreLeft += v.value
	}
	if accum+maxScoreLeft < *best {
		return accum
	}
	// sort descending
	sort.Slice(options, func(i, j int) bool {
		return options[i].value > options[j].value
	})

	// try each option in order
	var results []int
	for _, v := range options {
		results = append(results, mostFlowFrom(v.to, path+" "+start, left-v.cost-1, best, accum+v.value))
	}
	return utils.MaxInt(results...)
}

func mostFlowFromWithElephant(start, path string, left int, best *int, accum int) int {
	// sort options by value
	type option struct {
		to    string
		cost  int
		value int
	}
	var options []option
	for k, v := range cave[start].costs {
		if !strings.Contains(path, k) {
			if value := valueOfValve(cave[k].rate, v, left); value > 0 {
				options = append(options, option{to: k, cost: v, value: value})
			}
		}
	}
	// if there are no options, job's done
	if len(options) == 0 {
		if accum > *best {
			path += " " + start
			*best = accum
			fmt.Printf("new best path: %s, score %d\n", path, *best)
		}
		return accum
	}
	// sort descending
	sort.Slice(options, func(i, j int) bool {
		return options[i].value > options[j].value
	})

	var results []int
	// if the human stopped here, what route could the elephant have taken?
	eBest := 0
	path += " " + start
	eScore := mostFlowFrom(startValve, path+" | ", minutes-4, &eBest, 0)
	if accum+eScore > *best {
		*best = accum + eScore
		fmt.Printf("New best human path: %s, score %d\n", path, *best)
	}
	results = append(results, accum+eScore)
	// but if the human had kept going...
	for _, v := range options {
		results = append(results, mostFlowFromWithElephant(v.to, path, left-v.cost-1, best, accum+v.value))
	}
	return utils.MaxInt(results...)
}

func valueOfValve(rate, cost, left int) int {
	return (left - cost - 1) * rate
}
