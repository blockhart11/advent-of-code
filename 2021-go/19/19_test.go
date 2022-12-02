package _19

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	assert.Equal(t, 79, do(t, "sample.txt", false))
	assert.Equal(t, -1, do(t, "sample.txt", true))
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

	var scanners []scanner

	for {
		//first line is scanner line
		var padding string
		var scannerId, beaconId int
		var n int
		n, err = fmt.Fscanln(f, &padding, &padding, &scannerId, &padding)
		if n == 0 || err != nil {
			break
		}
		nextScanner := scanner{id: scannerId}
		if scannerId == 0 {
			nextScanner.location = zero()
		}
		for {
			var x, y, z int
			n, err = fmt.Fscanln(f, &x, &y, &z)
			if n == 0 || err != nil {
				break
			}
			nextScanner.beacons = append(nextScanner.beacons, beacon{beaconId, vector3{x, y, z}, -1, nil})
			beaconId++
		}
		scanners = append(scanners, nextScanner)
	}

	if !complicated {
		for _, s := range scanners {
			s.cacheDistances()
			fmt.Println(s)
		}
		return -1
	} else {
		return -1
	}
}
