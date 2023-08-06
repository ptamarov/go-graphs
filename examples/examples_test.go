package examples

import (
	"fmt"
	"testing"

	graph "github.com/ptamarov/go-graphs"
)

func TestExample(t *testing.T) {
	var g graph.Graph
	var err error

	g, err = graph.NewGraph(1, map[int][]int{1: {}})

	if err != nil {
		t.Error(err)
	}

	result := g.BreadthFirstSearchFrom(0)

	fmt.Println(result)
}
