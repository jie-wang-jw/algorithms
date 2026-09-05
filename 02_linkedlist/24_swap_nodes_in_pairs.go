package _2_linkedlist

/*
题目描述 / Problem Description
给定一个链表，两两交换相邻节点，并返回交换后的头节点。必须交换节点本身，不能只修改节点中的值。
Given a linked list, swap every two adjacent nodes and return the resulting head. Nodes themselves must be swapped rather than merely changing their values.

解题思路 / Solution Approach
使用虚拟头节点统一处理第一对节点。每轮保存一对节点及后续链表，重新连接三条指针，然后移动到下一对；不足两个节点时停止。
Use a dummy head to handle the first pair uniformly. For each pair, save both nodes and the remaining list, reconnect three links, and advance to the next pair.

时间与空间复杂度 / Time and Space Complexity
n 为节点数。时间 O(n)，每对节点进行固定次数的指针重连。辅助空间 O(1)，只增加虚拟头节点和几个临时指针。
n is the node count. Time O(n), with a constant number of link changes per pair. Auxiliary space O(1) for the dummy node and temporary pointers.
*/

/*
This is a linked-list pointer manipulation problem.
这是一个重新连接链表指针的问题。

I use a dummy node because the head may change after swapping the first pair.
因为交换第一对节点后头节点会改变，所以使用 dummy 虚拟头节点。

For each step, cur points to the node before the pair.
每一轮中，cur 都指向当前待交换节点对的前一个节点。

I take first and second, reconnect the three links to make cur point to second, second point to first, and first point to the next group.
取出 first 和 second 后，重新连接三条边：cur 指向 second，second 指向 first，first 指向下一组。

Then I move cur to first, which is now the tail of the swapped pair.
交换后 first 位于这一对的末尾，因此把 cur 移到 first，准备处理下一对。

cur 管前面
first 是第一个要交换的节点
second 是第二个要交换的节点
next 是后面还没处理的链表
最小脑内画面：
cur -> first -> second -> next
cur -> second -> first -> next
*/

func swapPairs(head *ListNode) *ListNode {
	// dummy keeps a stable node before the possibly changing head.
	// dummy 在可能变化的头节点前提供一个固定位置。
	dummy := &ListNode{Next: head}
	cur := dummy

	// A complete pair exists only when both cur.Next and cur.Next.Next exist.
	// 只有 cur.Next 和 cur.Next.Next 都不为 nil 时，才有完整的一对可以交换。
	for cur.Next != nil && cur.Next.Next != nil {
		// Remember the two nodes before reconnecting any links.
		// 修改指针前，先保存当前这一对的两个节点。
		first := cur.Next
		second := cur.Next.Next

		// 1. first points to the node after the pair. / first 先接到下一组。
		first.Next = second.Next
		// 2. second points back to first. / second 再指回 first。
		second.Next = first
		// 3. cur points to the new first node, second. / cur 最后指向交换后的头节点 second。
		cur.Next = second

		// first is now the pair's tail and becomes the predecessor of the next pair.
		// first 现在是这一对的尾节点，也就是下一对前面的节点。
		cur = first
	}

	// dummy.Next is the possibly updated head of the list.
	// dummy.Next 是交换后可能已经改变的新头节点。
	return dummy.Next
}
