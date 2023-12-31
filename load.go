package graph

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type testCase struct {
	Adj                  map[int][]int `json:"Adj"`
	NumVertices          int           `json:"NumVertices"`
	Directed             bool          `json:"Directed"`
	ExpectedBipartite    bool          `json:"ExpectedBipartite"`
	ExpectedSearches     map[int][]int `json:"ExpectedSearches"`
	ExpectedEdgeCount    int           `json:"ExpectedEdgecount"`
	ExpectedMatchingSize int           `json:"ExpectedMatchingSize"`
	ExpectedComponents   [][]int       `json:"ExpectedComponents"`
	ExpectedCycle        bool          `json:"ExpectedCycle"`
	// add more expected results for other tests
}

func loadTestsFromJSON(filepath string) map[string]testCase {
	filein, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer filein.Close()

	b, err := io.ReadAll(filein)
	if err != nil {
		log.Fatal(err)
	}

	testCaseByName := map[string]testCase{}

	err = json.Unmarshal(b, &testCaseByName)
	if err != nil {
		log.Fatal(err)
	}

	return testCaseByName
}

// NewFromJSON creates a new graph from a valid JSON input.
func NewFromJSON(filepath string) (Graph, error) {
	filein, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer filein.Close()

	b, err := io.ReadAll(filein)
	if err != nil {
		log.Fatal(err)
	}

	var t testCase
	err = json.Unmarshal(b, &t)
	if err != nil {
		log.Fatal(err)
	}

	var g Graph
	if t.Directed {
		g, err = New(t.NumVertices, t.Adj)
	} else {
		g, err = NewDirected(t.NumVertices, t.Adj)

	}

	if err != nil {
		return Graph{}, fmt.Errorf("while creating graph: %s", err)
	}

	return g, nil
}
