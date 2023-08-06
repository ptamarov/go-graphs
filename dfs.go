package graph

func (gr *graph) depthFirstSearchFromWithSubroutines(v int) {
	var (
		time     int
		finished bool
	)

	processed := make(map[int]bool, gr.numVertices)
	discovered := make(map[int]bool, gr.numVertices)
	entry_time := make(map[int]int, gr.numVertices)
	exit_time := make(map[int]int, gr.numVertices)
	parent := make(map[int]int, gr.numVertices)

	discovered[v] = true
	time++
	entry_time[v] = time

	finished = processVertexEarly(v)

	if finished {
		return
	}

	for _, y := range gr.adj[v] {
		if !discovered[y] {
			parent[y] = v
			processEdge(v, y)
		} else if !processed[y] || gr.directed {
			finished = processEdge(v, y)
			if finished { // subroutine could modify finished param
				return
			}
		}
	}

	processVertexLate(v)
	time++
	exit_time[v] = time
	processed[v] = true
}

func processEdge(v, y int) bool {
	// subroutine of an algorithm that uses DFS
	return true
}

func processVertexEarly(v int) bool {
	// subroutine of an algorithm that uses DFS
	return true
}

func processVertexLate(v int) {
	// subroutine of an algorithm that uses DFS
}
