package _2_linkedlist

/*
This is a linked-list two-pointer problem.
I use a dummy node because the head might be removed.
I move fast n steps ahead first, so fast and slow keep a gap of n nodes.
Then I move both pointers together until fast reaches the last node.
At that point, slow is right before the node we need to remove, so I skip slow.Next.

脑内画面
让 fast 先走 n 步，然后 slow 和 fast 一起走。
这样它们之间永远隔着 n 个节点：
我们希望 slow 停在“待删节点前一个节点”，不是待删节点本身
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	//这样 fast 和 slow 之间始终隔着 n 个节点
	for range n {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
