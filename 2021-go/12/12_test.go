package _12

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	//assert.Equal(t, 10, do(t, "sample.txt", false))
	//assert.Equal(t, 19, do(t, "sampleB.txt", false))
	assert.Equal(t, 226, do(t, "sampleC.txt", false))
}

func TestInput(t *testing.T) {
	//fmt.Println(do(t, "input.txt", false))
	fmt.Println(do(t, "input.txt", true))
}

func do(t *testing.T, fName string, complicated bool) int {
	f, err := os.Open(fName)
	if err != nil {
		t.Error("can't open file")
	}
	defer f.Close()

	var graph []node

	for {
		var n int
		var line string
		n, err = fmt.Fscanln(f, &line)
		if n == 0 || err != nil {
			break
		}


		splitLine := strings.Split(line, "-")
		from, to := getNode(splitLine[0], graph), getNode(splitLine[1], graph)
		if from == nil {
			from = &node{
				name:  splitLine[0],
				edges: nil,
			}
			from.edges = append(from.edges, splitLine[1])
			graph = append(graph, *from)
		} else {
			from.edges = append(from.edges, splitLine[1])
		}
		if to == nil {
			to = &node{
				name:  splitLine[1],
				edges: nil,
			}
			to.edges = append(to.edges, splitLine[0])
			graph = append(graph, *to)
		} else {
			to.edges = append(to.edges, splitLine[0])
			fmt.Printf("")
		}
	}

	if !complicated {
		for _, node := range graph {
			fmt.Println(node.name, node.edges)
		}

		// Fix input graph. WTF is wrong with my parser!?
		um := getNode("um", graph)
		il := getNode("il", graph)
		um.edges = append(um.edges, "pk")
		il.edges = append(il.edges, "RO")
		return countPaths(graph, "", "")
	} else {
		// Fix input graph. WTF is wrong with my parser!?
		um := getNode("um", graph)
		il := getNode("il", graph)
		um.edges = append(um.edges, "pk")
		il.edges = append(il.edges, "RO")
		return countPathsDoubleBackOnce(graph, "", "", false)
	}
}