package graph

import "fmt"

// MinimalPathsTreeFrom returns a map which sends a node to its parent in a minimal distance tree with root at the source node.
func (g *Graph) parentAndLevelFrom(source int) (map[int]int, map[int]int) {
	if source > g.numVertices {
		panic(fmt.Sprintf("source vertex %d out of range %d", source, g.numVertices))
	}

	level := make(map[int]int, g.numVertices)
	parent := make(map[int]int, g.numVertices)
	discovered := make(map[int]bool, g.numVertices)

	for i := 0; i < g.numVertices; i++ {
		parent[i] = -1 // -1 -> unreachable
	}

	currentLayer := []int{source}

	var nextLayer []int

	currentLevel := 0
	for len(currentLayer) != 0 {
		for _, node := range currentLayer {
			discovered[node] = true
			level[node] = currentLevel
			for _, next := range g.adj[node] {
				if !discovered[next] {
					nextLayer = append(nextLayer, next)
					discovered[next] = true
					parent[next] = node
				}
			}
		}
		currentLayer = nextLayer
		nextLayer = []int{}
	}
	return parent, level
}

func treeToPaths(parent map[int]int, source int) map[int][]int {
	paths := make(map[int][]int) // paths[i] = [minimal distance path from i to source]
	for node := range parent {
		path := []int{}
		current := node
		for parent[current] != -1 {
			path = append(path, current)
			current = parent[current]
		}
		if parent[node] != -1 {
			path = append(path, source)
		}
		paths[node] = path

		if node == source {
			paths[node] = []int{node}
		}
	}
	return paths
}
