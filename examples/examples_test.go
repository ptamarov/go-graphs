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
	g.BreadthFirstSearchFrom(1)

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

func TestExample2(t *testing.T) {
	var g graph.Graph
	var err error

	wanted := 10

	findSpecificVertex := func(v int) error {
		if v == wanted {
			return fmt.Errorf("found wanted vertex %d", wanted)
		}
		return nil
	}

	// count edges out of 0
	var edgeCount int
	countSpecificEdges := func(v, w int) error {
		if v == 0 {
			edgeCount++
		}
		return nil
	}

	// initialize a rooted tree at 0 and run DSF
	g, err = graph.NewDirected(12, map[int][]int{
		0:  {1, 2, 3},
		1:  {0},
		2:  {0, 4, 5},
		3:  {0, 6, 7, 8},
		4:  {2},
		5:  {2},
		6:  {3},
		7:  {3, 9, 10, 11},
		8:  {3},
		9:  {7},
		10: {7},
		11: {7},
	})
	if err != nil {
		t.Error(err)
	}

	g.ProcessNode = findSpecificVertex
	g.ProcessEdge = countSpecificEdges

	err = g.DepthFirstSearchFrom(0)
	if err != nil {
		fmt.Printf("dfs says: saw %d out edges and %s\n", edgeCount, err)
	} else {
		fmt.Printf("saw %d out edges but wanted vertex not found\n", edgeCount)
	}
	// Output: dfs says: saw 3 out edges and found wanted vertex 10

	var time int
	entryTimes := make(map[int]int, g.Order())
	exitTimes := make(map[int]int, g.Order())

	g.ProcessNode = func(v int) error {
		time++
		entryTimes[v] = time
		return nil
	}

	g.ProcessNodeLate = func(v int) error {
		time++
		exitTimes[v] = time
		return nil
	}

	g.DepthFirstSearchFrom(0)
	fmt.Println(entryTimes)
	fmt.Println(exitTimes)

	// Output:
	// map[0:1 1:2 2:4 3:10 4:5 5:7 6:11 7:13 8:21 9:14 10:16 11:18]
	// map[0:24 1:3 2:9 3:23 4:6 5:8 6:12 7:20 8:22 9:15 10:17 11:19]
}
