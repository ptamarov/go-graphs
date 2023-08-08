package graph

import (
	"fmt"
)

var ignoreEdges = func(_, _ int) error { return nil }
var ignoreVertices = func(_ int) {}

// BreadthFirstSearchFrom performs a breadth first search from the source vertex and processes nodes and edges as instructed by the
// input functions. The nodes are processed in the traversal order, and are _processed late_ once all of its neighbours have been
// discovered in the search. Edges are processed as they appear from a new discovered vertex to a vertex that has not yet been
// processed. If the edge processing function returns an error, the search stops and the function returns this error.
func (g *Graph) BreadthFirstSearchFrom(source int, processNode func(int), processNodeLate func(int), processEdge func(int, int) error) error {
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
		processNode(current)
		processed[current] = true

		for _, child := range g.adj[current] {
			if !processed[child] || g.directed {
				err := processEdge(current, child)
				if err != nil {
					return err
				}
			}
			if !discovered[child] {
				queue = append(queue, child)
				discovered[child] = true
			}
		}
		processNodeLate(current)
	}
	return nil
}

// DepthFirstSearchFrom performs a DepthFirstSearchFrom in the graph starting from the input source node.
func (g *Graph) DepthFirstSearchFrom(source int) []int {
	ci := g.newCachedInput()
	ci = g.cachedDepthFirstSearchFrom(source, ci)
	return ci.searchResult
}
