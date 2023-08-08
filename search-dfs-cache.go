package graph

type cachedInput struct {
	processed  map[int]bool
	discovered map[int]bool
	parent     map[int]int
}

func (g *Graph) newCachedInput() cachedInput {
	var ci cachedInput
	n := g.numVertices
	ci.processed = make(map[int]bool, n)
	ci.discovered = make(map[int]bool, n)
	ci.parent = make(map[int]int, n)
	return ci
}

func (g *Graph) cachedDepthFirstSearchFrom(v int, ci cachedInput) (cachedInput, error) {
	var err error

	ci.discovered[v] = true // mark node as discovered
	err = g.ProcessNode(v)
	if err != nil {
		return ci, err
	}

	for _, new := range g.adj[v] {
		if !ci.discovered[new] {
			ci.parent[new] = v
			err := g.ProcessEdge(v, new)
			if err != nil {
				return ci, err
			}
			ci, err = g.cachedDepthFirstSearchFrom(new, ci)
			if err != nil {
				return ci, err
			}

		} else if !ci.processed[new] || g.directed {
			err := g.ProcessEdge(v, new)
			if err != nil {
				return ci, err
			}
		}

	}
	err = g.ProcessNodeLate(v)
	if err != nil {
		return ci, err
	}
	ci.processed[v] = true

	return ci, err
}
