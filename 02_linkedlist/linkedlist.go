package _2_linkedlist

// ListNode stores one value and a pointer to the next node.
// ListNode 保存当前节点的值，以及指向下一个节点的指针。
// Val is the stored value; Next points to the next node, or nil at the end.
// Val 保存节点值；Next 指向下一个节点，nil 表示链表结束。
type ListNode struct {
	Val  int
	Next *ListNode
}
