/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
	p := dummy
	extra := 0
	for l1 != nil && l2!= nil {
		origin := l1.Val + l2.Val + extra
		l1 = l1.Next
		l2 = l2.Next
		val := origin%10
		extra = origin/10

		p.Next = &ListNode{
			Val: val,
		}
		p = p.Next
	}

	l := l1
	if l1 == nil {
		l = l2
	}
	for l != nil {
		origin := l.Val + extra
		l = l.Next
		val := origin%10
		extra = origin/10

		p.Next = &ListNode{
			Val: val,
		}
		p = p.Next
	}
	if extra != 0 {
		p.Next = &ListNode{
			Val: extra,
		}
	}
	return dummy.Next
}
