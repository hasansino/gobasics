//
// Package graph implements graph data structure and basic operations with it.
//
// https://en.wikipedia.org/wiki/Graph_theory
// https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
// https://afteracademy.com/blog/introduction-to-graph-in-programming
//
package graph

import (
	"github.com/hasansino/gobasics/structures/queue"
	"github.com/hasansino/gobasics/structures/stack"
)

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

// Radius is minimum eccentricity of a graph
func (g *Graph) Radius() int { return 0 }

// Diameter is maximum eccentricity of a graph
func (g *Graph) Diameter() int { return 0 }

// CentralPoint is node whose eccentricity is equal to radius
func (g *Graph) CentralPoint() int { return 0 }

// Circumference returns number of edges in longest path of graph
func (g *Graph) Circumference() int { return 0 }

// Eccentricity returns maximum distance between given vertex n to any other
func (g *Graph) Eccentricity() int { return 0 }

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

// BreathFirstSearch of a graph
// https://en.wikipedia.org/wiki/Breadth-first_search
// https://afteracademy.com/blog/graph-traversal-breadth-first-search
func (g *Graph) BreathFirstSearch(v interface{}) int {
	var (
		q       = queue.NewLLQueue(g.size)
		visited = make(map[int]bool, g.size)
	)

	if err := q.Enqueue(0); err != nil {
		return -1
	}

	for !q.Empty() {
		n := q.Dequeue().(int)
		if visited[n] {
			continue
		}
		if g.nodes[n].value == v {
			return n
		}
		visited[n] = true
		for j := 0; j < g.size; j++ {
			if g.edges[n][j] != nil {
				if err := q.Enqueue(j); err != nil {
					return -1
				}
			}
		}
	}

	return -1
}

// DepthFirstSearch of graph
// https://en.wikipedia.org/wiki/Depth-first_search
// https://afteracademy.com/blog/graph-traversal-depth-first-search
func (g *Graph) DepthFirstSearch(v interface{}) int {
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
			return n
		}
		visited[n] = true
		for j := 0; j < g.size; j++ {
			if g.edges[n][j] != nil {
				s.Push(j)
			}
		}
	}

	return -1
}

// ShortestPathQueueItem with path already traversed for backtracking
type ShortestPathQueueItem struct {
	path []int
	node int
}

// ShortestPath from root node to node with given value
func (g *Graph) ShortestPath(v interface{}) []int {
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

// DijkstrasShortestDistances algorithm
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// https://afteracademy.com/blog/dijkstras-algorithm
func (g *Graph) DijkstrasShortestDistances(from int) []int {
	var (
		visited   = make([]bool, g.size)
		distances = make([]int, g.size)
	)

	for j := 0; j < g.size; j++ {
		distances[j] = -1 // undefined
	}
	distances[from] = 0

	for j := 0; j < g.size; j++ {
		next := g.dijkstrasShortestPathNextNode(visited, distances)
		visited[next] = true
		for k := 0; k < g.size; k++ {
			if !visited[k] && g.edges[next][k] != nil {
				if distances[k] < 0 || distances[next]+g.edges[next][k].weight < distances[k] {
					distances[k] = distances[next] + g.edges[next][k].weight
				}
			}
		}
	}

	return distances
}

func (g *Graph) dijkstrasShortestPathNextNode(visited []bool, distances []int) int {
	var minIdx = -1
	for j := 0; j < g.size; j++ {
		if !visited[j] && distances[j] >= 0 &&
			(minIdx == -1 || distances[j] < distances[minIdx]) {
			minIdx = j
		}
	}
	return minIdx
}
