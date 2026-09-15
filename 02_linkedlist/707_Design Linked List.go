package _2_linkedlist

/*
707. 设计链表 / Design Linked List

题目描述 / Problem Description
实现单链表，支持按索引取值、头插、尾插、在第 index 个节点前插入、按索引删除。
索引从 0 开始。get 在索引无效时返回 -1。addAtIndex 在 index 等于长度时尾插，大于长度时不插入，小于 0 时头插。
Implement a singly linked list supporting get, addAtHead, addAtTail, addAtIndex, and deleteAtIndex.
Indices are 0-based. get returns -1 for an invalid index. addAtIndex appends when index equals the length,
ignores index greater than the length, and inserts at the head when index is negative.

解法：虚拟头节点 / Dummy Head
用 dummy 指向真实头节点，并维护 size。get 从 dummy.Next 走 index 步。
插入和删除都先走到待操作位置的前驱，前驱对头、尾、中间都存在，因此不需要单独处理头节点。
dummy points at the real head, and size tracks the number of real nodes. get walks index steps from dummy.Next.
Insert and delete first move to the predecessor, which exists for head, tail, and middle, so the head is not special.

时间与空间复杂度 / Time and Space Complexity
涉及 index 的操作为 O(index)，addAtHead 为 O(1)，addAtTail 为 O(n)。
空间 O(n)，保存全部节点和虚拟头。
Index operations take O(index) time, addAtHead O(1), addAtTail O(n).
Space is O(n) for the nodes plus the dummy.
*/

// MyLinkedList is a singly linked list with a dummy head, as Carl's article designs it.
// MyLinkedList 是带虚拟头节点的单链表，对应卡哥文章的写法。
// dummy.Next is the real head; size counts real nodes only.
// dummy.Next 才是真实头；size 只统计真实节点。
type MyLinkedList struct {
	dummy *ListNode
	size  int
}

// Constructor: dummy-head singly linked list
// 构造：虚拟头节点单链表
// Predecessors exist for every real node, so head/tail/middle share one insert/delete pattern.
// 每个真实节点都有前驱，头/尾/中间共用同一套插入删除逻辑。
// Time: get/addAtIndex/deleteAtIndex O(index), addAtHead O(1), addAtTail O(n); Space: O(n).
// 时间复杂度：get/addAtIndex/deleteAtIndex 为 O(index)，addAtHead 为 O(1)，addAtTail 为 O(n)；空间复杂度：O(n)。
func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &ListNode{}}
}

// Get the value at index, or -1 when the index is out of range.
// 返回下标 index 的值；索引无效时返回 -1。
// Time: O(index), Space: O(1).
// 时间复杂度：O(index)，空间复杂度：O(1)。
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.dummy.Next
	for range index {
		cur = cur.Next
	}
	return cur.Val
}

// AddAtHead inserts val before the current first node.
// AddAtHead 把 val 插到当前第一个节点之前。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val, Next: l.dummy.Next}
	l.dummy.Next = node
	l.size++
}

// AddAtTail appends val after the last real node.
// AddAtTail 把 val 接到最后一个真实节点后面。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
func (l *MyLinkedList) AddAtTail(val int) {
	cur := l.dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	l.size++
}

// AddAtIndex inserts val before index, or at the tail when index equals size.
// AddAtIndex 在下标 index 前插入 val；index 等于长度时尾插。
// Time: O(index), Space: O(1).
// 时间复杂度：O(index)，空间复杂度：O(1)。
//
// 步骤与要点 / Steps and notes:
//  1. Negative index is treated as insertion at the head.
//     负数下标按头插处理。
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}
	if index < 0 {
		index = 0
	}
	cur := l.dummy
	for range index {
		cur = cur.Next
	}
	node := &ListNode{Val: val, Next: cur.Next}
	cur.Next = node
	l.size++
}

// DeleteAtIndex removes the node at index when the index is valid.
// DeleteAtIndex 在索引有效时删除下标 index 的节点。
// Time: O(index), Space: O(1).
// 时间复杂度：O(index)，空间复杂度：O(1)。
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	cur := l.dummy
	for range index {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.size--
}
