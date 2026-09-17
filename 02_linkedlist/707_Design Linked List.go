package _2_linkedlist

/*
707. 设计链表 / Design Linked List

题目描述 / Problem Description
实现单链表，支持按索引取值、头插、尾插、在第 index 个节点前插入、按索引删除。
索引从 0 开始。get 在索引无效时返回 -1。addAtIndex 在 index 等于长度时尾插，大于长度时不插入，小于 0 时头插。
Implement a singly linked list supporting get, addAtHead, addAtTail, addAtIndex, and deleteAtIndex.
Indices are 0-based. get returns -1 for an invalid index. addAtIndex appends when index equals the length,
ignores index greater than the length, and inserts at the head when index is negative.
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
// Time/Space: O(1) for construction. / 构造本身时间、空间均 O(1)。
//
// 创建不存有效数据的虚拟头节点；首个有效节点始终是 dummy.Next，使头部插删也能统一操作前驱。
// Create a dummy head; the first real node is dummy.Next, so head insertion and deletion use the same predecessor logic.
func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &ListNode{}}
}

// Get the value at index, or -1 when the index is out of range.
// 返回下标 index 的值；索引无效时返回 -1。
// Time: O(index), Space: O(1).
// 时间复杂度：O(index)，空间复杂度：O(1)。
//
// 有效下标是 [0,size)；从真实头走 index 步，非法下标返回 -1。链表应由 Constructor 初始化。
// Valid indices are [0,size); walk index links from the real head or return -1. Initialize via Constructor.
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
//
// 新节点先指向原头，再让 dummy 指向新节点；保留原后缀并增加 size。
// Point the new node at the old head before attaching it to dummy, preserving the suffix and increasing size.
func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val, Next: l.dummy.Next}
	l.dummy.Next = node
	l.size++
}

// AddAtTail appends val after the last real node.
// AddAtTail 把 val 接到最后一个真实节点后面。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 从 dummy 寻找 Next==nil 的尾节点，空链也适用；接入新节点并增加 size。
// Find the node whose Next is nil starting at dummy, also covering an empty list; append and increment size.
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
// index<0 按头插处理，index==size 可尾插，index>size 不插入；从 dummy 走 index 步得到前驱。
// Negative index means head insertion, size permits append, and larger indices are ignored; index steps from dummy reach the predecessor.
// 先让新节点接住原后继，再修改前驱，避免丢链；成功插入才增加 size。
// Link the new node to the old successor before changing the predecessor; increment size only on insertion.
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
//
// 只接受 [0,size)；从 dummy 走 index 步定位前驱，跳过其后继并减少 size，删头与删中间一致。
// Accept only [0,size); find the predecessor from dummy, bypass its successor and decrement size, including head deletion.
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
