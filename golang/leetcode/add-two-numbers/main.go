package add_two_numbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1, l1Len := add(make([]byte, 0, 100), l1, 0)
	r2, l2Len := add(make([]byte, 0, 100), l2, 0)
	var (
		i          int
		acc1, acc2 byte
		isCarry    bool
		result     []byte
	)
	for {
		if i >= l1Len && i >= l2Len {
			if isCarry {
				result = append([]byte{1}, result...)
			}
			break
		}
		if i < l1Len {
			acc1 = r1[i]
		}
		if i < l2Len {
			acc2 = r2[i]
		}
		tmp := acc1 + acc2
		if isCarry {
			tmp += 1
		}
		if tmp > 9 {
			isCarry = true
			result = append([]byte{tmp % 10}, result...)
		} else {
			isCarry = false
			result = append([]byte{tmp}, result...)
		}
		acc1 = 0
		acc2 = 0
		i++
	}
	return createNode(result, len(result), nil)
}
func add(acc []byte, l *ListNode, count int) ([]byte, int) {
	if l.Next == nil {
		return append(acc, byte(l.Val)), count + 1
	} else {
		return add(append(acc, byte(l.Val)), l.Next, count+1)
	}
}
func createNode(num []byte, numLen int, l *ListNode) *ListNode {
	nLen := len(num)
	if nLen-1 == 0 {
		return &ListNode{
			Val:  int(num[0]),
			Next: l,
		}
	} else {
		if nLen == numLen {
			return createNode(num[1:], numLen, &ListNode{
				Val: int(num[0]),
			})
		} else {
			return createNode(num[1:], numLen, &ListNode{
				Val:  int(num[0]),
				Next: l,
			})
		}
	}
}
