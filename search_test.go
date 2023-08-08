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
				g.ProcessNode = func(v int) error {
					searchResult = append(searchResult, v)
					return nil
				}
				g.BreadthFirstSearchFrom(source)
				got := searchResult
				want := expectedSearchResult

				if err := isEqual[int](got, want); err != nil {
					t.Errorf("while doing BFS: %s", err)
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

		for source, expectedSearchResult := range test.ExpectedSearches {
			t.Run(fmt.Sprintf(name+"/source=%d", source), func(t *testing.T) {
				searchResult := []int{}
				g.ProcessNode = func(v int) error {
					searchResult = append(searchResult, v)
					return nil
				}
				g.DepthFirstSearchFrom(source)
				got := searchResult
				want := expectedSearchResult
				if err := isEqual[int](got, want); err != nil {
					t.Errorf("while doing DFS: %s", err)
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
				t.Errorf("wanted %d edges but got %d for %s", want, got, name)
			}
		})
	}
}
