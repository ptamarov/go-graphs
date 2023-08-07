package graph

// ConnectedComponents returns a list of list with the connected components of the graph
func (g *Graph) updateConnectedComponents() {
	discovered := make(map[int]bool, g.numVertices)
	components := [][]int{}

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
			current, stack = stack[0], stack[1:]
			for _, child := range g.adj[current] {
				if !discovered[child] {
					stack = append(stack, child)
					discovered[child] = true
				}
			}
			newComponent = append(newComponent, current)
		}
		components = append(components, newComponent)
	}

	g.components = components
}

// ConnectedComponentOf returns the connected component of the source node.
func (g *Graph) ConnectedComponentOf(source int) []int {
	discovered := make(map[int]bool, g.numVertices)
	component := []int{}

	discovered[source] = true
	stack := []int{source}

	var current int
	for len(stack) != 0 {
		current, stack = stack[0], stack[1:]
		for _, child := range g.adj[current] {
			if !discovered[child] {
				stack = append(stack, child)
				discovered[child] = true
			}
		}
		component = append(component, current)
	}
	return component
}
