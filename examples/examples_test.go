package examples

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	graph "github.com/ptamarov/go-graphs"
)

func TestExample(t *testing.T) {
	var g graph.Graph
	var err error

	// BFS in an undirected graph
	g, err = graph.NewGraph(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})

	if err != nil {
		t.Error(err)
	}

	search := g.BreadthFirstSearchFrom(1)
	fmt.Println(search)

	// BFS in a directed graph
	g, err = graph.NewDirGraph(5, map[int][]int{0: {2, 4}, 1: {0}, 2: {1}, 3: {0}, 4: {3}})

	if err != nil {
		t.Error(err)
	}

	search = g.BreadthFirstSearchFrom(1)
	fmt.Println(search)

	// Generate a 10 000 random graphs with 3 vertices and 2 edges in the Erdős–Rényi model.
	r := rand.New(rand.NewSource(time.Now().Unix()))
	results := make(map[int]int, 3)

	for i := 0; i < 10000; i++ {
		g, err := graph.RandomGraph(r, 3, 2)
		if err != nil {
			t.Error(err)
		}

		for i := 0; i < 3; i++ {
			results[i] += g.Degree(i) - 1
		}

	}

	fmt.Println(results)
	// Output: map[0:3300 1:3355 2:3345]

}
