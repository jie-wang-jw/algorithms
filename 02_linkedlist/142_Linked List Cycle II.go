package _2_linkedlist

/*
141: 判断有没有环
142: 如果有环，找到入环点

第一阶段：fast 每次比 slow 多走一步，所以有环一定会追上。
第二阶段：p1 和 p2 每次同速前进，所以会在入环点相遇。

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

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			//slow = 快慢指针在环里的相遇点
			//slow = the meeting point inside the cycle
			p1, p2 := head, slow

			/*
				第一次相遇：确认“有环”
				First meeting: proves there is a cycle.
				第二次相遇：定位“入口”
				Second meeting: finds the cycle entrance.
			*/
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}

			/*We return p1 because after the loop, p1 has moved from head to the cycle entrance.
			It is no longer necessarily the original head.*/
			return p1
		}
	}

	return nil
}
