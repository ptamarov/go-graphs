package graph

import (
	"fmt"
)

func popFrom(slice []int, fifo bool) (int, []int) {
	if len(slice) == 0 {
		panic("pop from empty slice")
	}
	var out int

	if fifo {
		out, slice = slice[0], slice[1:]
	} else {
		out, slice = slice[len(slice)-1], slice[:len(slice)-1]
	}
	return out, slice
}

func (g *Graph) searchFrom(source int, fifo bool) []int {
	if source < 0 || source > g.numVertices {
		panic(fmt.Sprintf("source vertex %d out of range %d", source, g.numVertices))
	}

	result := []int{}
	discovered := make(map[int]bool, g.numVertices)
	discovered[source] = true
	array := []int{source}

	var current int

	for len(array) != 0 {
		current, array = popFrom(array, fifo)
		result = append(result, current)

		var temp []int
		for _, child := range g.adj[current] {
			if !discovered[child] {
				temp = append(temp, child)
				discovered[child] = true
			}
		}
		if fifo {
			array = append(array, temp...)
		} else {
			array = append(array, reverse(temp)...)
		}
	}

	return result
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
