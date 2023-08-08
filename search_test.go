package graph

import (
	"fmt"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tests := loadTestsFromJSON("testdata/bfs-tests.json")

	for name, test := range tests {
		g, err := New(test.NumVertices, test.Adj)
		if err != nil {
			t.Errorf("while generating graph: %s", err)
		}

		for source, expectedSearchResult := range test.ExpectedSearches {
			t.Run(fmt.Sprintf(name+"/source=%d", source), func(t *testing.T) {
				searchResult := []int{}
				addToSearchResult := func(v int) {
					searchResult = append(searchResult, v)
				}
				g.BreadthFirstSearchFrom(source, addToSearchResult, ignoreVertices, ignoreEdges)
				got := searchResult
				want := expectedSearchResult
				for i := range got {
					if got[i] != want[i] {
						t.Errorf("BFS [%s] from %d: expected %d at position %d but got %d instead", name, source, want[i], i, got[i])
					}
				}
			})
		}
	}
}

func TestDepthFirstSearch(t *testing.T) {
	tests := loadTestsFromJSON("testdata/dfs-tests.json")

	for name, test := range tests {
		g, err := New(test.NumVertices, test.Adj)
		if err != nil {
			t.Errorf("while generating graph: %s", err)
		}

		for source, searchResult := range test.ExpectedSearches {
			t.Run(fmt.Sprintf(name+"/source=%d", source), func(t *testing.T) {

				got := g.DepthFirstSearchFrom(source)
				want := searchResult
				for i := range got {
					if got[i] != want[i] {
						// t.Errorf("DFS [%s] from %d: expected %d at position %d but got %d instead", name, source, want[i], i, got[i])
						t.Errorf("expected %v, got %v", want, got)
					}
				}
			})
		}
	}
}

func TestEdgeCount(t *testing.T) {
	tests := loadTestsFromJSON("testdata/bfs-tests.json")

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			g, err := New(test.NumVertices, test.Adj)
			if err != nil {
				t.Errorf("while generating graph: %s", err)
			}
			got, want := g.Size(), test.ExpectedEdgeCount
			if got != want {
				t.Errorf("Wanted %d edges but got %d for %s", want, got, name)
			}
		})
	}
}
