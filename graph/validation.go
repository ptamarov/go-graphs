package graph

import "fmt"

// validate verifies that the graph data is valid by performing several checks:
//   - Check for a positive number of vertices.
//   - Check that adjacency list is of the correct lenght.
//   - Check that keys of adjacency list are precisely [0, g.NumVertices)
//   - Check that no node is out of bounds in the values of the arrays in the adj. list.
//   - Check that no node is its own neighbour.
//   - Check that no node has repeated neighbours.
func (g *graph) validate() error {

	if g.NumVertices == 0 {
		return fmt.Errorf("graph cannot have no vertices")
	}

	if len(g.Adj) != g.NumVertices {
		return fmt.Errorf("adjacency list has length %d but have %d vertices", len(g.Adj), g.NumVertices)
	}

	for current := 0; current < g.NumVertices; current++ {
		var ok bool
		var children []int

		children, ok = g.Adj[current]

		if !ok {
			return fmt.Errorf("node %d missing in adjacency list", current)
		}

		for _, child := range children {
			seen := make(map[int]bool)

			if child < 0 || child >= g.NumVertices {
				return fmt.Errorf("node %d has neighbour %d out of bounds [%d]", current, child, g.NumVertices)
			}

			if child == current {
				return fmt.Errorf("node %d is its own neighbour", current)
			}

			if _, ok := seen[child]; ok {
				return fmt.Errorf("node %d has repeated neighbour %d", current, child)
			}

			seen[child] = true
		}
	}

	return nil
}
