package _17

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	// target area: x=20..30, y=-10..-5
	target := targetArea{vector2{20, -10}, vector2{30, -5}}
	assert.Equal(t, 112, countValidLaunches(target))
}

func TestInput(t *testing.T) {
	// target area: x=235..259, y=-118..-62
	target := targetArea{vector2{235, -118}, vector2{259, -62}}
	fmt.Println(countValidLaunches(target))
}
