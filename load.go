package graph

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	// outer:
	// 	for name := range db {
	// 		var g graph
	// 		var r result

	// 		testdata := db[name]
	// 		// get graph data from test data
	// 		g, err = loadGraphFromMap(testdata)
	// 		if err != nil {
	// 			new := fmt.Errorf("skipping: error while loading graph for test case %s: %s", name, err)
	// 			errors = append(errors, new)
	// 			continue outer
	// 		}

	// 		// get results
	// 		r, err = loadResultFromJSONDict(testdata, queries)
	// 		if err != nil {
	// 			new := fmt.Errorf("skipping: error while loading results for test case %s: %s", name, err)
	// 			errors = append(errors, new)
	// 			continue outer
	// 		}

	// 		out = append(out, testCase{g, r})

	// }
	// return out, errors
}

func toIntSlice(anyslice []any) ([]int, error) {
	out := make([]int, len(anyslice))

	for key, value := range anyslice {
		if intvalue, ok := value.(float64); ok {
			out[key] = int(intvalue)
		} else {
			return []int{}, fmt.Errorf("entry value %#v at key %d is not JSON integer", value, key)
		}
	}
	return out, nil
}

func loadResultFromJSONDict(data map[string]any, queries []string) (result, error) {
	var value any
	var ok bool
	var r result
	r.Searches = make(map[int][]int)

	for _, query := range queries {
		switch query {
		case "Bipartite":
			// try to get result from bipartite test
			value, ok = data[query]
			if !ok {
				return r, fmt.Errorf("missing key %s", query)
			} else {
				if isBipartite, ok := value.(bool); ok {
					r.Bipartite = isBipartite
				} else {
					return r, fmt.Errorf("value at key %s is not JSON boolean", query)
				}
			}
		case "Searches":
			value, ok = data[query]
			if !ok {
				return r, fmt.Errorf("missing key %s", query)
			} else {
				if searches, ok := value.(map[string]any); ok {
					for key, value := range searches {
						intkey, err := strconv.Atoi(key)
						if err != nil {
							return r, fmt.Errorf("error when converting key %#v to integer", key)
						}
						if array, ok := value.([]any); !ok {
							return r, fmt.Errorf("value %#v for key %d in searches not JSON array", value, intkey)
						} else {
							bfsresult, err := toIntSlice(array)
							if err != nil {
								return r, fmt.Errorf("error while converting JSON array %#v to []int", array)
							}
							r.Searches[intkey] = bfsresult
						}
					}
				} else {
					return r, fmt.Errorf("value at key %s is not JSON struct", query)
				}
			}
		case "EdgeCount":
			// try to get result for edge counting test
			value, ok = data[query]
			if !ok {
				return r, fmt.Errorf("missing key %s", query)
			} else {
				if edgeCount, ok := value.(float64); ok {
					r.EdgeCount = int(edgeCount)
				} else {
					return r, fmt.Errorf("value at key %s is not JSON integer (float64)", query)
				}
			}
		default:
			return r, fmt.Errorf("query %#v is invalid", query)
		}
	}

	return r, nil
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

	db := map[string]any{}

	err = json.Unmarshal(b, &db)
	if err != nil {
		log.Fatal(err)
	}

	g, err = loadGraphFromMap(db)
	if err != nil {
		return z, err
	}
	err = g.validate()
	if err != nil {
		return z, err
	}

	return g, nil
}

// loadGraphFromMap attempts to load graph data into a Graph object
// Returns an error if load is unsuccessful.
func loadGraphFromMap(graphdata map[string]any) (graph, error) {
	var value any
	var ok bool
	var g graph

	// try to get graph name
	value, ok = graphdata["Name"]
	if !ok {
		return g, fmt.Errorf("missing key 'Name'")
	} else {
		if gname, ok := value.(string); ok {
			g.Name = gname
		} else {
			return g, fmt.Errorf("value for key 'Name' is not string")
		}
	}

	// try to get number of vertices
	value, ok = graphdata["NumVertices"]
	if !ok {
		return g, fmt.Errorf("error: missing key 'NumVertices'")
	} else {
		if numVertices, ok := value.(float64); ok {
			g.NumVertices = int(numVertices)
		} else {
			return g, fmt.Errorf("value for key 'NumVertices' is not JSON integer (float64)")
		}
	}

	// try to get directedness
	value, ok = graphdata["Directed"]
	if !ok {
		return g, fmt.Errorf("missing key 'Directed'")
	} else {
		if isDirected, ok := value.(bool); ok {
			g.Directed = isDirected
		} else {
			return g, fmt.Errorf("value for key 'Directed' is not JSON boolean (bool)")
		}
	}

	// try to get adjacency list
	value, ok = graphdata["Adj"]
	if !ok {
		return g, fmt.Errorf("missing key 'Adj'")
	} else {
		Adj := make(map[int][]int)

		if FloatAdj, ok := value.(map[string]any); ok {
			for key, value := range FloatAdj {
				nodeKey, err := strconv.Atoi(key)
				if err != nil {
					return g, fmt.Errorf("at 'Adj'; failed to convert key of adjacency list to integer (not a JSON int?): %s", err)
				}

				if slice, ok := value.([]any); ok {
					intSlice, err := toIntSlice(slice)
					if err != nil {
						return g, fmt.Errorf("at 'Adj'; error when converting JSON array to integer slice: %s", err)
					}
					Adj[nodeKey] = intSlice
				} else {
					return g, fmt.Errorf("at 'Adj'; value at node key %s is not JSON array", key)
				}
				g.Adj = Adj
			}
		} else {
			return g, fmt.Errorf("value at key 'Adj' is not JSON object")
		}
	}

	return g, nil
}
