package graph

import (
	"fmt"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tests := loadTestsFromJSON("testdata/bfs-tests.json", []string{"Searches"})

	for _, test := range tests {
		for source, searchResult := range test.Searches {
			t.Run(fmt.Sprintf(test.Name+"/source=%d", source), func(t *testing.T) {
				got := test.BreadthFirstSearchFrom(source)
				want := searchResult

				for i := range got {
					if got[i] != want[i] {
						t.Errorf("BFS [%s] from %d: expected to get node %d at position %d but got node %d instead", test.Name, source, want[i], i, got[i])
					}
				}
			})
		}
	}
}

func TestEdgeCount(t *testing.T) {
	tests := loadTestsFromJSON("testdata/bfs-tests.json", []string{"EdgeCount"})

	for _, test := range tests {
		err := test.validate()
		if err != nil {
			t.Errorf("invalid input [%s]", test.Name)
		}
		got := test.EdgeNumber()
		want := test.EdgeCount
		if got != want {
			t.Errorf("Wanted %d edges but got %d for %s", want, got, test.Name)
		}
	}
}
