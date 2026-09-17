package _2_linkedlist

/*
题目描述 / Problem Description
给定一个链表的头节点 head 和整数 n，删除链表的倒数第 n 个节点，并返回删除后的头节点。
Given the head of a linked list and an integer n, remove the nth node from the end and return the resulting head.
*/

// 1. Fast and slow pointers: one pass, with fast kept exactly n links ahead of slow.
// 1. 快慢指针：一次遍历，让 fast 始终领先 slow 恰好 n 条边。
// Time: O(L), Space: O(1).
// 时间复杂度：O(L)，空间复杂度：O(1)。
//
// 要求 1<=n<=链表长度。从 dummy 出发让 fast 先走 n 步，两指针始终相距 n 条边。
// Require 1<=n<=list length; starting at dummy, advance fast n links and retain that gap.
// fast 停在尾节点时，slow 正好是倒数第 n 个节点的前驱；修改 slow.Next，dummy 统一处理删头。
// When fast reaches the tail, slow precedes the nth node from the end; dummy makes head deletion identical.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

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

// 2. Two passes: measure the length first, then walk to the target's predecessor.
// 2. 两次遍历：先求出链表长度，再走到待删除节点的前驱。
// Time: O(L), Space: O(1).
// 时间复杂度：O(L)，空间复杂度：O(1)。
//
// 要求 n 在有效范围内。倒数第 n 个节点的零基下标是 length-n，从 dummy 走这么多步恰到其前驱。
// Require a valid n; the target's zero-based index is length-n, so that many steps from dummy reach its predecessor.
func removeNthFromEndLength(head *ListNode, n int) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	steps := length - n
	for range steps {
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	return dummy.Next
}

// 3. Stack: push every node, then pop n entries so the top becomes the predecessor.
// 3. 栈解法：把所有节点入栈，再弹出 n 个，栈顶就是待删除节点的前驱。
// Time: O(L), Space: O(L) for the stack.
// 时间复杂度：O(L)，空间复杂度：O(L)，由栈产生。
//
// 栈按顺序保存 dummy 和全部节点；移去末尾 n 个引用后，栈顶恰好是目标节点的前驱。
// Store dummy and every node in order; removing the final n references leaves the target's predecessor on top.
// 删除通过修改前驱的 Next 完成，切片截短本身不改变链表；要求 1<=n<=链表长度。
// Rewire the predecessor's Next to delete; truncating the slice alone does not change the list. Require valid n.
func removeNthFromEndStack(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	stack := []*ListNode{}
	for node := dummy; node != nil; node = node.Next {
		stack = append(stack, node)
	}

	stack = stack[:len(stack)-n]
	prev := stack[len(stack)-1]

	prev.Next = prev.Next.Next
	return dummy.Next
}
