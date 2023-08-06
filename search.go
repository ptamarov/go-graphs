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

func (g *graph) searchFrom(source int, fifo bool) []int {
	if source < 0 || source > g.numVertices {
		panic(fmt.Sprintf("source vertex %d out of range %d", source, g.numVertices))
	}

	result := []int{}
	discovered := make(map[int]bool, g.numVertices)
	processed := make(map[int]bool, g.numVertices)
	discovered[source] = true
	array := []int{source}

	var current int

	for len(array) != 0 {
		current, array = popFrom(array, fifo)

		for _, child := range g.adj[current] {
			if !discovered[child] {
				array = append(array, child)
				discovered[child] = true // a node is discovered when its the first time it appears in bfs
			}
		}

		result = append(result, current)
		processed[current] = true // a node is processed once all its children have been discovered

	}

	return result
}
