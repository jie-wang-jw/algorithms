package _2_linkedlist

/*
203. 移除链表元素 / Remove Linked List Elements

题目描述 / Problem Description
删除链表中所有节点值等于 val 的节点，并返回新的头节点。
Remove every node whose value equals val and return the new head.

示例 / Example
head = [1,2,6,3,4,5,6], val = 6 → [1,2,3,4,5]
*/

// 1. Delete on the original list
// 1. 直接在原链表上删除
// Strip matching heads first; then from predecessor cur, skip matching successors and stay put after a skip.
// 先摘掉连续匹配的头；再由前驱 cur 跳过匹配后继，跳过后停在 cur。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 先连续移除匹配的头节点；随后 cur 始终指向保留节点，检查的是 cur.Next。
// Remove matching heads first; thereafter cur is retained and the candidate is cur.Next.
// 删除后 cur 不移动，因为新接上的后继也可能要删；保留后继时才前进。
// After deletion keep cur fixed to check its new successor; advance only when retaining that successor.
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
//
// dummy 为真实头也提供前驱；cur.Next 是待检查节点。删除后 cur 不动，确保连续匹配项都被删除。
// dummy supplies a predecessor even for the head; inspect cur.Next and keep cur fixed after deletion to catch consecutive matches.
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
//
// 递归返回“删除完成后的子链表头”。当前值要删除就直接返回后缀结果，否则接回处理后的后缀再返回 head。
// Each call returns the filtered sublist head; discard the current node or reconnect the filtered suffix and return head.
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
