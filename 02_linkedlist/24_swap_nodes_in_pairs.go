package _2_linkedlist

/*This is a linked-list pointer manipulation problem.
I use a dummy node because the head may change after swapping the first pair.
For each step, cur points to the node before the pair.
I take first and second, reconnect the three links to make cur point to second, second point to first, and first point to the next group.
Then I move cur to first, which is now the tail of the swapped pair.
cur 管前面
first 是第一个要交换的节点
second 是第二个要交换的节点
next 是后面还没处理的链表
最小脑内画面：
cur -> first -> second -> next
cur -> second -> first -> next
*/

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second

		cur = first
	}

	return dummy.Next
}
