package _2_linkedlist

/*
This is a linked-list two-pointer problem.
这是一个链表双指针问题。
I use a dummy node because the head might be removed.
因为头节点也可能被删除，所以使用 dummy 虚拟头节点。
I move fast n steps ahead first, so fast and slow keep a gap of n nodes.
先让 fast 领先 n 步，使 fast 和 slow 始终保持 n 个节点的距离。
Then I move both pointers together until fast reaches the last node.
然后两个指针一起移动，直到 fast 到达最后一个节点。
At that point, slow is right before the node we need to remove, so I skip slow.Next.
此时 slow 正好位于待删除节点前面，因此跳过 slow.Next 即可。

脑内画面
让 fast 先走 n 步，然后 slow 和 fast 一起走。
这样它们之间永远隔着 n 个节点：
我们希望 slow 停在“待删节点前一个节点”，不是待删节点本身
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// dummy makes deleting the original head identical to deleting any other node.
	// dummy 让“删除原头节点”和“删除普通节点”使用相同逻辑。
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// Move fast n steps ahead so the two pointers keep a gap of n nodes.
	// 先让 fast 走 n 步，使 fast 和 slow 之间始终保持 n 个节点的距离。
	for range n {
		fast = fast.Next
	}

	// Stop at fast.Next == nil so slow stays before the node to remove.
	// 在 fast.Next == nil 时停止，这样 slow 会停在待删除节点的前一个节点。
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Skip the target node by linking slow directly to the following node.
	// 让 slow 直接连接待删除节点的后一个节点，从而跳过目标节点。
	slow.Next = slow.Next.Next
	// The real result starts after dummy.
	// 真正的链表结果从 dummy.Next 开始。
	return dummy.Next
}
