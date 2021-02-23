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

	assert.True(t, g.BreathFirstSearch(4))
	assert.True(t, g.BreathFirstSearch(6))
	assert.True(t, g.BreathFirstSearch(33))
	assert.True(t, g.BreathFirstSearch(654))
	assert.True(t, g.BreathFirstSearch(2))
	assert.True(t, g.BreathFirstSearch(234))
	assert.True(t, g.BreathFirstSearch(546))

	assert.False(t, g.BreathFirstSearch(0))
	assert.False(t, g.BreathFirstSearch(1))
	assert.False(t, g.BreathFirstSearch(999))
}

func TestGraph_DFS(t *testing.T) {
	g := NewGraph(len(testNodes))
	for _, v := range testNodes {
		g.InsertNode(v)
	}
	for _, v := range testEdges {
		g.CreateEdge(v.i, v.j, v.w)
	}

	assert.True(t, g.DepthFirstSearch(4))
	assert.True(t, g.DepthFirstSearch(6))
	assert.True(t, g.DepthFirstSearch(33))
	assert.True(t, g.DepthFirstSearch(654))
	assert.True(t, g.DepthFirstSearch(2))
	assert.True(t, g.DepthFirstSearch(234))
	assert.True(t, g.DepthFirstSearch(546))

	assert.False(t, g.DepthFirstSearch(0))
	assert.False(t, g.DepthFirstSearch(1))
	assert.False(t, g.DepthFirstSearch(999))
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
