package _2_linkedlist

/*
题目描述 / Problem Description
给定单链表的头节点 head，反转链表并返回反转后的新头节点。
Given the head of a singly linked list, reverse the list and return its new head.

解题思路 / Solution Approach
迭代法使用 prev 和 cur，先保存下一个节点，再令 cur.Next 指向 prev。递归法先反转后半段，再把当前节点接到反转后链表的末尾。
The iterative method uses prev and cur, saving the next node before reversing cur.Next. The recursive method reverses the suffix first and then attaches the current node at its end.
*/

/*
This is a linked-list pointer reversal problem.
这是一个修改链表指针方向的问题。

I keep two pointers: prev is the already reversed part, and cur is the node I am processing.
我维护两个指针：prev 指向已经反转好的部分，cur 指向当前正在处理的节点。

For each node, I save cur.Next first, reverse cur.Next to prev, then move both pointers forward.
处理每个节点时，先保存 cur.Next，再让 cur.Next 指向 prev，最后同时推进两个指针。

At the end, prev becomes the new head.
循环结束时，prev 就是反转后链表的新头节点。

prev：已经反转好的链表头
cur：当前正在处理的节点
next：原链表里 cur 后面的节点，先保存，防止断链

next := cur.Next // 先保存后面的节点，不然会断链
cur.Next = prev  // 当前节点反向指回前一个节点
prev = cur       // prev 往前走
cur = next       // cur 往前走
*/
// Iterative two-pointer solution. / 迭代双指针解法。
func reverseList(head *ListNode) *ListNode {
	// prev starts as nil because the new tail must point to nil.
	// prev 初始为 nil，因为反转后的尾节点应该指向 nil。
	var prev *ListNode
	cur := head

	for cur != nil {
		// Save the original next node before changing cur.Next.
		// 修改 cur.Next 前先保存原来的下一个节点，防止后半段链表丢失。
		next := cur.Next
		// Reverse the current link so cur points to the processed part.
		// 反转当前指针，让 cur 指向已经处理好的部分。
		cur.Next = prev
		// Move prev and cur one node forward in the original list.
		// prev 和 cur 沿原链表各向前移动一个节点。
		prev = cur
		cur = next
	}

	// cur is nil and prev points to the last original node, now the new head.
	// 此时 cur 为 nil，prev 指向原链表尾节点，也就是新头节点。
	return prev
}

// Recursive solution: reverse the remaining list first, then attach head at the end.
// 递归解法：先递归反转后半段，再把 head 接到末尾。
func reverseList2(head *ListNode) *ListNode {
	// An empty list or a single-node list is already reversed.
	// 空链表或只有一个节点的链表，本身就是反转结果。
	if head == nil || head.Next == nil {
		return head
	}

	// Reverse the list after head and keep its new head.
	// 先反转 head 后面的链表，并保存后半段反转后的新头节点。
	newHead := reverseList2(head.Next)

	// Before: head -> head.Next. After: head.Next -> head.
	// 原来是 head -> head.Next，现在改成 head.Next -> head。
	head.Next.Next = head
	// head becomes the new tail, so its Next must be nil.
	// head 变成新的尾节点，因此它的 Next 必须设为 nil。
	head.Next = nil

	// The head of the reversed tail is also the head of the whole result.
	// 后半段的新头节点，也是整个反转链表的新头节点。
	return newHead
}
