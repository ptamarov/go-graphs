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
	var searchResult []int

	// decide how to process vertices during the search
	appendToSearchResult := func(v int) {
		searchResult = append(searchResult, v)
	}

	// count edges
	var edgeCount int
	increaseEdgeCounter := func(_, _ int) {
		edgeCount++
	}

	// do not process vertices after going through children
	processVertexLate := func(_ int) {}

	// initialize graph and perform BSF
	g, err = graph.New(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})
	if err != nil {
		t.Error(err)
	}
	g.BreadthFirstSearchFrom(1, appendToSearchResult, processVertexLate, increaseEdgeCounter)
	fmt.Println(searchResult)
	fmt.Println(edgeCount)

	// clear cache
	searchResult = []int{}
	edgeCount = 0

	// initialize directed graph and perform BSF
	g, err = graph.NewDirected(5, map[int][]int{0: {2, 4}, 1: {0}, 2: {1}, 3: {0}, 4: {3}})
	if err != nil {
		t.Error(err)
	}
	g.BreadthFirstSearchFrom(1, appendToSearchResult, processVertexLate, increaseEdgeCounter)
	fmt.Println(searchResult)
	fmt.Println(edgeCount)

	// Generate a 10 000 random graphs with 3 vertices and 2 edges in the Erdős–Rényi model.
	r := rand.New(rand.NewSource(time.Now().Unix()))
	results := make(map[int]int, 3)

	for i := 0; i < 10000; i++ {
		g, err := graph.NewRandom(r, 3, 2)
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
