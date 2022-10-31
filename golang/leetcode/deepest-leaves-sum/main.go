package deepest_leaves_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/deepest-leaves-sum/discuss/463316/Golang-solution-100
func deepestLeavesSum(root *TreeNode) int {
	_, x := get(root, 0, 0, 0)
	return x
}

func get(root *TreeNode, level, highestLevel, sum int) (int, int) {
	if root.Left != nil {
		highestLevel, sum = get(root.Left, level+1, highestLevel, sum)
	}

	if root.Right != nil {
		highestLevel, sum = get(root.Right, level+1, highestLevel, sum)
	}

	if root.Left == nil && root.Right == nil {
		if level > highestLevel {
			highestLevel = level
			sum = root.Val
		} else if level == highestLevel {
			sum += root.Val
		}
	}

	return highestLevel, sum
}
