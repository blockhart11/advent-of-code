package _14

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 1588, do(t, "sample.txt", 10))
	assert.Equal(t, 2188189693529, do(t, "sample.txt", 40))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", 10))
	fmt.Println(do(t, "input.txt", 40))
}

func do(t *testing.T, fName string, cycleCount int) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var template polymerTemplate
	var n int
	var line, padding, insertion string

	n, err = fmt.Fscanln(f, &line)
	if n == 0 || err != nil {
		t.Error("can't parse first line")
	}

	template.startState = line

	// parse the empty line next
	n, err = fmt.Fscanln(f, &line)
	//fmt.Println("expected empty line error:", err)

	for {
		n, err = fmt.Fscanln(f, &line, &padding, &insertion)
		if n == 0 || err != nil || line == "" {
			break
		}

		template.rules = append(template.rules, rule{
			pattern: line,
			insert:  insertion,
		})
	}

	return mostLessLeastWithCache(template.startState, template.rules, cycleCount)
}
