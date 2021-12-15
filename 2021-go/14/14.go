package _14

import (
	"fmt"
)

type polymerTemplate struct {
	startState string
	rules      []rule
}

type rule struct {
	pattern string
	insert  string
}

//example
// {"NN": [1]: {"N":1, "C":1}}
type cache map[string][]map[string]int

func mostLessLeastWithCache(state string, ruleList []rule, cycles int) int {
	// initialize cache
	megaCache := make(cache, len(ruleList))
	for _, v := range ruleList {
		megaCache[v.pattern] = make([]map[string]int, cycles+1)
		for i, _ := range megaCache[v.pattern] {
			megaCache[v.pattern][i] = make(map[string]int)
		}
	}

	for i, _ := range state[:len(state)-1] {
		//fmt.Println("outer loop -", i)
		megaCache[state[i:i+2]][cycles] = cycleNWithCache(state[i:i+2], ruleList, megaCache, cycles)
	}

	// calculate answer
	result := make(map[string]int)
	var max, min int
	for _, pattern := range megaCache {
		copyMap(result, pattern[cycles])
	}
	// final character
	result[state[len(state)-1:]]++

	for _, v := range result {
		if max < v {
			max = v
		}
		if min > v || min == 0 {
			min = v
		}
	}
	//fmt.Println(result)
	//fmt.Println(megaCache)
	fmt.Println("max", max)
	fmt.Println("min", min)
	return max - min
}

func cycleNWithCache(pattern string, ruleList []rule, myCache cache, cycles int) map[string]int {
	result := make(map[string]int)
	// check cache hit
	if len(myCache[pattern][cycles]) != 0 {
		return myCache[pattern][cycles]
	}

	// cache miss
	if cycles == 0 {
		result[pattern[:1]]++
	} else {
		myRule := getRule(pattern, ruleList)
		lhs := pattern[:1] + myRule.insert
		rhs := myRule.insert + pattern[1:]
		copyMap(result, cycleNWithCache(lhs, ruleList, myCache, cycles-1))
		copyMap(result, cycleNWithCache(rhs, ruleList, myCache, cycles-1))
	}

	// add answer to cache
	myCache[pattern][cycles] = result

	return result
}

func getRule(pattern string, ruleList []rule) *rule {
	for i, r := range ruleList {
		if r.pattern == pattern {
			return &ruleList[i]
		}
	}
	return nil
}

func copyMap(dest, src map[string]int) {
	for k, v := range src {
		dest[k] += v
	}
}
