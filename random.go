package graph

import (
	"fmt"
	"math/rand"
)

type edge struct {
	s int
	t int
}

// RandomGraph generates a random graph among those graphs with n vertices and m edges in the Erdős–Rényi model.
// In this model each graph with n vertices and m edge is equiprobable.
func RandomGraph(r *rand.Rand, n int, m int) (Graph, error) {

	if 2*m > n*(n-1) {
		panic(fmt.Sprintf("too many edges (%d) for a graph with %d nodes", m, n))
	}

	Adj := make(map[int][]int, n)

	univ := [](edge){}

	for i := 0; i < n; i++ {
		Adj[i] = []int{}
		for j := 0; j < n; j++ {
			if j > i {
				univ = append(univ, edge{i, j})
			}
		}
	}

	for i := 0; i < m; i++ {
		var e edge
		n := len(univ)
		randIndex := r.Intn(n)

		univ[n-1], univ[randIndex] = univ[randIndex], univ[n-1]

		e, univ = univ[n-1], univ[:n-1]
		s, t := e.s, e.t
		Adj[s] = append(Adj[s], t)
		Adj[t] = append(Adj[t], s)
	}

	g, err := NewGraph(n, Adj)

	if err != nil {
		return Graph{}, err
	}

	return g, nil
}
