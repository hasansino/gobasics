package sorting

import (
	"github.com/hasansino/gobasics/structures/binarytree"
)

func TreeSort(data []int) []int {
	bt := binarytree.NewBinaryTree(func(v1, v2 interface{}) bool {
		return v1.(int) <= v2.(int)
	})
	for _, v := range data {
		bt.Insert(v)
	}
	tmp := make([]int, 0, len(data))
	bt.Traverse(binarytree.LNR, func(n *binarytree.Node) bool {
		tmp = append(tmp, n.Value().(int))
		return true
	})
	return tmp
}
