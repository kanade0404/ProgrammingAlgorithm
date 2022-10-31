package merge_nodes_in_between_zeros

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	var cur = head
	var result *ListNode
	var tmp int
	for cur.Next != nil {
		if cur.Val == 0 && tmp != 0 {
			result.Val = tmp
			tmp = 0
		} else {
			tmp += cur.Val
		}
		cur = cur.Next
	}
	return result
}
