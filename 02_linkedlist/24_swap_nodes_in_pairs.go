package _2_linkedlist

/*
题目描述 / Problem Description
给定一个链表，两两交换相邻节点，并返回交换后的头节点。必须交换节点本身，不能只修改节点中的值。
Given a linked list, swap every two adjacent nodes and return the resulting head.
Nodes themselves must be swapped rather than merely changing their values.
*/

// 1. Dummy-head iteration: reconnect three links per pair and advance to the next pair.
// 1. 虚拟头节点迭代：每对节点重连三条指针，然后移动到下一对。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// cur 是当前二元组的前驱，first/second 是待交换节点；先接好剩余后缀，再连 second→first 和 cur→second。
// cur precedes the pair; reconnect the suffix first, then second→first and cur→second.
// 交换后 first 成为本组尾节点，也是下一组前驱；少于两个节点时停止，保留末尾单节点。
// first becomes the pair's tail and next pair's predecessor; stop when fewer than two nodes remain.
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second

		cur = first
	}

	return dummy.Next
}

// 2. Recursion: swap the suffix first, then swap the current pair in front of it.
// 2. 递归解法：先交换后面的链表，再把当前这一对接到它前面。
// Time: O(n), Space: O(n) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由递归调用栈产生。
//
// 先递归交换 second.Next 后的剩余链，将结果接到原 head，再令 second 指向 head；本组新头为 second。
// Swap the suffix after second recursively, attach it to head, then link second to head and return second.
func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	second := head.Next

	head.Next = swapPairsRecursive(second.Next)

	second.Next = head

	return second
}
