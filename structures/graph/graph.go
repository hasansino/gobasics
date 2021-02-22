package graph

// https://en.wikipedia.org/wiki/Graph_theory
// https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
// https://afteracademy.com/blog/introduction-to-graph-in-programming
// https://afteracademy.com/blog/graph-traversal-breadth-first-search
// https://afteracademy.com/blog/graph-traversal-depth-first-search

// Graph is ... well, a graph
type Graph struct {
	size  int
	nodes []Node
	edges [][]*Edge // Adjacency Matrix
}

// Node of a graph
type Node struct {
	value interface{}
}

// Edge of graph
type Edge struct {
	weight int
}

// NewGraph creates new graph of given size
func NewGraph(size int) *Graph {
	edges := make([][]*Edge, size)
	for i := 0; i < size; i++ {
		edges[i] = make([]*Edge, size)
	}
	return &Graph{
		size:  size,
		nodes: make([]Node, 0, size),
		edges: edges,
	}
}

// Eccentricity returns maximum distance between given n to all others
func (g *Graph) Eccentricity(n int) int { return 0 }

// Radius is minimum eccentricity of a graph
func (g *Graph) Radius() int { return 0 }

// Diameter is maximum eccentricity of a graph
func (g *Graph) Diameter() int { return 0 }

// CentralPoint is node whose eccentricity is equal to radius
func (g *Graph) CentralPoint() int { return 0 }

// Circumference returns number of edges in longest path of graph.
func (g *Graph) Circumference() int { return 0 }

// InsertNode into graph with given value
// Returns index of insertion
func (g *Graph) InsertNode(v interface{}) int {
	g.nodes = append(g.nodes, Node{value: v})
	return len(g.nodes)
}

// RemoveNode from graph
func (g *Graph) RemoveNode(i int) {}

// CreateEdge between i and j with weight w and direction d
func (g *Graph) CreateEdge(i, j, w int) {
	g.edges[i][j] = &Edge{weight: w}
}

// RemoveEdge between i and j
func (g *Graph) RemoveEdge(i, j int) {
	g.edges[i][j] = nil
}

// Adjacent tests if i and j have an edge
func (g *Graph) Adjacent(i, j int) bool {
	return g.edges[i][j] != nil
}
