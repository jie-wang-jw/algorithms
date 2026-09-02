package _5_stack_queue

/*
Use one queue to simulate a LIFO stack.
使用一个队列模拟后进先出的栈。

The front of the queue always represents the top of the stack.
队首元素始终表示栈顶元素。

After adding a new element, rotate all older elements to the back.
加入新元素后，将所有旧元素依次移动到队尾。

This moves the new element to the front, making it the new stack top.
这样新元素就会移动到队首，成为新的栈顶。
*/

type MyStack struct {
	// queue stores elements in stack order:
	// front = stack top, back = stack bottom.
	// queue 按照栈的顺序保存元素：
	// 队首 = 栈顶，队尾 = 栈底。
	queue []int
}

// StackConstructor creates an empty stack.
// StackConstructor 创建一个空栈。
func StackConstructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

// Push adds x onto the top of the stack.
// Push 将 x 压入栈顶。
func (s *MyStack) Push(x int) {
	// Save the number of existing elements.
	// 保存加入新元素之前的旧元素数量。
	size := len(s.queue)

	// Add the new element to the back of the queue.
	// 将新元素加入队尾。
	s.queue = append(s.queue, x)

	// Move all older elements from the front to the back.
	// 将所有旧元素从队首依次移动到队尾。
	for i := 0; i < size; i++ {
		// Read the front element.
		// 读取队首元素。
		front := s.queue[0]

		// Remove the front element.
		// 删除队首元素。
		s.queue = s.queue[1:]

		// Add it back to the end of the queue.
		// 将该元素重新加入队尾。
		s.queue = append(s.queue, front)
	}

	// The new element is now at the front.
	// 此时，新元素已经位于队首。
}

// Pop removes and returns the top element.
// Pop 删除并返回栈顶元素。
func (s *MyStack) Pop() int {
	// The queue front represents the stack top.
	// 队首元素就是栈顶元素。
	top := s.queue[0]

	// Remove the front element from the queue.
	// 从队列中删除队首元素。
	s.queue = s.queue[1:]

	return top
}

// Top returns the top element without removing it.
// Top 返回栈顶元素，但不删除它。
func (s *MyStack) Top() int {
	// The queue front always represents the stack top.
	// 队首元素始终表示栈顶元素。
	return s.queue[0]
}

// Empty returns whether the stack contains no elements.
// Empty 判断栈中是否没有元素。
func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}
