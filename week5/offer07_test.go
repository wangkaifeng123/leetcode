package week5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := new(TreeNode)
	headVal := preorder[0]
	head.Val = headVal
	for k, v := range inorder {
		if v == preorder[0] {
			head.Left = buildTree(preorder[1:k+1], inorder[:k])
			head.Right = buildTree(preorder[k+1:], inorder[k:])
			break
		}
	}
	return head
}
