package remove_duplicates_from_sorted_list

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
func deleteDuplicates(head *ListNode) *ListNode {
	r := &ListNode{
		Val: head.Val,
	}
	r.AppendNext(head.Next)
	return r
}

func (l *ListNode) AppendNext(node *ListNode) *ListNode {
	if node.Next == nil {
		return l.Next
	}
	if l.Val != node.Val {
		l.Next = node
	}
	return l.Next.AppendNext(node.Next)
}
