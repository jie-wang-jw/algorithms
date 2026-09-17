package _2_linkedlist

/*
题目描述 / Problem Description
给定一个链表的头节点 head，如果链表中存在环，返回环的入口节点；如果不存在环，则返回 nil。不能修改链表。
Given the head of a linked list, return the node where a cycle begins, or nil if the list has no cycle.
The list must not be modified.
*/

// 1. Floyd's slow and fast pointers: find the meeting point, then walk to the entrance.
// 1. Floyd 快慢指针：先找到相遇点，再从头和相遇点同速走到入环点。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 快指针每轮比慢指针多走一步，有环时必相遇；访问 fast.Next.Next 前先检查 fast 和 fast.Next。
// The fast pointer gains one step per round and must catch the slow one in a cycle; guard both links before advancing.
// 设入口前长度 a、入口到相遇点距离 b、环长 c，慢指针相遇前走 t=a+b+k*c，快指针多走 t，故 t 是 c 的倍数。
// Let a be the entry distance, b the entry-to-meeting distance and c the cycle length; slow travels t=a+b+k*c and fast gains t, a multiple of c.
// 因此 a+b≡0(mod c)：从头和相遇点各走 a 步都会到入口；两者改为每次一步，首次相遇即入口。
// Thus a+b is divisible by c; moving one pointer from the head and one from the meeting point at equal speed finds the entry.
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p1, p2 := head, slow

			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}

			return p1
		}
	}

	return nil
}

// 2. Visited node set: the first node seen twice is the cycle entrance.
// 2. 节点集合法：第一个被重复访问到的节点就是入环点。
// Time: O(n) expected, Space: O(n) for the set.
// 时间复杂度：期望 O(n)，空间复杂度：O(n)，由集合产生。
//
// 按节点地址记录访问；第一次再次遇见的节点就是环入口，值相同并不代表同一节点。
// Track node addresses; the first repeated node is the cycle entry, whereas equal values do not imply identity.
func detectCycleHash(head *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)

	for node := head; node != nil; node = node.Next {
		if seen[node] {
			return node
		}

		seen[node] = true
	}

	return nil
}
