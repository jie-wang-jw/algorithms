package _2_linkedlist

/*
面试题 02.07. 链表相交 / Intersection of Two Linked Lists
同 160. Intersection of Two Linked Lists。

题目描述 / Problem Description
找出两个单链表相交的起始节点；不相交则返回 nil。交点比较的是指针，不是节点值。
不能破坏原链表结构。题目保证无环。
Return the node where two singly linked lists intersect, or nil if they do not.
Equality is pointer identity, not node values. Do not modify the lists. The lists are acyclic.
*/

// 1. Align tails by length, then walk together
// 1. 按长度对齐尾部，再同步前进
// Advance the longer list by |lenA-lenB|, then walk both until pointers meet (intersection or both nil).
// 让较长链表先走 |lenA-lenB|，再同步前进直到指针相等（交点或同为 nil）。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
//
// 先让长链走长度差，使两指针到尾部的剩余长度相等，再同步前进；相交时同刻到交点，否则同刻到 nil。
// Advance the longer list by the length difference; equal remaining distances make pointers meet at the intersection or nil.
// 比较节点地址而非 Val；前提是两条链表均无环。
// Compare node addresses, not values; both lists must be acyclic.
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	for cur := headA; cur != nil; cur = cur.Next {
		lenA++
	}
	for cur := headB; cur != nil; cur = cur.Next {
		lenB++
	}

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

// 2. Switch to the other list at nil
// 2. 走到空时改走另一条链表
// Both pointers travel n+m nodes and meet at the intersection, or both become nil if none.
// 两个指针都走 n+m 个节点，在交点相遇；没有交点时同时变成 nil。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
//
// pA 走 A 后转 B，pB 走 B 后转 A，交换链头抵消两条非公共前缀的长度差。
// pA follows A then B while pB follows B then A; switching heads cancels the unequal prefix lengths.
// 相交就以节点地址相等结束；不相交则都走完两链后在 nil 相遇。前提是无环链表。
// Equal addresses identify the intersection; disjoint acyclic lists end with both pointers at nil.
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
