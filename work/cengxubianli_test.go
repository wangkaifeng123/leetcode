package work

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree() *TreeNode {
	//[3,9,20,null,null,15,7]
	node6 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node5 := TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	node2 := TreeNode{
		Val:   20,
		Left:  &node5,
		Right: &node6,
	}
	node1 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	node0 := TreeNode{
		Val:   3,
		Left:  &node1,
		Right: &node2,
	}
	return &node0
}

var res []int

func TestBfs(t *testing.T) {
	root := NewTree()
	bfs(root)
	t.Log("res", res)
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	bfs(root.Left)
	bfs(root.Right)
	return
}

func TestDfs(t *testing.T) {
	root := NewTree()
	t.Log(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var count int          //层数
	nodes := []*TreeNode{} //放每一层的节点
	nodes = append(nodes, root)
	res = append(res, []int{root.Val})
	for len(nodes) != 0 {
		tempRes := []int{}
		count++
		length := len(nodes) //记录每层节电的长度
		for i := 0; i < length; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
				tempRes = append(tempRes, nodes[i].Left.Val)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
				tempRes = append(tempRes, nodes[i].Right.Val)
			}
		}

		if len(tempRes) != 0 {
			res = append(res, tempRes)
		}
		nodes = nodes[length:]
	}
	return res
}
