package graph

import (
	"fmt"
)

// BreadthFirstSearchFrom performs a breadth first search from the source vertex
// and processes nodes and edges as instructed by the graph search functions.
//
// The nodes are processed in the traversal order, and are processed late once all of
// its neighbours have been discovered in the search. Edges are processed as they
// appear from a new discovered vertex to a vertex that has not yet been processed.
//
// If any of the search functions returns an error during the search, the search stops
// and the function returns this error.
func (g *Graph) BreadthFirstSearchFrom(source int) error {
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
		err := g.ProcessNode(current)
		if err != nil {
			return err
		}
		processed[current] = true

		for _, child := range g.adj[current] {
			if !processed[child] || g.directed {
				err := g.ProcessEdge(current, child)
				if err != nil {
					return err
				}
			}
			if !discovered[child] {
				queue = append(queue, child)
				discovered[child] = true
			}
		}
		g.ProcessNodeLate(current)
	}
	return nil
}

// DepthFirstSearchFrom performs a depth first search from the source vertex
// and processes nodes and edges as instructed by the graph search functions.
//
// The nodes are processed in the traversal order, and are processed late once all of
// its children have been processed in the traversal order. Edges are processed as they
// appear from a new discovered vertex to a vertex that has not yet been processed.
//
// If any of the search functions returns an error during the search, the search stops
// and the function returns this error.
func (g *Graph) DepthFirstSearchFrom(source int) error {
	ci := g.newCachedInput()
	_, err := g.cachedDepthFirstSearchFrom(source, ci)
	if err != nil {
		return err
	}
	return nil
}
