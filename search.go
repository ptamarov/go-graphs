package graph

import (
	"fmt"
)

// BreadthFirstSearch performs a BreadthFirstSearch in the graph starting from the input source node.
func (g *Graph) BreadthFirstSearchFrom(source int) []int {
	if source < 0 || source > g.numVertices {
		panic(fmt.Sprintf("source vertex %d out of range %d", source, g.numVertices))
	}

	result := []int{}
	discovered := make(map[int]bool, g.numVertices)
	queue := []int{source}

	var current int

	for len(queue) != 0 {
		current, queue = queue[0], queue[1:]
		if !discovered[current] {
			discovered[current] = true
		}

		result = append(result, current)
		for _, child := range g.adj[current] {
			if !discovered[child] {

				queue = append(queue, child)
				discovered[child] = true
			}
		}
	}

	return result
}

// DepthFirstSearchFrom performs a DepthFirstSearchFrom in the graph starting from the input source node.
func (g *Graph) DepthFirstSearchFrom(source int) []int {
	ci := g.newCachedInput()
	ci = g.CachedDepthFirstSearchFrom(source, ci)
	return ci.searchResult
}
