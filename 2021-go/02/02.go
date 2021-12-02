package _2

import "fmt"

type command struct {
	cmd string
	value int
}

func moveAll(c []command) (int, int) {
	d, h := 0, 0
	for _, v := range c {
		switch v.cmd {
		case "forward":
			h += v.value
		case "up":
			d -= v.value
		case "down":
			d += v.value
		default:
			fmt.Errorf("help me I'm lost")
		}
	}
	return d, h
}

func moveAllAim(c []command) (int, int) {
	d, h, a := 0, 0, 0
	for _, v := range c {
		switch v.cmd {
		case "forward":
			h += v.value
			d += a*v.value
		case "up":
			a -= v.value
		case "down":
			a += v.value
		default:
			fmt.Errorf("help me I'm lost")
		}
	}
	return d, h
}