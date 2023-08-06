package graph

// ConnectedComponents returns a list of list with the connected components of the graph
func (g *graph) updateConnectedComponents(countEdges bool) {
	discovered := make(map[int]bool, g.numVertices)
	processed := make(map[int]bool, g.numVertices)
	components := [][]int{}

	if countEdges {
		g.numEdges = 0
	}

loop:
	for i := 0; i < g.numVertices; i++ {
		if discovered[i] {
			continue loop
		}
		discovered[i] = true

		newComponent := []int{}
		stack := []int{i}

		var current int
		for len(stack) != 0 {
			current, stack = stack[0], stack[1:] // dequeue for BFS
			for _, child := range g.adj[current] {
				if !processed[child] || g.directed {
					if countEdges {
						g.numEdges++ // a new edge goes from an unprocessed node to an processed node
					}
				}
				if !discovered[child] {
					stack = append(stack, child)
					discovered[child] = true
				}
			}
			processed[current] = true
			newComponent = append(newComponent, current)
		}
		components = append(components, newComponent)
	}

	g.components = components
}
