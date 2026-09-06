package _2_linkedlist

/*
题目描述 / Problem Description
给定一个链表的头节点 head 和整数 n，删除链表的倒数第 n 个节点，并返回删除后的头节点。
Given the head of a linked list and an integer n, remove the nth node from the end and return the resulting head.

解题思路 / Solution Approach
两个指针都从 dummy 出发，让 fast 先走 n 步，再同速移动，始终保持从 slow 到 fast 需要走 n 条 Next 边。
fast.Next==nil 时，fast 在最后一个节点。从 slow.Next 到 fast（两端都算）恰好有 n 个节点，所以 slow.Next 就是倒数第 n 个，slow 是它的前驱。
令 slow.Next=slow.Next.Next 跳过目标即可。dummy 为原头节点也提供了前驱，因此删除头节点无需特殊分支。
Start both pointers at dummy, advance fast n steps, then move both equally so fast remains n Next links ahead.
When fast.Next==nil, fast is the last node. From slow.Next through fast inclusive there are exactly n nodes,
making slow.Next the nth node from the end and slow its predecessor.
Skip the target with slow.Next=slow.Next.Next. The dummy also gives the original head a predecessor, avoiding a separate head-deletion case.

时间与空间复杂度 / Time and Space Complexity
L 为链表长度，题目参数 n 为倒数位置。时间 O(L)，fast 总共沿链表前进一次，slow 随后跟进。辅助空间 O(1)，虚拟头节点和指针数量固定。
L is the list length; parameter n is the position from the end. Time O(L),
with fast traversing the list once and slow following. Auxiliary space O(1) for a dummy node and pointers.
*/

/*
This is a linked-list two-pointer problem.
这是一个链表双指针问题。
I use a dummy node because the head might be removed.
因为头节点也可能被删除，所以使用 dummy 虚拟头节点。
I move fast n steps ahead: the distance is n links, not n nodes strictly between the pointers.
先让 fast 领先 n 步：距离是 n 条边，不是两者中间夹着 n 个节点。
Then I move both pointers together until fast reaches the last node.
然后两个指针一起移动，直到 fast 到达最后一个节点。
The n nodes from slow.Next through the last node form the final n nodes, so slow.Next is the target.
从 slow.Next 到最后节点恰好是末尾 n 个节点，因此 slow.Next 就是待删除节点。

脑内画面 / Mental picture (n=2):
slow -> target -> fast -> nil

	倒数第2   倒数第1
	2nd last  last

若改为 fast==nil 才停止，slow 也会多走一步到 target；不能只改停止条件而不改初始间距。
Stopping at fast==nil would also move slow one step onto target; the stopping condition must match the initial gap.
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// dummy makes deleting the original head identical to deleting any other node.
	// dummy 让“删除原头节点”和“删除普通节点”使用相同逻辑。
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// Advancing fast n links leaves exactly n nodes from slow.Next through fast.
	// fast 先走 n 条边，使 slow.Next 到 fast（含两端）恰好有 n 个节点。
	for range n {
		fast = fast.Next
	}

	// At the last node, those n nodes are the final n nodes, so slow.Next is the target.
	// fast 到最后节点时，这 n 个节点就是末尾 n 个，所以 slow.Next 是倒数第 n 个。
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
