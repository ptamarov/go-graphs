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
func (g *graph) FindTwoColoring() (map[int]int, error) {
	var coloringError error
	var currentColor int

	discovered := make(map[int]bool, g.NumVertices)
	coloring := make(map[int]int, g.NumVertices)
	for i := 0; i < g.NumVertices; i++ {
		coloring[i] = uncolored
	}
	currentColor = white

loop:
	for i := 0; i < g.NumVertices; i++ {
		if discovered[i] {
			continue loop
		}

		discovered[i] = true

		currentLayer := []int{i}
		nextLayer := []int{}

		for len(currentLayer) != 0 {
			for _, parent := range currentLayer {
				coloring[parent] = currentColor
				for _, child := range g.Adj[parent] {
					if discovered[child] || g.Directed {
						coloringError = checkEdgeColoring(parent, child, coloring)
						if coloringError != nil {
							return nil, coloringError
						}
					}
					if !discovered[child] {
						nextLayer = append(nextLayer, child)
						discovered[child] = true
					}
				}
			}
			currentLayer, nextLayer = nextLayer, []int{}
			currentColor = complement(currentColor) // color changes from layer to layer
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

func checkEdgeColoring(a, b int, coloring map[int]int) error {
	if coloring[a] == coloring[b] {
		return fmt.Errorf("warning: not bipartite due to (%d, %d)", a, b)
	}
	return nil
}
