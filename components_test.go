package graph

import "fmt"

func F() {
	g := graph{Adj: map[int][]int{0: {1, 2, 5}, 1: {0, 3, 4}, 2: {0, 3, 4}, 3: {1, 2, 5}, 4: {1, 2, 5}, 5: {0, 3, 4}, 6: {7, 8}, 7: {6, 8}, 8: {6, 7}}, NumVertices: 9, Directed: false}
	fmt.Println(g.ConnectedComponents())

}
