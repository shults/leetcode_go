package t00160

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var aLen = listLen(headA)
	var bLen = listLen(headB)

	if aLen == bLen {
		//
	} else if aLen > bLen {
		for aLen != bLen {
			aLen--
			headA = headA.Next
		}
	} else {
		for bLen != aLen {
			bLen--
			headA = headB.Next
		}
	}

	for {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}
}

func listLen(node *ListNode) int {
	var result int

	for node != nil {
		result++
		node = node.Next
	}

	return result
}
