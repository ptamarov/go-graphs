package graph

type cachedInput struct {
	processed    map[int]bool
	discovered   map[int]bool
	entryTime    map[int]int
	exitTime     map[int]int
	parent       map[int]int
	searchResult []int
	timeCounter  int
}

func (g *Graph) newCachedInput() cachedInput {
	var ci cachedInput
	n := g.numVertices
	ci.processed = make(map[int]bool, n)
	ci.discovered = make(map[int]bool, n)
	ci.entryTime = make(map[int]int, n)
	ci.exitTime = make(map[int]int, n)
	ci.parent = make(map[int]int, n)
	return ci
}

func (g *Graph) CachedDepthFirstSearchFrom(v int, ci cachedInput) cachedInput {
	ci.discovered[v] = true
	ci.entryTime[v] = ci.timeCounter
	ci.timeCounter++
	ci.searchResult = append(ci.searchResult, v)

	for _, new := range g.adj[v] {
		if !ci.discovered[new] {
			ci.parent[new] = v
			ci = g.CachedDepthFirstSearchFrom(new, ci)
		}
	}

	ci.processed[v] = true
	ci.timeCounter++
	ci.exitTime[v] = ci.timeCounter

	return ci
}
