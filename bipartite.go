package graph

import "fmt"

const (
	uncolored = iota - 1
	white
	black
)

// FindTwoColoring attempts to find a two coloring of the graph. Returns a coloring in the form
// of a dictionary and nil if a two coloring exists, or a nil map and an error reporting a problematic
// edge if the two coloring could not be found.
func (g *Graph) FindTwoColoring() (map[int]int, error) {
	discovered := make(map[int]bool, g.numVertices)
	coloring := make(map[int]int, g.numVertices)

	// initialize all nodes as uncolored
	for i := 0; i < g.numVertices; i++ {
		coloring[i] = uncolored
	}

	// mark nodes as discovered in BFS
	markAsDiscovered := func(v int) {
		discovered[v] = true
	}

	// check that coloring is sound, else return error
	checkEdgeColoring := func(a, b int) error {
		if coloring[a] == coloring[b] {
			return fmt.Errorf("warning: not bipartite due to (%d, %d)", a, b)
		}
		coloring[b] = complement(coloring[a])
		return nil
	}

	for i := 0; i < g.numVertices; i++ {
		if !discovered[i] {
			coloring[i] = white
			err := g.BreadthFirstSearchFrom(i, markAsDiscovered, ignoreVertices, checkEdgeColoring)
			if err != nil {
				return map[int]int{}, err
			}
		}
	}
	return coloring, nil
}

func complement(color int) int {
	if color == white {
		return black
	}
	if color == black {
		return white
	}
	return uncolored
}
