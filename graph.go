package graph

type Graph struct {
	numVertices     int
	numEdges        int
	directed        bool
	adj             map[int][]int
	ProcessNode     func(int) error
	ProcessNodeLate func(int) error
	ProcessEdge     func(int, int) error
}

// New creates a new undirected graph or returns an error if the input data is invalid.
func New(NumVertices int, Adj map[int][]int) (Graph, error) {
	g := Graph{numVertices: NumVertices, adj: Adj, directed: false}
	err := g.validate()
	if err != nil {
		var z Graph
		return z, err
	}
	var count int
	for i := range g.adj {
		count += len(g.adj[i])
	}

	g.numEdges = count / 2
	g.ProcessNode = func(i int) error { return nil }
	g.ProcessEdge = func(i1, i2 int) error { return nil }
	g.ProcessNodeLate = func(i int) error { return nil }

	return g, nil

}

// NewDirected creates a new directed graph or returns an error if the
// input data is invalid.
func NewDirected(NumVertices int, Adj map[int][]int) (Graph, error) {
	g := Graph{numVertices: NumVertices, adj: Adj, directed: true}
	err := g.validate()
	if err != nil {
		var z Graph
		return z, err
	}

	var count int
	for i := range g.adj {
		count += len(g.adj[i])
	}

	g.numEdges = count
	g.ProcessEdge = func(i1, i2 int) error { return nil }
	g.ProcessNode = func(i int) error { return nil }
	g.ProcessNodeLate = func(i int) error { return nil }

	return g, nil
}

// Order returns the number of nodes in the graph.
func (g *Graph) Order() int {
	return g.numVertices
}

// Size returns the number of edges in the graph.
func (g *Graph) Size() int {
	return g.numEdges
}

// Degree returns the degree of a node. If the graph is oriented, it returns the outdegree.
func (g *Graph) Degree(node int) int {
	return len(g.adj[node])
}

// DistanceFrom returns a map recording the distance of each node in the graph to the source node.
func (g *Graph) DistanceFrom(source int) map[int]int {
	_, level := g.parentAndLevelFrom(source)
	return level
}

// ShortestDistanceTreeFrom returns a map which sends a node to a its parent in a shortest path from the source node. A parent value
// of -1 means there is no path unless key is source.
func (g *Graph) ShortestDistanceTreeFrom(source int) map[int]int {
	parent, _ := g.parentAndLevelFrom(source)
	return parent
}

// ShortestPathsFrom returns a map which sends a node to a minimal path from the source node. An empty list means that no path exists.
func (g *Graph) ShortestPathsFrom(source int) map[int][]int {
	parent, _ := g.parentAndLevelFrom(source)
	return treeToPaths(parent, source)
}
