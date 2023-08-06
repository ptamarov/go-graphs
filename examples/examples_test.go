package examples

import (
	"fmt"
	"testing"

	graph "github.com/ptamarov/go-graphs"
)

func TestExample(t *testing.T) {
	var g graph.Graph
	var err error

	g, err = graph.NewGraph(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})

	if err != nil {
		t.Error(err)
	}

	search := g.BreadthFirstSearchFrom(1)
	fmt.Println(search)
}
