package graph

type unionFind struct {
	componentPointer map[int]int
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

// ConnectedComponents returns an array of arrays containing the nodes in each connected component of the graph.
func (g *Graph) ConnectedComponents() [][]int {
	uf := unionFind{make(map[int]int)}
	for i := 0; i < g.numVertices; i++ {
		uf.componentPointer[i] = -1 // all components have size 1
	}

	for i := 0; i < g.numVertices; i++ {
		for _, j := range g.adj[i] {
			var iRep int
			var jRep int
			iRep = uf.computeRep(i)
			jRep = uf.computeRep(j)
			if iRep != jRep {
				uf.mergeComponents(iRep, jRep)
			}
		}
	}

	componentMap := make([][]int, g.numVertices)

	for i := 0; i < g.numVertices; i++ {
		rep := uf.computeRep(i)
		componentMap[rep] = append(componentMap[rep], i)
	}

	out := [][]int{}

	for i := range componentMap {
		if len(componentMap[i]) != 0 {
			out = append(out, componentMap[i])
		}
	}
	return out

}

func (uf *unionFind) computeRep(i int) int {
	// go down the tree until you get to the negative value (stores size of component)
	if uf.componentPointer[i] < 0 {
		return i
	}
	rep := uf.computeRep(uf.componentPointer[i])
	uf.componentPointer[i] = rep
	return rep
}

func (uf *unionFind) mergeComponents(a int, b int) {
	// get the component sizes
	aSize := -uf.componentPointer[a]
	bSize := -uf.componentPointer[b]

	// merge the smaller component to the larger one
	// this keeps trees shallow and makes run essentially linear
	if aSize < bSize {
		uf.componentPointer[a] = b
		uf.componentPointer[b] = -aSize - bSize
	} else {
		uf.componentPointer[b] = a
		uf.componentPointer[a] = -aSize - bSize
	}
}
