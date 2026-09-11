package _2_linkedlist

/*
题目描述 / Problem Description
给定一个链表的头节点 head，如果链表中存在环，返回环的入口节点；如果不存在环，则返回 nil。不能修改链表。
Given the head of a linked list, return the node where a cycle begins, or nil if the list has no cycle.
The list must not be modified.

解题思路 / Solution Approach
使用 Floyd 快慢指针。slow 每次走一步，fast 每次走两步；如果相遇则存在环。
随后一个指针回到 head，两个指针同速前进，再次相遇的位置就是环入口。
Use Floyd's slow and fast pointers. If they meet, a cycle exists. Move one pointer back to head,
advance both one step at a time, and their next meeting point is the cycle entrance.

关键逻辑：为什么这样做 / Why This Works
为什么回到 head 后会在入口相遇？设 head 到入口距离为 a，入口沿环到相遇点距离为 b，环长为 c。
相遇时 slow 走了 a+b+t*c 步，fast 是它的两倍；两者距离差又是整圈 k*c，因此 a+b+t*c=k*c，即 a+b 是 c 的倍数。
于是相遇点再走 a 步，其环内位置是 (b+a) mod c=0，恰好回到入口；head 出发的指针走 a 步也恰到入口。
在此之前，head 指针还在环外，不可能与环内指针相遇。如果 a=0，两个指针设置好时已经在入口相等，不必再走。
Let a be the distance from head to the entrance, b the forward distance from entrance to meeting point,
and c the cycle length. At the meeting, slow traveled a+b+t*c steps. Fast traveled twice as far,
and their distance difference is k*c, so a+b is a multiple of c.
Walking a more steps from the meeting point gives cycle position (b+a) mod c=0: the entrance.
The pointer from head also reaches it after a steps. Before then it is outside the cycle and cannot meet the other pointer.
If a=0, both are already equal at the entrance when reset.

时间与空间复杂度 / Time and Space Complexity
n 为不同节点数。时间 O(n)：快慢指针找相遇点、再找入口都只需线性步数。辅助空间 O(1)，仅保存几个节点指针。
n is the number of distinct nodes. Time O(n): both finding the meeting point and locating the entrance take linear steps.
Auxiliary space O(1) for a few node pointers.

补充解法：节点集合 / Alternative: Visited Node Set
detectCycleHash 保存走过的节点指针。第一次再次遇到的节点就是环入口：
入环前节点不会重复，入环后按 Next 绕一圈首先回到入口。比较地址而非 Val。
平均时间 O(n)，辅助空间 O(n)；快慢指针省掉了集合的空间。
Store visited pointers, not values. Prefix nodes never repeat; the first complete lap returns to the entry.
Expected time O(n), auxiliary space O(n); Floyd's existing method avoids the set.
*/

/*
141: Check whether a cycle exists. / 判断有没有环。
142: If a cycle exists, find its entrance. / 如果有环，找到入环点。

Phase 1: fast gains one node on slow each round, so they must meet inside a cycle.
第一阶段：fast 每轮比 slow 多走一步，所以有环时一定会追上 slow。

Phase 2: p1 and p2 move at the same speed and meet at the cycle entrance.
第二阶段：p1 和 p2 同速前进，并在入环点相遇。

After slow and fast meet inside the cycle, one pointer starts from head and the other starts from the meeting point.
If both move one step at a time, they meet at the cycle entrance.

This is Floyd's cycle detection algorithm.
First, I use slow and fast pointers to check whether a cycle exists.
Slow moves one step and fast moves two steps.
If they meet, there is a cycle.
Then I put one pointer at the head and keep the other at the meeting point.
I move both one step at a time, and their next meeting point is the cycle entrance.
If fast reaches nil, there is no cycle.
*/

// 1. Floyd's slow and fast pointers: find the meeting point, then walk to the entrance.
// 1. Floyd 快慢指针：先找到相遇点，再从头和相遇点同速走到入环点。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
func detectCycle(head *ListNode) *ListNode {
	// Both pointers start at head; slow moves 1 step and fast moves 2 steps.
	// 两个指针都从 head 出发；slow 每次走 1 步，fast 每次走 2 步。
	slow, fast := head, head

	// fast and fast.Next must both exist before fast can move two steps.
	// fast 连走两步前，必须确保 fast 和 fast.Next 都不为 nil。
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// Pointer equality means both variables point to the exact same node.
		// 指针相等表示 slow 和 fast 指向内存中的同一个节点，而不只是节点值相同。
		if slow == fast {
			// slow is the first meeting point inside the cycle, not necessarily the entrance.
			// slow 是快慢指针在环内的第一次相遇点，但它不一定是入环点。
			p1, p2 := head, slow

			/*
				第一次相遇：确认“有环”
				First meeting: proves there is a cycle.
				第二次相遇：定位“入口”
				Second meeting: finds the cycle entrance.
			*/
			for p1 != p2 {
				// After a steps, p1 reaches the entrance and p2 reaches cycle offset (b+a) mod c=0.
				// 走 a 步后 p1 从头到入口，p2 的环内位置为 (b+a) mod c=0，也恰好到入口。
				p1 = p1.Next
				p2 = p2.Next
			}

			/*
				Return p1 because after the loop both p1 and p2 point to the cycle entrance.
				返回 p1，因为循环结束后 p1 和 p2 都已经指向入环点；此时 p1 不一定还是原来的 head。
			*/
			return p1
		}
	}

	// Reaching nil means the list ends, so no cycle exists.
	// fast 能走到 nil，说明链表存在终点，因此没有环。
	return nil
}

// 2. Visited node set: the first node seen twice is the cycle entrance.
// 2. 节点集合法：第一个被重复访问到的节点就是入环点。
// Time: O(n) expected, Space: O(n) for the set.
// 时间复杂度：期望 O(n)，空间复杂度：O(n)，由集合产生。
func detectCycleHash(head *ListNode) *ListNode {
	// The key is the node pointer, not its value, because values may repeat.
	// 键是节点指针而不是节点值，因为不同节点的值可能相同。
	seen := make(map[*ListNode]bool)

	for node := head; node != nil; node = node.Next {
		// Nodes before the cycle are never revisited, so the first repeat is the entrance.
		// 入环前的节点不会被重复访问，因此第一个重复出现的节点就是入环点。
		if seen[node] {
			return node
		}

		seen[node] = true
	}

	// Reaching nil means the list ends, so no cycle exists.
	// 能走到 nil 说明链表存在终点，因此没有环。
	return nil
}
