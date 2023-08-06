# The `graphs` Go package

This a Go package to work with simple (un)directed graphs.

## Initialization

A graph object is defined by specifying

1. a number of vertices `NumVertices`,
2. a map `Adj` assigning each integer in `[0,n)` to its array of neighbours.

A graph can be constructed with any of the following functions:

### func `NewGraph`
```go
func NewGraph(NumVertices int, Adj map[int][]int) (graph, error)
```
`NewGraph` creates a new undirected graph or returns an error if the
input data is invalid. 

_Example (NewGraph)_. The following example creates a "butterfly" graph with five vertices
and performs a breadth first search starting from vertex `1`.
```go
package main 

import (
    "fmt"

    graph "github.com/ptamarov/go-graphs"
)

func main() {
	g, _ := graph.NewGraph(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})
    search := g.BreadthFirstSearchFrom(1)
    fmt.Println(search)
    // Output: [1 0 2 3 4]
}
```

### func `NewDirGraph`
```go 
func NewDirGraph(NumVertices int, Adj map[int][]int) (graph, error)
```
`NewDirGraph` returns a directed graph.

### func `NewGraphFromJSON`
```go
func NewGraphFromJSON(filepath string) (graph, error) 
```
`NewGraphFromJSON` creates a graph from a JSON file or returns an error if the
input data is invalid.

_Example (NewGraphFromJSON)_. The following example shows a valid JSON format to initalize the butterfly graph
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

## Methods 

The package support standard graph traversal algorithms to search through the graph.

#### func `BreadthFirstSearchFrom`
```go
func (g *graph) BreadthFirstSearchFrom(source int) []int
```
`BreadthFirstSearchFrom` performs a breadth first search from the source vertex
and returns the discovered vertices in the resulting traversal order. 

#### func `DepthFirstSearchFrom`
```go
func (g *Graph).DepthFirstSearchFrom(source int) []int
```
`DepthFirstSearchFrom` performs a depth first search from the source vertex and 
returns the discovered vertices in the resulting traversal order.  

#### func `ConnectedComponents`
```go
func (g *Graph).ConnectedComponents() [][]int
```
`ConnectedComponents` returns an array of arrays, where each individual array 
corresponds to a single connected component of the graph.

#### func `DistanceFrom`
```go
func (*Graph).DistanceFrom(source int) map[int]int
```
`DistanceFrom` returns a map that assigns each vertex of the graph to the
shortest distance to the source vertex. A value of `-1` reports the vertex is
unreachable from source. 

#### func `NumberOfEdges`
```go
func (g *Graph).NumberOfEdges() int
```
`NumberOfEdges` returns the number of edges of the graph.

#### func `FindTwoColoring`
```go
func (g *Graph).FindTwoColoring() (map[int]int, error)
```
`FindTwoColoring` attempts to find a two coloring of the graph. It returns
a map assigning each node to `0` or `1` if successful and no error, or an 
empty map and an error reporting a problematic edge if the attempt fails.

#### func `ShortestDistanceTreeFrom`
```go
func (g *Graph).ShortestDistanceTreeFrom(source int) map[int]int
```
`ShortestDistanceTreeFrom` returns a map assigning a vertex to its parent
in a shortest distance tree to the source. Any unreachable vertex is assigned
the parent parent `-1`; source is also assigned the parent `-1`.

#### func `ShortestPathsFrom`
```go
func (g *Graph).ShortestPathsFrom(source int) map[int][]int
```
`ShortestPathsFrom` returns a map sending a node to a shortest
path to the source node. An empty path indicated that the node
is unreachable from the source node.
