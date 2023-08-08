package graph

func (g *Graph) updateConnectedComponents() {
	discovered := make(map[int]bool, g.numVertices)
	components := [][]int{}

	var newComponent []int
	appendToComponentAndDiscover := func(v int) {
		discovered[v] = true
		newComponent = append(newComponent, v)
	}

	for i := 0; i < g.numVertices; i++ {
		if !discovered[i] {
			discovered[i] = true
			g.BreadthFirstSearchFrom(i, appendToComponentAndDiscover, func(_ int) {}, func(_, _ int) {})
			components = append(components, newComponent)
			newComponent = []int{}
		}
	}
	g.components = components
}

// ConnectedComponentOf returns the connected component of the source node.
func (g *Graph) ConnectedComponentOf(source int) []int {
	var newComponent []int
	appendToComponent := func(v int) {
		newComponent = append(newComponent, v)
	}

	g.BreadthFirstSearchFrom(source, appendToComponent, func(_ int) {}, func(_, _ int) {})

	return newComponent
}
