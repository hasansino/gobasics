package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//        234
	//       /   \
	//	4 - 6    546
	//   \   \   /
	//   33   654
	//    \  /
	//      2
	//                        0  1  2   3    4  5    6
	testNodes = []interface{}{4, 6, 33, 654, 2, 234, 546}
	testEdges = []struct {
		i, j, w int
	}{
		{0, 1, 1},
		{1, 0, 1},
		{1, 5, 1},
		{5, 1, 1},
		{5, 6, 1},
		{6, 5, 1},
		{0, 2, 1},
		{2, 0, 1},
		{1, 3, 1},
		{3, 1, 1},
		{6, 3, 1},
		{3, 6, 1},
		{2, 4, 1},
		{4, 2, 1},
		{3, 4, 1},
		{4, 3, 1},
	}
)

func TestGraph_Build(t *testing.T) {
	g := NewGraph(len(testNodes))
	for _, v := range testNodes {
		g.InsertNode(v)
	}
	for _, v := range testEdges {
		g.CreateEdge(v.i, v.j, v.w)
	}
}

func TestGraph_BFS(t *testing.T) {
	g := NewGraph(len(testNodes))
	for _, v := range testNodes {
		g.InsertNode(v)
	}
	for _, v := range testEdges {
		g.CreateEdge(v.i, v.j, v.w)
	}

	assert.Equal(t, 0, g.BreathFirstSearch(4))
	assert.Equal(t, 1, g.BreathFirstSearch(6))
	assert.Equal(t, 2, g.BreathFirstSearch(33))
	assert.Equal(t, 3, g.BreathFirstSearch(654))
	assert.Equal(t, 4, g.BreathFirstSearch(2))
	assert.Equal(t, 5, g.BreathFirstSearch(234))
	assert.Equal(t, 6, g.BreathFirstSearch(546))

	assert.Equal(t, -1, g.BreathFirstSearch(0))
	assert.Equal(t, -1, g.BreathFirstSearch(1))
	assert.Equal(t, -1, g.BreathFirstSearch(999))
}

func TestGraph_DFS(t *testing.T) {
	g := NewGraph(len(testNodes))
	for _, v := range testNodes {
		g.InsertNode(v)
	}
	for _, v := range testEdges {
		g.CreateEdge(v.i, v.j, v.w)
	}

	assert.Equal(t, 0, g.DepthFirstSearch(4))
	assert.Equal(t, 1, g.DepthFirstSearch(6))
	assert.Equal(t, 2, g.DepthFirstSearch(33))
	assert.Equal(t, 3, g.DepthFirstSearch(654))
	assert.Equal(t, 4, g.DepthFirstSearch(2))
	assert.Equal(t, 5, g.DepthFirstSearch(234))
	assert.Equal(t, 6, g.DepthFirstSearch(546))

	assert.Equal(t, -1, g.DepthFirstSearch(0))
	assert.Equal(t, -1, g.DepthFirstSearch(1))
	assert.Equal(t, -1, g.DepthFirstSearch(999))
}

func TestGraph_ShortestPath(t *testing.T) {
	g := NewGraph(len(testNodes))
	for _, v := range testNodes {
		g.InsertNode(v)
	}
	for _, v := range testEdges {
		g.CreateEdge(v.i, v.j, v.w)
	}
	assert.Equal(t, []int{0}, g.ShortestPath(4))
	assert.Equal(t, []int{0, 2, 4}, g.ShortestPath(2))
	assert.Equal(t, []int{0, 1, 3, 6}, g.ShortestPath(546))
}

func TestGraph_DijkstrasShortestDistances(t *testing.T) {
	var (
		//        234
		//       /   \
		//	4 - 6    546
		//   \   \   /
		//   33   654
		//    \  /
		//      2
		//                        0  1  2   3    4  5    6
		testNodes = []interface{}{4, 6, 33, 654, 2, 234, 546}
		testEdges = []struct {
			i, j, w int
		}{
			{0, 1, 2},
			{1, 0, 2},
			{1, 5, 2},
			{5, 1, 2},
			{5, 6, 2},
			{6, 5, 2},
			{6, 3, 2},
			{3, 6, 2},
			{3, 4, 2},
			{4, 3, 2},

			{0, 2, 9},
			{2, 0, 9},
			{1, 3, 9},
			{3, 1, 9},
			{2, 4, 9},
			{4, 2, 9},
		}
	)

	g := NewGraph(len(testNodes))
	for _, v := range testNodes {
		g.InsertNode(v)
	}
	for _, v := range testEdges {
		g.CreateEdge(v.i, v.j, v.w)
	}

	assert.Equal(t, []int{0, 2, 9, 8, 10, 4, 6}, g.DijkstrasShortestDistances(0))
	assert.Equal(t, []int{8, 6, 11, 0, 2, 4, 2}, g.DijkstrasShortestDistances(3))
}
