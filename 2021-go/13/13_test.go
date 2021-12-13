package _13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 17, do(t, "sample.txt", false))
	assert.Equal(t, 0, do(t, "sample.txt", true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", false))
	fmt.Println(do(t, "input.txt", true))
	// GJZGLUPJ
}

func do(t *testing.T, fName string, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var paper paper

	for {
		var n int
		var line string
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil || line == "" {
			break
		}

		var nextLine point
		splitLine := strings.Split(line, ",")
		nextLine.x, _ = strconv.Atoi(splitLine[0])
		nextLine.y, _ = strconv.Atoi(splitLine[1])

		paper.points = append(paper.points, nextLine)
	}

	// now do the folds
	for {
		var n int
		var line string
		var padding string
		// ex: "fold along y=7"
		n, err = fmt.Fscanln(f, &padding, &padding, &line)
		if n == 0 || err != nil {
			break
		}

		var nextFold fold
		splitLine := strings.Split(line, "=")
		nextFold.xAxis = splitLine[0] == "x"
		nextFold.value, _ = strconv.Atoi(splitLine[1])

		paper.folds = append(paper.folds, nextFold)
	}

	if !complicated {
		return foldOnceAndCount(&paper)
	} else {
		paper.foldAll()
		paper.print()
		return 0
	}
}