package merge_two_sorted_lists

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var values []int
	cur1 := list1
	cur2 := list2
	for cur1 != nil || cur2 != nil {
		if cur1 != nil {
			values = append(values, cur1.Val)
			cur1 = cur1.Next
		}
		if cur2 != nil {
			values = append(values, cur2.Val)
			cur2 = cur2.Next
		}
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	var cur *ListNode
	root := &ListNode{}
	for i := range values {
		root.Val = values[i]
		tmp := root
		cur = &ListNode{
			Val: values[i],
		}
		tmp.Next = root
	}
	return root
}
