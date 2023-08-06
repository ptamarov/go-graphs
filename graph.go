package graph

type Graph struct {
	numVertices int
	directed    bool
	adj         map[int][]int
	numEdges    int
	components  [][]int
}

func NewGraph(NumVertices int, Adj map[int][]int) (Graph, error) {
	g := Graph{numVertices: NumVertices, adj: Adj, directed: false}
	err := g.validate()
	if err != nil {
		var z Graph
		return z, err
	}
	return g, nil
}

func NewDirGraph(NumVertices int, Adj map[int][]int) (Graph, error) {
	g := Graph{numVertices: NumVertices, adj: Adj, directed: true}
	err := g.validate()
	if err != nil {
		var z Graph
		return z, err
	}
	return g, nil
}

// BreadthFirstSearch performs a BreadthFirstSearch in the graph starting from the input source node.
func (g *Graph) BreadthFirstSearchFrom(source int) []int {
	return g.searchFrom(source, true)
}

// DepthFirstSearchFrom performs a DepthFirstSearchFrom in the graph starting from the input source node.
func (g *Graph) DepthFirstSearchFrom(source int) []int {
	return g.searchFrom(source, false)
}

// ConnectedComponents counts the number of connected components of the graph.
func (g *Graph) ConnectedComponents() [][]int {
	if len(g.components) == 0 {
		g.updateConnectedComponents(false)
	}
	return g.components
}

// EdgeNumber counts the number of edges in the graph by running iterated BFSs.
func (g *Graph) EdgeNumber() int {
	if g.numEdges == 0 {
		g.updateConnectedComponents(true)
	}
	return g.numEdges
}

// DistanceFrom returns a map recording the distance of each node in the graph to the source node.
func (g *Graph) DistanceFrom(source int) map[int]int {
	_, level := g.parent_and_level_from(source)
	return level
}

// ShortestDistanceTreeFrom returns a map which sends a node to a its parent in a shortest path from the source node. A parent value
// of -1 means there is no path unless key is source.
func (g *Graph) ShortestDistanceTreeFrom(source int) map[int]int {
	parent, _ := g.parent_and_level_from(source)
	return parent
}

// ShortestPathsFrom returns a map which sends a node to a minimal path from the source node. An empty list means that no path exists.
func (g *Graph) ShortestPathsFrom(source int) map[int][]int {
	parent, _ := g.parent_and_level_from(source)
	return treeToPaths(parent, source)
}
