/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    var reverse func(*ListNode) (*ListNode, *ListNode) 
	reverse = func(l *ListNode) (*ListNode, *ListNode) {
		head := &ListNode{}
		tail := l
		for l != nil {
			cur := l
			l = l.Next

			cur.Next = head.Next
			head.Next = cur
		}
		return head.Next, tail
	}

	dummy := &ListNode{}
	p := dummy
	q := head
	cnt := 0
	from := head
	for q != nil {
		cur := q
		q = q.Next
		cnt++
		
		if cnt == k {
			cur.Next = nil
			head, tail := reverse(from)
			p.Next = head
			p = tail

			cnt = 0
			from = q
		}
	}
	if cnt > 0 {
		p.Next = from
	}

	return dummy.Next
}
