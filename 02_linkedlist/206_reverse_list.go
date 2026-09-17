package _2_linkedlist

/*
题目描述 / Problem Description
给定单链表的头节点 head，反转链表并返回反转后的新头节点。
Given the head of a singly linked list, reverse the list and return its new head.
*/

// 1. Iterative two pointers: prev heads the reversed part, cur heads the rest, flipping one link per step.
// 1. 迭代双指针：prev 是已反转部分的头，cur 是未处理部分的头，每轮反转一条指针。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// prev 是已反转部分的头，cur 是未处理部分的头；先存 next，再反转 cur.Next，最后同时推进两个部分。
// prev heads the reversed part and cur the untouched part; save next before reversing the link, then advance both fronts.
// 不先保存 next 就会丢失原后缀；cur==nil 时 prev 是整条反转链的头。
// Saving next preserves the original suffix; when cur becomes nil, prev heads the complete reversed list.
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

// 2. Back-to-front recursion: reverse the remaining list first, then attach head at the end.
// 2. 从后向前递归：先递归反转后半段，再把 head 接到末尾。
// Time: O(n), Space: O(n) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由递归调用栈产生。
//
// 先反转 head.Next 之后的链；原后继现在是该段尾节点，让它指回 head，再把 head.Next 清空以断开旧边。
// Reverse the suffix first; the original successor becomes its tail. Link it back to head, then clear head.Next to avoid a cycle.
// newHead 从最深层原尾节点一路返回；空链或单节点无需反转。
// newHead propagates the original tail from the deepest call; empty and single-node lists are base cases.
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}

// 3. Front-to-back recursion: the iterative prev and cur become recursion parameters.
// 3. 从前往后递归：把迭代版的 prev 和 cur 直接当作递归参数传递。
// Time: O(n), Space: O(n) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由递归调用栈产生。
//
// 递归参数 prev/cur 分别是已反转和未处理部分；保存后继、反转当前边，再递归处理下一节点。
// prev/cur separate reversed and untouched parts; save the successor, reverse the current link, then recurse.
// cur==nil 返回 prev；Go 不保证尾调用消除，因此递归仍占 O(n) 栈空间。
// Return prev when cur is nil; Go does not guarantee tail-call elimination, so stack space is O(n).
func reverseListFromFront(head *ListNode) *ListNode {
	var reverse func(prev, cur *ListNode) *ListNode

	reverse = func(prev, cur *ListNode) *ListNode {
		if cur == nil {
			return prev
		}

		next := cur.Next
		cur.Next = prev

		return reverse(cur, next)
	}

	return reverse(nil, head)
}
