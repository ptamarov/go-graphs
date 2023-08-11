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

### func `NewDirected`
```go
func NewDirected(NumVertices int, Adj map[int][]int) (graph, error)
```
`NewDirected` creates a new directed graph or returns an error if the input data is invalid. 

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

The full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/ptamarov/go-graphs).

## Methods and usage 

The package supports standard graph-traversal-based algorithms to query and analyse simple (un)directed graphs. A
graph stores three functions that are called when running a BFS or a DFS, which can return an `error` and in this
way stop the run. 

```go
type Graph struct {
	ProcessNode     func(int) error
	ProcessNodeLate func(int) error
	ProcessEdge     func(int, int) error
}
```

#### func `BreadthFirstSearchFrom`
```go
func (g *graph) BreadthFirstSearchFrom(source int) error
```
`BreadthFirstSearchFrom` runs a breadth first search from the source node, processing the nodes and edges 
according to the graph search functions. It returns the first error raised by any of the search functions if 
this happens during the run.

_Example_: the following code runs a breadth first search, returns the nodes encountered and counts the number 
of edges along the way.

```go
import (
	"fmt"

	graph "github.com/ptamarov/go-graphs"
)

func main() {
	var g graph.Graph
	var err error
	var searchResult []int

	// add node to result when discovered
	appendToSearchResult := func(v int) error {
		searchResult = append(searchResult, v)
		return nil
	}

	// count edges as they are discovered
	var edgeCount int
	increaseEdgeCounter := func(_, _ int) error {
		edgeCount++
		return nil
	}

	// initialize graph and perform breadth first search
	g, err = graph.New(5, map[int][]int{0: {1, 2, 3, 4}, 1: {0, 2}, 2: {0, 1}, 3: {0, 4}, 4: {0, 3}})
	if err != nil {
		fmt.Printf("while running bfs: %s\n", err)
	}

	g.ProcessNode = appendToSearchResult
	g.ProcessEdge = increaseEdgeCounter
	g.BreadthFirstSearchFrom(1)

	fmt.Println(searchResult) // [1 0 2 3 4]
	fmt.Println(edgeCount)    // 6
}
```

#### func `DepthFirstSearchFrom`
```go
func (g *Graph) DepthFirstSearchFrom(source int) error
```
`DepthFirstSearchFrom` runs a depth first search from the source node, processing the nodes and edges according to the graph 
search functions. It returns the first error raised by any of the search functions if this happens during the run. 

_Example (DepthFirstSearchFrom)_. The following example initializes a directed tree rooted at `0` and runs a DFS from this node.
It stops once it finds vertex `10`. Along the way, it counts the number of edges out of `0` that are encountered. 

```go 
import (
	"fmt"

	graph "github.com/ptamarov/go-graphs"
)

func main() {
	var g graph.Graph
	var err error

	wanted := 10

	findSpecificVertex := func(v int) error {
		if v == wanted {
			return fmt.Errorf("found wanted vertex %d", wanted)
		}
		return nil
	}

	// count edges out of 0
	var edgeCount int
	countSpecificEdges := func(v, w int) error {
		if v == 0 {
			edgeCount++
		}
		return nil
	}

	// initialize a rooted tree at 0 and set-up a DFS
	g, err = graph.NewDirected(12, map[int][]int{
		0:  {1, 2, 3},
		1:  {0},
		2:  {0, 4, 5},
		3:  {0, 6, 7, 8},
		4:  {2},
		5:  {2},
		6:  {3},
		7:  {3, 9, 10, 11},
		8:  {3},
		9:  {7},
		10: {7},
		11: {7},
	})
	if err != nil {
		fmt.Printf("while creating graph: %s\n", err)
	}

	g.ProcessNode = findSpecificVertex
	g.ProcessEdge = countSpecificEdges

	err = g.DepthFirstSearchFrom(0)
	if err != nil {
		fmt.Printf("dfs says: saw %d out edges and %s\n", edgeCount, err)
	} else {
		fmt.Printf("saw %d out edges but wanted vertex not found\n", edgeCount)
	}
	// Output: dfs says: saw 3 out edges and found wanted vertex 10
...
```

_Example (DepthFirstSearchFrom)_. The following (continued) example computes entry and exit times in this
rooted tree. 

```go
...
	var time int
	entryTimes := make(map[int]int, g.Order())
	exitTimes := make(map[int]int, g.Order())

	g.ProcessNode = func(v int) error {
		time++
		entryTimes[v] = time
		return nil
	}

	g.ProcessNodeLate = func(v int) error {
		time++
		exitTimes[v] = time
		return nil
	}

	g.DepthFirstSearchFrom(0)
	fmt.Println(entryTimes)
	fmt.Println(exitTimes)

	// Output:
	// map[0:1 1:2 2:4 3:10 4:5 5:7 6:11 7:13 8:21 9:14 10:16 11:18]
	// map[0:24 1:3 2:9 3:23 4:6 5:8 6:12 7:20 8:22 9:15 10:17 11:19]
}
```


#### func `FindTwoColoring`
```go
func (g *Graph) FindTwoColoring() (map[int]int, error)
```
`FindTwoColoring` attempts to find a two coloring of the graph. It returns
a map assigning each node to `0` or `1` if successful and no error, or an 
empty map and an error reporting a problematic edge if the attempt fails.

### Random graph generation

The package supports the generation of random graphs. 

#### func `NewRandom`
```go
func NewRandom(r *rand.Rand, n int, m int) (Graph, error) 
```

`NewRandom` generates a random graph with `n` nodes and `m` edges in
in the [Erdős–Rényi model](https://en.wikipedia.org/wiki/Erd%C5%91s%E2%80%93R%C3%A9nyi_model).

_Example (NewRandom)_. There are exactly 3 labelled graphs with 3 nodes and 2 edges, and each is uniquely
determined by a unique node of degree two. The following generates 10000 random graphs with 
3 nodes and 2 edges and prints the number of ocurrences of each. 

```go
package main 

import (
    "fmt"
	"rand"
	"time"

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
    // Output: map[0:33?? 1:33?? 2:33??] 
	// Missing digits may vary.
}
```