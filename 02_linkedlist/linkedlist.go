package _2_linkedlist

// ListNode stores one value and a pointer to the next node.
// ListNode 保存当前节点的值，以及指向下一个节点的指针。
type ListNode struct {
	Val  int       // Value stored in this node. / 当前节点保存的值。
	Next *ListNode // Next node; nil means the end of the list. / 下一个节点；nil 表示链表结束。
}
