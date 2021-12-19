package _18

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 4140, do(t, "sample.txt", false))
	assert.Equal(t, 3993, do(t, "sample.txt", true))
}

func TestInput(t *testing.T) {
	fmt.Println(do(t, "input.txt", false))
	fmt.Println(do(t, "input.txt", true))
}

func do(t *testing.T, fName string, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	//var sNumbers []sNumber
	var sNumbers []string
	var n int
	var line string

	for {
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil || line == "" {
			break
		}

		sNumbers = append(sNumbers, line)
		//nextSNum, _ := toSNumber(line)
		//sNumbers = append(sNumbers, nextSNum)
	}

	if !complicated {
		sumSNumbers := sumAndReduceAll(sNumbers)
		return magnitude(sumSNumbers)
	} else {
		maxMag := 0
		for i, v := range sNumbers {
			for j, w := range sNumbers[i+1:] {
				mag := magnitude(reduce(sum(v, w)))
				if mag > maxMag {
					fmt.Printf("New max at %d + %d\n%s PLUS %s = %d\n", i, i+j, v, w, mag)
					maxMag = mag
				}
				mag = magnitude(reduce(sum(w, v)))
				if mag > maxMag {
					fmt.Printf("New max at %d + %d\n%s PLUS %s = %d\n", i+j, i, w, v, mag)
					maxMag = mag
				}
			}
		}
		return maxMag
	}
}
