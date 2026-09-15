package _2_linkedlist

/*
203. 移除链表元素 / Remove Linked List Elements

题目描述 / Problem Description
删除链表中所有节点值等于 val 的节点，并返回新的头节点。
Remove every node whose value equals val and return the new head.

示例 / Example
head = [1,2,6,3,4,5,6], val = 6 → [1,2,3,4,5]

解法一：直接在原链表上删除 / Method 1: Delete on the Original List
头节点没有前驱，需要先用 while 把连续等于 val 的头节点摘掉。
其余节点由前驱 cur 检查 cur.Next：相等则跳过，否则 cur 前进。
The head has no predecessor, so first drop every leading node equal to val.
For the rest, predecessor cur inspects cur.Next: skip equals, otherwise advance cur.

解法二：虚拟头节点（推荐） / Method 2: Dummy Head (Recommended)
给原头补一个 dummy，所有待删节点都有前驱，删除逻辑统一。返回 dummy.Next。
A dummy in front of the original head gives every target a predecessor, so deletion is uniform. Return dummy.Next.

解法三：递归 / Method 3: Recursion
空链表返回 nil。头等于 val 则答案是对后续链表递归的结果；否则把头接到递归后的后续链表上。
An empty list returns nil. If the head equals val, the answer is the recursive result on the suffix;
otherwise attach the head to that suffix.

时间与空间复杂度 / Time and Space Complexity
n 为节点数。三种解法时间都是 O(n)。解法一、二辅助空间 O(1)；解法三递归栈 O(n)。
n is the node count. All three take O(n) time. Methods 1 and 2 use O(1) auxiliary space; method 3 uses O(n) call-stack space.
*/

// 1. Delete on the original list
// 1. 直接在原链表上删除
// Strip matching heads first; then from predecessor cur, skip matching successors and stay put after a skip.
// 先摘掉连续匹配的头；再由前驱 cur 跳过匹配后继，跳过后停在 cur。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 步骤与要点 / Steps and notes:
//  1. Stay at cur: the successor that slid in may also equal val.
//     停在 cur：补上来的后继仍可能等于 val。
func removeElementsDirect(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 2. Dummy head (recommended)
// 2. 虚拟头节点（推荐）
// Dummy gives every real node a predecessor so deletion uses one loop; stay at cur after a skip.
// 虚拟头让每个真实节点都有前驱，删除只需一套逻辑；跳过后停在 cur。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 3. Recursion
// 3. 递归
// Drop a matching head; otherwise keep it and reconnect to the cleaned suffix.
// 头等于 val 就丢掉；否则保留头并接到清理后的后缀上。
// Time: O(n), Space: O(n) for the call stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由递归栈产生。
func removeElementsRecursive(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return removeElementsRecursive(head.Next, val)
	}
	head.Next = removeElementsRecursive(head.Next, val)
	return head
}
