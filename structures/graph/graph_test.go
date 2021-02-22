package graph

import "testing"

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
		{1, 5, 1},
		{5, 6, 1},
		{0, 2, 1},
		{1, 3, 1},
		{6, 3, 1},
		{2, 4, 1},
		{3, 4, 1},
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
