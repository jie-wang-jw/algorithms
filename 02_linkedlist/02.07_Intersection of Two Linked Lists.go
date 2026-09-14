package _2_linkedlist

/*
面试题 02.07. 链表相交 / Intersection of Two Linked Lists
同 160. Intersection of Two Linked Lists。

题目描述 / Problem Description
找出两个单链表相交的起始节点；不相交则返回 nil。交点比较的是指针，不是节点值。
不能破坏原链表结构。题目保证无环。
Return the node where two singly linked lists intersect, or nil if they do not.
Equality is pointer identity, not node values. Do not modify the lists. The lists are acyclic.

解法一：先对齐再同走 / Method 1: Align Tails, Then Walk Together
分别求长度。让更长的链表先走 |lenA-lenB| 步，使两个指针距离各自尾部的剩余长度相同。
再同步前进，第一次指针相等就是交点；走到末尾仍不相等则没有交点。
Count both lengths. Advance the longer list by |lenA-lenB| so both pointers have the same remaining distance to the tail.
Then walk together; the first equal pointers are the intersection, or nil if none.

解法二：拼接后同步移动 / Method 2: Switch Heads and Walk Together
pA 走完 A 后接到 B 的头，pB 走完 B 后接到 A 的头。两条路径长度都是 lenA+lenB，
因此会在交点相遇；没有交点时同时走到 nil。
After pA finishes A it continues from B's head, and pB continues from A's head.
Both walks have length lenA+lenB, so they meet at the intersection, or both become nil if there is none.

时间与空间复杂度 / Time and Space Complexity
n、m 为两条链表长度。两种解法时间 O(n+m)，辅助空间 O(1)。
For list lengths n and m, both methods take O(n+m) time and O(1) auxiliary space.
*/

// 1. Align the tails by length, then walk both pointers until they meet.
// 1. 按长度对齐尾部，再让两个指针同步前进直到相遇。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	for cur := headA; cur != nil; cur = cur.Next {
		lenA++
	}
	for cur := headB; cur != nil; cur = cur.Next {
		lenB++
	}

	// fast walks the longer list first so both pointers line up with the tails.
	// fast 先走较长的链表，使两个指针距离尾部的剩余长度相同。
	var fast, slow *ListNode
	var step int
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for range step {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 2. Switch to the other list at nil so both pointers travel n+m nodes and meet at the intersection.
// 2. 走到空时改走另一条链表，两个指针都走 n+m 个节点，并在交点相遇。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
func getIntersectionNodeTwoPointers(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
