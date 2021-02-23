package graph

import (
	"github.com/hasansino/gobasics/structures/queue"
	"github.com/hasansino/gobasics/structures/stack"
)

// https://en.wikipedia.org/wiki/Graph_theory
// https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
// https://afteracademy.com/blog/introduction-to-graph-in-programming

// Graph is ... well, a graph
type Graph struct {
	size  int
	nodes []*Node
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
		nodes: make([]*Node, 0, size),
		edges: edges,
	}
}

// Eccentricity returns maximum distance between given n to all others
func (g *Graph) Eccentricity(n int) int { return 0 } // @TODO

// Radius is minimum eccentricity of a graph
func (g *Graph) Radius() int { return 0 } // @TODO

// Diameter is maximum eccentricity of a graph
func (g *Graph) Diameter() int { return 0 } // @TODO

// CentralPoint is node whose eccentricity is equal to radius
func (g *Graph) CentralPoint() int { return 0 } // @TODO

// Circumference returns number of edges in longest path of graph.
func (g *Graph) Circumference() int { return 0 } // @TODO

// InsertNode into graph with given value
// Returns index of insertion
func (g *Graph) InsertNode(v interface{}) int {
	g.nodes = append(g.nodes, &Node{value: v})
	return len(g.nodes)
}

// RemoveNode from graph
func (g *Graph) RemoveNode(i int) {
	g.nodes[i] = nil
	for j := 0; j < g.size; j++ {
		g.edges[i][j] = nil
	}
	for j := 0; j < g.size; j++ {
		g.edges[j][i] = nil
	}
}

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
	return g.edges[i][j] != nil || g.edges[j][i] != nil
}

// https://en.wikipedia.org/wiki/Breadth-first_search
// https://afteracademy.com/blog/graph-traversal-breadth-first-search

// BreathFirstSearch of a graph
func (g *Graph) BreathFirstSearch(v interface{}) bool {
	if g.nodes[0] == nil {
		return false // empty graph
	}

	var (
		q       = queue.NewLLQueue(g.size)
		visited = make(map[int]bool, g.size)
	)

	if err := q.Enqueue(0); err != nil {
		return false
	}

	for !q.Empty() {
		n := q.Dequeue().(int)
		if visited[n] {
			continue
		}
		if g.nodes[n].value == v {
			return true
		}
		visited[n] = true
		for j := 0; j < g.size; j++ {
			if g.edges[n][j] != nil {
				if err := q.Enqueue(j); err != nil {
					return false
				}
			}
		}
	}

	return false
}

// https://en.wikipedia.org/wiki/Depth-first_search
// https://afteracademy.com/blog/graph-traversal-depth-first-search

// DepthFirstSearch of graph
func (g *Graph) DepthFirstSearch(v interface{}) bool {
	if g.nodes[0] == nil {
		return false // empty graph
	}

	var (
		s       = stack.NewLLStack()
		visited = make(map[int]bool, g.size)
	)

	s.Push(0)

	for !s.Empty() {
		n := s.Pop().(int)
		if visited[n] {
			continue
		}
		if g.nodes[n].value == v {
			return true
		}
		visited[n] = true
		for j := 0; j < g.size; j++ {
			if g.edges[n][j] != nil {
				s.Push(j)
			}
		}
	}

	return false
}

// ShortestPathQueueItem with path already traversed for backtracking
type ShortestPathQueueItem struct {
	path []int
	node int
}

// ShortestPath from root node to node with given value
func (g *Graph) ShortestPath(v interface{}) []int {
	if g.nodes[0] == nil {
		return nil // empty graph
	}

	var (
		q       = queue.NewLLQueue(g.size)
		visited = make(map[int]bool, g.size)
	)

	if err := q.Enqueue(ShortestPathQueueItem{
		node: 0, path: make([]int, 0),
	}); err != nil {
		return nil
	}

	for !q.Empty() {
		qi := q.Dequeue().(ShortestPathQueueItem)
		if visited[qi.node] {
			continue
		}
		if g.nodes[qi.node].value == v {
			qi.path = append(qi.path, qi.node)
			return qi.path
		}
		visited[qi.node] = true
		for j := 0; j < g.size; j++ {
			if g.edges[qi.node][j] != nil {
				if err := q.Enqueue(ShortestPathQueueItem{
					node: j,
					path: append(qi.path, qi.node),
				}); err != nil {
					return nil
				}
			}
		}
	}

	return nil
}
