# The `graphs` Go package

This a Go package to work with simple (un)directed graphs.

## Initialization

A graph object is defined by specifying

1. a number of nodes `n`,
2. a map `Adj` assigning each integer in `[0,n)` to its array of neighbours.

A graph can be constructed with any of the following functions:

### func `New`
```go
func New(NumVertices int, Adj map[int][]int) (graph, error)
```
`New` creates a new undirected graph or returns an error if the input data is invalid. 

### func `NewFromJSON`
```go
func NewFromJSON(filepath string) (graph, error) 
```
`NewFromJSON` creates a graph from a JSON file or returns an error if the input data is invalid.

_Example (NewFromJSON)_. The following example shows a valid JSON format to initalize the butterfly graph
of the first example:
```json
{
    "NumVertices": 5,
    "Directed": false,
    "Adj": {
        "0": [1, 2, 3, 4],
        "1": [0, 2],
        "2": [0, 1],
        "3": [0, 4],
        "4": [0, 3]
    }
}
```

## Documentation 

The documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/ptamarov/go-graphs).

## Methods 

The package supports standard graph-traversal-based algorithms to query and analyse simple (un)directed graphs.

#### func `BreadthFirstSearchFrom`
```go
func (g *graph) BreadthFirstSearchFrom(source int, processNode func(int), processNodeLate func(int), processEdge func(int, int) error) error
```
`BreadthFirstSearchFrom` performs a breadth first search from the source node and processes nodes and edges as instructed by the
input functions. The nodes are processed in the traversal order, and are _processed late_ once all of its neighbours have been
discovered in the search. Edges are processed as they appear from a new discovered node to a node that has not yet been 
processed. If the edge processing function returns an error, it will stop the search and return the error encountered.

_Example_: the following performs a standard breadth first search and counts the number of edges of the initialized graph.

```go
	var g graph.Graph
	var err error
	var searchResult []int

	// add node to result when discovered
	appendToSearchResult := func(v int) {
		searchResult = append(searchResult, v)
	}

	// count edges as they are discovered
	var edgeCount int
	increaseEdgeCounter := func(_, _ int) {
		edgeCount++
	}

	// do no late node processing 
	processNodeLate := func(_ int) {}

	// initialize graph and perform breadth first search
	g, err = graph.New(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})
	if err != nil {
		t.Error(err)
	}
	g.BreadthFirstSearchFrom(1, appendToSearchResult, processNodeLate, increaseEdgeCounter)
	fmt.Println(searchResult) // [1 0 2 3 4]
	fmt.Println(edgeCount) // 6
```

#### func `DepthFirstSearchFrom`
```go
func (g *Graph) DepthFirstSearchFrom(source int) []int
```
`DepthFirstSearchFrom` performs a depth first search from the source node and 
returns the discovered nodes in the resulting traversal order.  

#### func `ConnectedComponents`
```go
func (g *Graph) ConnectedComponents() [][]int
```
`ConnectedComponents` returns an array of arrays, where each individual array 
corresponds to a single connected component of the graph.

#### func `ConnectedComponentOf`
```go
func (g *Graph) ConnectedComponentOf(source int) [][]int
```
`ConnectedComponentOf` returns an array containig all nodes in the connected
component of the source node.

#### func `FindTwoColoring`
```go
func (g *Graph) FindTwoColoring() (map[int]int, error)
```
`FindTwoColoring` attempts to find a two coloring of the graph. It returns
a map assigning each node to `0` or `1` if successful and no error, or an 
empty map and an error reporting a problematic edge if the attempt fails.

#### func `DistanceFrom`
```go
func (*Graph) DistanceFrom(source int) map[int]int
```
`DistanceFrom` returns a map that assigns each node of the graph to the
shortest distance to the source node. A value of `-1` reports the node is
unreachable from source. 

#### func `ShortestDistanceTreeFrom`
```go
func (g *Graph) ShortestDistanceTreeFrom(source int) map[int]int
```
`ShortestDistanceTreeFrom` returns a map assigning a node to its parent
in a shortest distance tree to the source. Any unreachable node is assigned
the parent parent `-1`; source is also assigned the parent `-1`.

#### func `ShortestPathsFrom`
```go
func (g *Graph) ShortestPathsFrom(source int) map[int][]int
```
`ShortestPathsFrom` returns a map sending a node to a shortest
path to the source node. An empty path indicates that the node
is unreachable from the source node.

#### func `NewRandom`
```go
func NewRandom(r *rand.Rand, n int, m int) (Graph, error) 
```

`NewRandom` generates a random graph with `n` nodes and `m` edges in
in the [Erdős–Rényi model](https://en.wikipedia.org/wiki/Erd%C5%91s%E2%80%93R%C3%A9nyi_model).

_Example_. There are exactly 3 labelled graphs with 3 nodes and 2 edges, and each is uniquely
determined by a unique node of degree two. The following generates 10000 random graphs with 
3 nodes and 2 edges and prints the number of ocurrences of each. 

```go
package main 

import (
    "fmt"

    graph "github.com/ptamarov/go-graphs"
)

func main() {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    
    results := make(map[int]int, 3)
    for i := 0; i < 10000; i++ {
        g, _ := graph.RandomGraph(r, 3, 2)
        for i := 0; i < 3; i++ {
            results[i] += g.Degree(i) - 1
        }
    }

    fmt.Println(results)
    // Output: map[0:3300 1:3355 2:3345]
}
```