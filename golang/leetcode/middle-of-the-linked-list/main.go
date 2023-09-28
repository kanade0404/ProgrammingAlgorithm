package middle_of_the_linked_list

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
func middleNode(head *ListNode) *ListNode {
	var vals []int
	n := head
	if n.Next == nil {
		return &ListNode{
			Val: n.Val,
		}
	}
	for {
		vals = append([]int{n.Val}, vals...)
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	results := make([]int, len(vals)+1, len(vals)+1)
	if len(vals)%2 == 0 {
		results = vals[:len(vals)/2]
	} else {
		results = vals[:(len(vals)+1)/2]
	}
	result := &ListNode{
		Val: results[0],
	}
	results = results[1:]
	for i := range results {
		tmp := result
		result = &ListNode{Val: results[i], Next: tmp}
	}
	return result
}
