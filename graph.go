package graph

type graph struct {
	NumVertices int `json:"NumVertices"`
	numEdges    int
	Adj         map[int][]int `json:"Adj"`
	Directed    bool          `json:"Directed"`
	Name        string        `json:"Name"`
	components  [][]int
}

func NewGraph(NumVertices int, Adj map[int][]int, Name string) (graph, error) {
	g := graph{NumVertices: NumVertices, Adj: Adj, Directed: false, Name: Name}
	err := g.validate()
	if err != nil {
		var z graph
		return z, err
	}
	return g, nil
}

func NewDirGraph(NumVertices int, Adj map[int][]int, Name string) (graph, error) {
	g := graph{NumVertices: NumVertices, Adj: Adj, Directed: true, Name: Name}
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
