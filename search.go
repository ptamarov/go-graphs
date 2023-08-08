package graph

import (
	"fmt"
)

// BreadthFirstSearchFrom performs a breadth first search from the source vertex and processes vertices and edges as instructed by the
// input functions. The vertices are processed in the traversal order, and are _processed late_ once all of its neighbours have been
// discovered in the search. Edges are processed as they appear from a new discovered vertex to a vertex that has not yet been
// processed.
func (g *Graph) BreadthFirstSearchFrom(source int, processVertex func(int), processVertexLate func(int), processEdge func(int, int)) {
	if source < 0 || source > g.numVertices {
		panic(fmt.Sprintf("source vertex %d out of range %d", source, g.numVertices))
	}

	processed := make(map[int]bool, g.numVertices)
	discovered := make(map[int]bool, g.numVertices)
	queue := []int{source}

	discovered[source] = true
	var current int

	for len(queue) != 0 {
		current, queue = queue[0], queue[1:]
		processVertex(current)
		processed[current] = true

		for _, child := range g.adj[current] {
			if !processed[child] || g.directed {
				processEdge(current, child)
			}
			if !discovered[child] {
				queue = append(queue, child)
				discovered[child] = true
			}
		}
		processVertexLate(current)
	}
}

// DepthFirstSearchFrom performs a DepthFirstSearchFrom in the graph starting from the input source node.
func (g *Graph) DepthFirstSearchFrom(source int) []int {
	ci := g.newCachedInput()
	ci = g.cachedDepthFirstSearchFrom(source, ci)
	return ci.searchResult
}
