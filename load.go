package graph

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type testCase struct {
	graph
	result
}

type result struct {
	Bipartite bool          `json:"Bipartite"`
	Searches  map[int][]int `json:"Searches"`
	EdgeCount int           `json:"Edgecount"`
	// add more results for other tests
}

func loadTestsFromJSON(filepath string, queries []string) map[string]testCase {
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

func NewGraphFromJSON(filepath string) (graph, error) {
	var g graph
	var z graph

	filein, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer filein.Close()

	b, err := io.ReadAll(filein)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &g)
	if err != nil {
		log.Fatal(err)
	}

	err = g.validate()
	if err != nil {
		return z, err
	}

	return g, nil
}
