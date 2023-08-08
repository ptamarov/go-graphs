package graph

import "fmt"

// validate verifies that the graph data is valid by performing several checks:
//   - Check for a positive number of nodes.
//   - Check that adjacency list is of the correct lenght.
//   - Check that keys of adjacency list are precisely [0, g.NumVertices)
//   - Check that no node is out of bounds in the values of the arrays in the adj. list.
//   - Check that no node is its own neighbour.
//   - Check that no node has repeated neighbours.
func (g *Graph) validate() error {

	if g.numVertices == 0 {
		return fmt.Errorf("graph cannot have no nodes")
	}

	if len(g.adj) != g.numVertices {
		return fmt.Errorf("adjacency list has length %d but have %d nodes", len(g.adj), g.numVertices)
	}

	for current := 0; current < g.numVertices; current++ {
		var ok bool
		var children []int

		children, ok = g.adj[current]

		if !ok {
			return fmt.Errorf("node %d missing in adjacency list", current)
		}

		for _, child := range children {
			seen := make(map[int]bool)

			if child < 0 || child >= g.numVertices {
				return fmt.Errorf("node %d has neighbour %d out of bounds [%d]", current, child, g.numVertices)
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
