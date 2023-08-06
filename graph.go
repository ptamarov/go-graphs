package graph

type graph struct {
	numVertices int
	numEdges    int
	adj         map[int][]int
	directed    bool
	components  [][]int
}

func NewGraph(NumVertices int, Adj map[int][]int) (graph, error) {
	g := graph{numVertices: NumVertices, adj: Adj, directed: false}
	err := g.validate()
	if err != nil {
		var z graph
		return z, err
	}
	return g, nil
}

func NewDirGraph(NumVertices int, Adj map[int][]int) (graph, error) {
	g := graph{numVertices: NumVertices, adj: Adj, directed: true}
	err := g.validate()
	if err != nil {
		var z graph
		return z, err
	}
	return g, nil
}

// BreadthFirstSearch performs a BreadthFirstSearch in the graph starting from the input source node.
func (g *graph) BreadthFirstSearchFrom(source int) []int {
	return g.searchFrom(source, true)
}

// DepthFirstSearchFrom performs a DepthFirstSearchFrom in the graph starting from the input source node.
func (g *graph) DepthFirstSearchFrom(source int) []int {
	return g.searchFrom(source, false)
}

// ConnectedComponents counts the number of connected components of the graph.
func (g *graph) ConnectedComponents() [][]int {
	if len(g.components) == 0 {
		g.updateConnectedComponents(false)
	}
	return g.components
}

// EdgeNumber counts the number of edges in the graph by running iterated BFSs.
func (g *graph) EdgeNumber() int {
	if g.numEdges == 0 {
		g.updateConnectedComponents(true)
	}
	return g.numEdges
}

// DistanceFrom returns a map recording the distance of each node in the graph to the source node.
func (g *graph) DistanceFrom(source int) map[int]int {
	_, level := g.parent_and_level_from(source)
	return level
}

// ShortestDistanceTreeFrom returns a map which sends a node to a its parent in a shortest path from the source node. A parent value
// of -1 means there is no path unless key is source.
func (g *graph) ShortestDistanceTreeFrom(source int) map[int]int {
	parent, _ := g.parent_and_level_from(source)
	return parent
}

// ShortestPathsFrom returns a map which sends a node to a minimal path from the source node. An empty list means that no path exists.
func (g *graph) ShortestPathsFrom(source int) map[int][]int {
	parent, _ := g.parent_and_level_from(source)
	return treeToPaths(parent, source)
}
