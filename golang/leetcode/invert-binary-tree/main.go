package invert_binary_tree

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

func invertTree(root *TreeNode) *TreeNode {
	head := &TreeNode{
		Val: root.Val,
	}

}

func dig(node *TreeNode) *TreeNode {
	if node.Left != nil {
		return dig(node.Left)
	}
	if node.Right != nil {
		return dig(node.Right)
	}
	return node
}
