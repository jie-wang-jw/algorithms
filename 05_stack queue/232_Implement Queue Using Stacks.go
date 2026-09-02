package _5_stack_queue

/*
Use two stacks to simulate a FIFO queue.
使用两个栈模拟先进先出的队列。

inStack stores newly added elements.
inStack 保存新加入的元素。

outStack stores elements in queue-removal order.
outStack 按照队列出队顺序保存元素。

When outStack is empty, move every element from inStack to outStack.
当 outStack 为空时，将 inStack 中的所有元素转移到 outStack。

The transfer reverses the order, so the oldest element becomes the top
of outStack and can be removed first.
转移会反转元素顺序，因此最早加入的元素会来到 outStack 栈顶，
从而最先被删除。
*/

type MyQueue struct {
	// inStack handles incoming elements.
	// inStack 负责接收新加入的元素。
	inStack []int

	// outStack handles front-element access and removal.
	// outStack 负责读取和删除队首元素。
	outStack []int
}

// Constructor creates an empty queue.
// Constructor 创建一个空队列。
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

// Push adds x to the back of the queue.
// Push 将 x 加入队尾。
func (q *MyQueue) Push(x int) {
	// New elements always enter inStack first.
	// 新元素始终先进入 inStack。
	q.inStack = append(q.inStack, x)
}

// moveToOut moves elements into outStack when necessary.
// moveToOut 在需要时将元素转移到 outStack。
func (q *MyQueue) moveToOut() {
	// Keep the existing order if outStack still has elements.
	// 如果 outStack 中还有元素，保持当前顺序，不进行转移。
	if len(q.outStack) > 0 {
		return
	}

	// Move every element from inStack to outStack.
	// 将 inStack 中的所有元素依次转移到 outStack。
	for len(q.inStack) > 0 {
		// Find the index of the top element in inStack.
		// 找到 inStack 栈顶元素的下标。
		last := len(q.inStack) - 1

		// Push the top element of inStack into outStack.
		// 将 inStack 的栈顶元素压入 outStack。
		q.outStack = append(q.outStack, q.inStack[last])

		// Remove the top element from inStack.
		// 删除 inStack 的栈顶元素。
		q.inStack = q.inStack[:last]
	}
}

// Pop removes and returns the front element.
// Pop 删除并返回队首元素。
func (q *MyQueue) Pop() int {
	// Make sure the front element is on top of outStack.
	// 确保队首元素位于 outStack 的栈顶。
	q.moveToOut()

	// The top of outStack is the front of the queue.
	// outStack 的栈顶就是队首元素。
	last := len(q.outStack) - 1
	value := q.outStack[last]

	// Remove the top element from outStack.
	// 删除 outStack 的栈顶元素。
	q.outStack = q.outStack[:last]

	return value
}

// Peek returns the front element without removing it.
// Peek 返回队首元素，但不删除它。
func (q *MyQueue) Peek() int {
	// Make sure the front element is on top of outStack.
	// 确保队首元素位于 outStack 的栈顶。
	q.moveToOut()

	// Read the top element without modifying the stack.
	// 读取栈顶元素，但不修改栈。
	return q.outStack[len(q.outStack)-1]
}

// Empty returns whether the queue contains no elements.
// Empty 判断队列中是否没有元素。
func (q *MyQueue) Empty() bool {
	// The queue is empty only when both stacks are empty.
	// 只有两个栈都为空时，队列才为空。
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
