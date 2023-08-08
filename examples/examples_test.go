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

	// add node to result when discovered
	appendToSearchResult := func(v int) error {
		searchResult = append(searchResult, v)
		return nil
	}

	// count edges as they are discovered
	var edgeCount int
	increaseEdgeCounter := func(_, _ int) error {
		edgeCount++
		return nil
	}

	// initialize graph and perform breadth first search
	g, err = graph.New(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})
	if err != nil {
		t.Error(err)
	}

	g.ProcessNode = appendToSearchResult
	g.ProcessEdge = increaseEdgeCounter
	_ = g.BreadthFirstSearchFrom(1)

	fmt.Println(searchResult) // [1 0 2 3 4]
	fmt.Println(edgeCount)    // 6

	// clear cache
	searchResult = []int{}
	edgeCount = 0

	// initialize directed graph and perform BSF
	g, err = graph.NewDirected(5, map[int][]int{0: {2, 4}, 1: {0}, 2: {1}, 3: {0}, 4: {3}})
	if err != nil {
		t.Error(err)
	}
	g.ProcessEdge = increaseEdgeCounter
	g.ProcessNode = appendToSearchResult
	g.BreadthFirstSearchFrom(1)
	fmt.Println(searchResult) // [1 0 2 4 3]
	fmt.Println(edgeCount)    // 6

	// Generate a 10 000 random graphs with 3 nodes and 2 edges in the Erdős–Rényi model.
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
	// Output: map[0:33?? 1:33?? 2:33??]

}
