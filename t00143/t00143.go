package t00143

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(items ...int) *ListNode {
	var root = new(ListNode)
	var tail = root

	for _, item := range items {
		tail.Next = new(ListNode)
		tail.Next.Val = item
		tail = tail.Next
	}

	return root.Next
}

func (l *ListNode) ToSlice() []int {
	var res []int

	curr := l

	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	return res
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	oldHead := head

	tail, length := reverseAndCalc(head)
	mid := length / 2

	var root = new(ListNode)
	var curr = root

	for i := 0; i < mid; i++ {
		var node = new(ListNode)
		node.Val = head.Val
		curr.Next = node
		curr = node

		var node2 = new(ListNode)
		node2.Val = tail.Val
		curr.Next = node2
		curr = node2

		head = head.Next
		tail = tail.Next
	}

	if length&1 == 1 {
		var node = new(ListNode)
		node.Val = head.Val
		curr.Next = node
	}

	*oldHead = *root.Next
}

func reverseAndCalc(head *ListNode) (*ListNode, int) {
	var length int

	var res *ListNode

	for head != nil {
		var n ListNode
		n.Val = head.Val
		n.Next = res
		res = &n
		length++
		head = head.Next
	}

	return res, length
}
