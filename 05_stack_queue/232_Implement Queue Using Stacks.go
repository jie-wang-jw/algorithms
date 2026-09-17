package _5_stack_queue

/*
232. 用栈实现队列 / Implement Queue Using Stacks

题目描述 / Problem Description
只使用栈的标准操作实现一个先进先出的队列，支持 Push、Pop、Peek 和 Empty。
Implement a first-in-first-out queue using only standard stack operations,
supporting Push, Pop, Peek, and Empty.
*/

// 1. Two stacks with lazy transfer (recommended): move inStack into outStack only when outStack is empty.
// 1. 双栈惰性转移：推荐；仅当 outStack 为空时才把 inStack 全部倒入其中。
// Time: Push amortized O(1); Pop/Peek amortized O(1), worst O(n); Empty O(1). Space: O(n).
// 时间复杂度：Push 均摊 O(1)；Pop/Peek 均摊 O(1)、单次最坏 O(n)；Empty O(1)。空间复杂度：O(n)。
// inStack handles incoming elements; outStack handles front-element access and removal.
// inStack 负责接收新加入的元素；outStack 负责读取和删除队首元素。
type MyQueue struct {
	inStack  []int
	outStack []int
}

// Create an empty queue with two empty stacks.
// 用两个空栈创建一个空队列。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

// Push x onto inStack; the front is never rearranged here.
// 把 x 压入 inStack，入队时不整理队首。
// Time: amortized O(1), worst O(n) on reallocation; backing storage O(n).
// 时间均摊 O(1)，扩容时最坏 O(n)；底层存储 O(n)。
//
// 追加到输入端，后续的取出操作负责调整顺序；append 均摊 O(1)，扩容时单次 O(n)。
// Append at the input end; removal handles reordering. append is amortized O(1), with O(n) time on reallocation.
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// Transfer every inStack element into outStack only when outStack is empty, reversing arrival order.
// 仅当 outStack 为空时，把 inStack 全部倒入 outStack，从而反转入队顺序。
// Amortized O(1), worst-case O(n) time per operation; stack storage O(n), including transfer allocations.
// 每次均摊 O(1)、最坏 O(n) 时间；两栈存储 O(n)，转移时可能分配输出栈。
//
// 仅当 outStack 为空时倒入全部 inStack；反转使最早进入的元素来到 outStack 栈顶。
// Transfer all of inStack only when outStack is empty; reversal places the oldest value on top.
// 若 outStack 未空就转移，新元素会挡住更早元素，破坏 FIFO；每项只转移一次，均摊成本为 O(1)。
// Transferring earlier would put newer values before older ones; each value transfers once, giving amortized O(1) operations.
func (q *MyQueue) moveToOut() {
	if len(q.outStack) > 0 {
		return
	}

	for len(q.inStack) > 0 {
		last := len(q.inStack) - 1

		q.outStack = append(q.outStack, q.inStack[last])

		q.inStack = q.inStack[:last]
	}
}

// Ensure the oldest element sits on outStack, then pop that top.
// 先保证队首已在 outStack 栈顶，再弹出它。
// Amortized O(1), worst-case O(n) time per operation; stack storage O(n), including transfer allocations.
// 每次均摊 O(1)、最坏 O(n) 时间；两栈存储 O(n)，转移时可能分配输出栈。
//
// 先在需要时转移，再弹出 outStack 栈顶，即最早入队的值；调用前要求队列非空。
// Transfer if needed, then remove outStack's top, the oldest value; require a nonempty queue.
func (q *MyQueue) Pop() int {
	q.moveToOut()

	last := len(q.outStack) - 1
	value := q.outStack[last]

	q.outStack = q.outStack[:last]

	return value
}

// Ensure the oldest element sits on outStack, then read that top without removing it.
// 先保证队首已在 outStack 栈顶，再只读不删。
// Amortized O(1), worst-case O(n) time per operation; stack storage O(n), including transfer allocations.
// 每次均摊 O(1)、最坏 O(n) 时间；两栈存储 O(n)，转移时可能分配输出栈。
//
// 与 Pop 共用惰性转移，但只读不删；返回最早入队值，要求队列非空。
// Use the same lazy transfer as Pop but read without removal; return the oldest value from a nonempty queue.
func (q *MyQueue) Peek() int {
	q.moveToOut()

	return q.outStack[len(q.outStack)-1]
}

// The queue is empty only when both stacks are empty.
// 只有两个栈都为空时，队列才为空。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 元素可能在任一栈中，因此必须两个栈都空才能判队列为空。
// Values may be in either stack, so the queue is empty only when both stacks are empty.
func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

// 2. Eager reordering on Push: keep the oldest element on a single stack top; this is not amortized O(1).
// 2. 入队时整理顺序：始终把最早元素放在唯一栈的栈顶；这不是均摊 O(1)。
// Time: Push O(n) every call, not amortized O(1); Pop/Peek/Empty O(1). Space: O(n).
// 时间复杂度：Push 每次都是确定的 O(n)，不是均摊 O(1)；Pop/Peek/Empty O(1)。空间复杂度：O(n)。
// stack keeps the queue inverted: top = oldest = queue front, bottom = newest = queue back.
// stack 以倒置方式保存队列：栈顶 = 最旧 = 队首，栈底 = 最新 = 队尾。
// The zero value is ready to use: a nil slice is an empty queue.
// 零值即可使用：nil 切片就是一个空队列，无需构造函数。
type MyQueueEager struct {
	stack []int
}

// Drain the main stack into a temp stack, install x as the new bottom, then restore the older elements.
// 先把旧元素全部弹入临时栈，再把 x 压成新栈底，最后倒回旧元素。
// Time: O(n), Space: O(n) for the temporary stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由临时栈产生。
//
// 目标是让最旧元素始终在栈顶；先把旧值全部移到 temp，把新值压到空栈底，再倒回旧值。
// Keep the oldest value on top: move old values to temp, place the newest at the bottom, then restore the old values.
// 每次 Push 都搬动全部旧元素，单次 O(n)，不是惰性转移的均摊 O(1)。
// Every Push moves all existing values, so it is O(n), unlike amortized lazy transfer.
func (q *MyQueueEager) Push(x int) {
	temp := make([]int, 0, len(q.stack))

	for len(q.stack) > 0 {
		last := len(q.stack) - 1
		temp = append(temp, q.stack[last])
		q.stack = q.stack[:last]
	}

	q.stack = append(q.stack, x)

	for len(temp) > 0 {
		last := len(temp) - 1
		q.stack = append(q.stack, temp[last])
		temp = temp[:last]
	}

}

// Pop the stack top, which Push already arranged to be the queue front.
// 直接弹出栈顶：Push 已经把最早入队的元素安排在栈顶。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 栈顶保持为最早入队元素，直接弹出切片末项；要求队列非空。
// The oldest value is kept at the stack top; remove the final slice entry, requiring a nonempty queue.
func (q *MyQueueEager) Pop() int {
	last := len(q.stack) - 1
	value := q.stack[last]

	q.stack = q.stack[:last]

	return value
}

// Read the stack top without removing it; that top is the queue front.
// 只读栈顶、不删除；栈顶就是队首。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 只读栈顶即最早入队值；要求队列非空。
// Read the stack top, the oldest queued value; require a nonempty queue.
func (q *MyQueueEager) Peek() int {
	return q.stack[len(q.stack)-1]
}

// Report emptiness from the single stack's length.
// 所有元素都在一个栈里，只看它的长度。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (q *MyQueueEager) Empty() bool {
	return len(q.stack) == 0
}
