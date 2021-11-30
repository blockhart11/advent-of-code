package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHoHoHo(t *testing.T) {
	e := "Ho ho hos!"
	assert.Equal(t, e, HoHoHo())
}