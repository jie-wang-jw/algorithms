package _2_linkedlist

/*This is a linked-list pointer reversal problem.
I keep two pointers: prev is the already reversed part, and cur is the node I am processing.
For each node, I save cur.Next first, reverse cur.Next to prev, then move both pointers forward.
At the end, prev becomes the new head.
prev：已经反转好的链表头
cur：当前正在处理的节点
next：原链表里 cur 后面的节点，先保存，防止断链

next := cur.Next // 先保存后面的节点，不然会断链
cur.Next = prev  // 当前节点反向指回前一个节点
prev = cur       // prev 往前走
cur = next       // cur 往前走
*/
//双指针
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// 递归版
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}
