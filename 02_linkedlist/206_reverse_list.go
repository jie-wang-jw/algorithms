package _2_linkedlist

/*
题目描述 / Problem Description
给定单链表的头节点 head，反转链表并返回反转后的新头节点。
Given the head of a singly linked list, reverse the list and return its new head.

解题思路 / Solution Approach
迭代法使用 prev 和 cur，先保存下一个节点，再令 cur.Next 指向 prev。
从后向前递归先反转后半段，再把当前节点接到反转后链表的末尾。
The iterative method uses prev and cur, saving the next node before reversing cur.Next.
The back-to-front recursion reverses the suffix first and then attaches the current node at its end.

关键逻辑：为什么这样做 / Why This Works
迭代时 prev 是已经反转部分的头，cur 是尚未处理部分的头。cur.Next=prev 把当前节点接到已反转部分前面，
所以 prev=cur；原来的后续节点必须提前保存，否则改向后就无法沿原链继续。cur=nil 表示未处理部分为空，因此 prev 就是完整结果。
递归返回后，原来的 head.Next 已成为反转后半段的尾节点。
让它的 Next 指回 head，就把 head 接到尾部；再将 head.Next=nil，切断原来的正向边，否则两节点会形成环。
Iteratively, prev heads the reversed prefix and cur heads the unprocessed suffix. Linking cur.Next to prev prepends cur,
making it the new prev. Save the old next pointer first so the remaining input stays reachable.
When cur is nil, prev heads the complete result.
After recursion, the original head.Next is the tail of the reversed suffix.
Pointing its Next to head appends head. Set head.Next=nil to remove the original forward edge; otherwise the pair forms a cycle.

时间与空间复杂度 / Time and Space Complexity
n 为节点数。reverseList：时间 O(n)，每条指针改向一次，辅助空间 O(1)。
reverseList2：时间 O(n)，递归深度为 n，调用栈占 O(n) 辅助空间。两者均复用原节点。
n is the node count. reverseList takes O(n) time and O(1) auxiliary space.
reverseList2 takes O(n) time and O(n) call-stack space due to recursion depth n.
Both reuse the original nodes.

补充解法：从前往后递归 / Additional Approach: Front-to-Back Recursion
reverseListFromFront 把迭代版的 prev 和 cur 直接作为递归参数，每层只反转一条指针，再带着新的 prev、cur 进入下一层。
cur==nil 时未处理部分为空，返回 prev 就是新头节点；它与 reverseList 使用同一套不变量，只是把循环写成了递归形式。
时间 O(n)；递归深度为 n，且 Go 不保证尾调用优化，因此调用栈占 O(n) 辅助空间，空间不如迭代版的 O(1)。
reverseListFromFront passes the iterative prev and cur as recursion parameters, reversing one link per level
and recursing with the updated pair. When cur is nil the unprocessed part is empty, so prev is the new head.
It shares reverseList's invariant and only rewrites the loop as recursion.
Time O(n); recursion depth n costs O(n) call-stack space because Go does not guarantee tail-call optimization.
*/

/*
This is a linked-list pointer reversal problem.
这是一个修改链表指针方向的问题。

I keep two pointers: prev is the already reversed part, and cur is the node I am processing.
我维护两个指针：prev 指向已经反转好的部分，cur 指向当前正在处理的节点。

For each node, I save cur.Next first, reverse cur.Next to prev, then move both pointers forward.
处理每个节点时，先保存 cur.Next，再让 cur.Next 指向 prev，最后同时推进两个指针。

At the end, prev becomes the new head.
循环结束时，prev 就是反转后链表的新头节点。

prev：已经反转好的链表头
cur：当前正在处理的节点
next：原链表里 cur 后面的节点，先保存，防止断链

next := cur.Next // 先保存后面的节点，不然会断链
cur.Next = prev  // 当前节点反向指回前一个节点
prev = cur       // prev 往前走
cur = next       // cur 往前走
*/
// 1. Iterative two-pointer solution. / 1. 迭代双指针解法。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
func reverseList(head *ListNode) *ListNode {
	// prev starts as nil because the new tail must point to nil.
	// prev 初始为 nil，因为反转后的尾节点应该指向 nil。
	var prev *ListNode
	cur := head

	for cur != nil {
		// Save the original next node before changing cur.Next.
		// 修改 cur.Next 前先保存原来的下一个节点，防止后半段链表丢失。
		next := cur.Next
		// Reverse the current link so cur points to the processed part.
		// 反转当前指针，让 cur 指向已经处理好的部分。
		cur.Next = prev
		// Move prev and cur one node forward in the original list.
		// prev 和 cur 沿原链表各向前移动一个节点。
		prev = cur
		cur = next
	}

	// cur is nil and prev points to the last original node, now the new head.
	// 此时 cur 为 nil，prev 指向原链表尾节点，也就是新头节点。
	return prev
}

// 2. Back-to-front recursion: reverse the remaining list first, then attach head at the end.
// 2. 从后向前递归：先递归反转后半段，再把 head 接到末尾。
// Time: O(n), Space: O(n) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由递归调用栈产生。
func reverseList2(head *ListNode) *ListNode {
	// An empty list or a single-node list is already reversed.
	// 空链表或只有一个节点的链表，本身就是反转结果。
	if head == nil || head.Next == nil {
		return head
	}

	// Reverse the list after head and keep its new head.
	// 先反转 head 后面的链表，并保存后半段反转后的新头节点。
	newHead := reverseList2(head.Next)

	// Before: head -> head.Next. After: head.Next -> head.
	// 原来是 head -> head.Next，现在改成 head.Next -> head。
	head.Next.Next = head
	// head becomes the new tail, so its Next must be nil.
	// head 变成新的尾节点，因此它的 Next 必须设为 nil。
	head.Next = nil

	// The head of the reversed tail is also the head of the whole result.
	// 后半段的新头节点，也是整个反转链表的新头节点。
	return newHead
}

// 3. Front-to-back recursion: the iterative prev and cur become recursion parameters.
// 3. 从前往后递归：把迭代版的 prev 和 cur 直接当作递归参数传递。
// Time: O(n), Space: O(n) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由递归调用栈产生。
func reverseListFromFront(head *ListNode) *ListNode {
	// reverse reverses exactly one link per call; prev always heads the finished part.
	// reverse 每层只反转一条指针；prev 始终是已经反转好的那部分的头节点。
	var reverse func(prev, cur *ListNode) *ListNode

	reverse = func(prev, cur *ListNode) *ListNode {
		// An empty unprocessed part means prev already heads the complete result.
		// 未处理部分为空，说明 prev 就是完整反转结果的头节点。
		if cur == nil {
			return prev
		}

		// Save the original next node before changing cur.Next.
		// 修改 cur.Next 前先保存原来的下一个节点，防止后半段链表丢失。
		next := cur.Next
		// Reverse the current link so cur points to the processed part.
		// 反转当前指针，让 cur 指向已经处理好的部分。
		cur.Next = prev

		// cur has joined the reversed part, so it becomes the next level's prev.
		// cur 已经并入已反转部分，因此它就是下一层的 prev。
		return reverse(cur, next)
	}

	// The new tail must point to nil, so the first prev is nil.
	// 反转后的尾节点必须指向 nil，因此第一层的 prev 为 nil。
	return reverse(nil, head)
}
