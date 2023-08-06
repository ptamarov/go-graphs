# The `graphs` Go package

This a Go package to work with simple (un)directed graphs.

## Introduction

A graph object is defined by specifying:

1. a number of vertices `NumVertices`,
2. a map `Adj` assigning each integer in `[0,n)` to its array of neighbours,
3. optionally, a name `Name`.

## Initialization

To create a new undirected graph, you can use the buit-in method `NewGraph`, which 
returns an undirected graph or an error if the input data is invalid. The corresponding 
built-in method `NewDirGraph` returns a directed graph.

```go
func NewGraph(NumVertices int, Adj map[int][]int, Name string) (graph, error)
func NewDirGraph(NumVertices int, Adj map[int][]int, Name string) (graph, error)
```

A graph can also be created from a JSON file, using the in-built method `NewGraphFromJSON`.

```go
func NewGraphFromJSON(filepath string) (graph, error) 
```

The following example shows a valid JSON format to initalize an undirected graph; all fields
are required, although of course the `Name` field can be safely set to be the empty string.

```json
{
    "Name": "four star",
    "NumVertices": 5,
    "Directed": false,
    "Adj": {
        "0": [1, 2, 3, 4],
        "1": [0],
        "2": [0],
        "3": [0],
        "4": [0]
    }
}
```

## Methods 

The package support standard graph traversal algorithms to search through the graph. The method 

```go
func (g *graph) BreadthFirstSearchFrom(source int) []int
```

performs a breadth first search from the source vertex and returns the discovered vertices in
the resulting traversal order.  
