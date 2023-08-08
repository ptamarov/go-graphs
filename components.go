package graph

func (g *Graph) updateConnectedComponents() {
	discovered := make(map[int]bool, g.numVertices)
	components := [][]int{}

	var newComponent []int

	// add nodes to new component
	g.ProcessNode = func(v int) error {
		discovered[v] = true
		newComponent = append(newComponent, v)
		return nil
	}

	for i := 0; i < g.numVertices; i++ {
		if !discovered[i] {
			discovered[i] = true
			g.BreadthFirstSearchFrom(i)
			components = append(components, newComponent)
			newComponent = []int{}
		}
	}
	g.components = components
}

// ConnectedComponentOf returns the connected component of the source node.
func (g *Graph) ConnectedComponentOf(source int) []int {
	var newComponent []int

	g.ProcessNode = func(v int) error {
		newComponent = append(newComponent, v)
		return nil
	}

	g.BreadthFirstSearchFrom(source)
	return newComponent
}
