package _5_stack_queue

/*
225. 用队列实现栈 / Implement Stack Using Queues

题目描述 / Problem Description
只使用队列的标准操作实现一个后进先出的栈，支持 Push、Pop、Top 和 Empty。
Implement a last-in-first-out stack using only standard queue
operations, supporting Push, Pop, Top, and Empty.
*/

// MyStack 1. Rotate on Push (recommended): one queue held in stack order, so its front is always the stack top.
// 1. 入栈时旋转：推荐；单个队列始终按栈序保存元素，队首永远是栈顶。
// Time: Push O(n), Pop/Top/Empty O(1). Space: O(n).
// 时间复杂度：Push O(n)，Pop/Top/Empty O(1)。空间复杂度：O(n)。
// queue stores elements in stack order: front = stack top, back = stack bottom.
// queue 按照栈的顺序保存元素：队首 = 栈顶，队尾 = 栈底。
type MyStack struct {
	queue []int
}

// StackConstructor Create an empty stack backed by one queue already stored in stack order.
// 创建一个空栈，底层单个队列已按栈序准备好。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func StackConstructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

// Push Append x, then rotate every older element behind it so the front becomes the new top.
// 先把 x 追加到队尾，再把所有旧元素依次转到后面，使队首成为新的栈顶。
// Time: O(n); storage O(n) for slice-backed queues, including possible append reallocations.
// 时间 O(n)；切片队列及可能的扩容需要 O(n) 存储，不能把底层队列空间也算成 O(1)。
//
// 入队后把原 size 个元素依次旋转到队尾，新值就来到队首；此后队首始终是栈顶。
// After appending, rotate the original size elements to the back; the new value becomes the front and thus the stack top.
func (s *MyStack) Push(x int) {
	size := len(s.queue)

	s.queue = append(s.queue, x)

	for range size {
		front := s.queue[0]

		s.queue = s.queue[1:]

		s.queue = append(s.queue, front)
	}

}

// Pop the queue front, which each Push rotation already made the stack top.
// 直接弹出队首：入栈时的旋转已经保证队首就是栈顶。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// Push 已将栈顶放在队首，移除 queue[0] 即可；调用前要求栈非空。
// Push maintains the top at queue[0]; remove that entry, requiring a nonempty stack.
func (s *MyStack) Pop() int {
	top := s.queue[0]

	s.queue = s.queue[1:]

	return top
}

// Read the queue front without removing it; that front is the stack top.
// 只读队首、不删除；队首就是栈顶。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 只读取队首，不删除；调用前要求栈非空。
// Read the front without removal; require a nonempty stack.
func (s *MyStack) Top() int {
	return s.queue[0]
}

// Report emptiness from the queue length alone.
// 只看队列长度判断栈是否为空。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

// 2. Rotate on Pop: keep arrival order and defer the O(n) rotation until removal.
// 2. 出栈时旋转：队列按入栈先后保存，把 O(n) 成本推迟到 Pop。
// Time: Push amortized O(1), Pop/Top O(n), Empty O(1). Space: O(n).
// 时间复杂度：Push 均摊 O(1)，Pop/Top O(n)，Empty O(1)。空间复杂度：O(n)。
// queue stores elements in arrival order: front = oldest = stack bottom, back = newest = stack top.
// queue 按照入栈先后保存元素：队首 = 最旧 = 栈底，队尾 = 最新 = 栈顶。
// The zero value is ready to use: a nil slice is an empty stack.
// 零值即可使用：nil 切片就是一个空栈，无需构造函数。
type MyStackLazy struct {
	queue []int
}

// Append x at the back; the back is the stack top, so Push does not rotate.
// 直接追加到队尾：队尾就是栈顶，入栈不必旋转。
// Time: amortized O(1), worst O(n) on reallocation; backing storage O(n).
// 时间均摊 O(1)，扩容时最坏 O(n)；底层存储 O(n)。
//
// 队尾保存最新入栈元素，直接追加即可；重新排列的成本留给 Pop/Top。
// The newest value is at the back; append now and defer reordering to Pop/Top.
func (s *MyStackLazy) Push(x int) {
	s.queue = append(s.queue, x)
}

// Rotate the oldest n-1 elements to the back, then dequeue the newest value at the front.
// 把前 n-1 个旧元素转到队尾，再从队首弹出最新入栈的值。
// Time: O(n); storage O(n) for slice-backed queues, including possible append reallocations.
// 时间 O(n)；切片队列及可能的扩容需要 O(n) 存储，不能把底层队列空间也算成 O(1)。
//
// 将前 n-1 个旧元素旋转到队尾，使最新元素来到队首后删除；其余元素仍保持原入队顺序。
// Rotate the first n-1 values so the newest reaches the front, then remove it; remaining values keep arrival order.
// 要求栈非空；队列首尾分别对应最旧和最新元素。
// Require a nonempty stack; queue front/back represent oldest/newest values.
func (s *MyStackLazy) Pop() int {
	for range len(s.queue) - 1 {
		front := s.queue[0]

		s.queue = s.queue[1:]

		s.queue = append(s.queue, front)
	}

	value := s.queue[0]
	s.queue = s.queue[1:]

	return value
}

// Reuse Pop then Push so the newest value is read without accessing the queue back.
// 复用 Pop 取出栈顶，再用 Push 加回队尾，避免直接读队尾。
// Time: O(n); storage O(n) for slice-backed queues, including possible append reallocations.
// 时间 O(n)；切片队列及可能的扩容需要 O(n) 存储，不能把底层队列空间也算成 O(1)。
//
// 先 Pop 取得栈顶，再 Push 放回队尾，恢复原栈内容和顺序；所以读取也需 O(n)，且要求非空。
// Pop then Push restores the value and order; peeking is O(n) and requires nonempty input.
func (s *MyStackLazy) Top() int {
	value := s.Pop()

	s.Push(value)

	return value
}

// Empty Report emptiness from the queue length alone.
// 只看队列长度判断栈是否为空。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (s *MyStackLazy) Empty() bool {
	return len(s.queue) == 0
}

// MyStackTwoQueues 3. Two queues with a backup: transfer all but the newest into a second queue, then swap roles.
// 3. 两个队列，备份队列中转：把除最新元素以外的全部转移到 backup，再交换角色。
// Time: Push amortized O(1), Pop/Top O(n), Empty O(1). Space: O(n).
// 时间复杂度：Push 均摊 O(1)，Pop/Top O(n)，Empty O(1)。空间复杂度：O(n)。
// main holds every element in arrival order: front = oldest = stack bottom, back = newest = stack top.
// main 按入栈先后保存全部元素：队首 = 最旧 = 栈底，队尾 = 最新 = 栈顶。
// backup is a staging area used only inside Pop, and is empty between operations.
// backup 只在 Pop 内部充当中转区，两次操作之间一定为空。
// The zero value is ready to use: two nil slices are an empty stack.
// 零值即可使用：两个 nil 切片就是一个空栈，无需构造函数。
type MyStackTwoQueues struct {
	main   []int
	backup []int
}

// Push Append x onto main; its back is the stack top.
// 把 x 追加到 main 队尾，队尾就是栈顶。
// Time: amortized O(1), worst O(n) on reallocation; backing storage O(n).
// 时间均摊 O(1)，扩容时最坏 O(n)；底层存储 O(n)。
//
// 追加到输入端，后续的取出操作负责调整顺序；append 均摊 O(1)，扩容时单次 O(n)。
// Append at the input end; removal handles reordering. append is amortized O(1), with O(n) time on reallocation.
func (s *MyStackTwoQueues) Push(x int) {
	s.main = append(s.main, x)
}

// Pop Drain every older element from main into backup, dequeue the leftover top, then swap the two queues.
// 把旧元素全部转移到 backup，弹出 main 里剩下的栈顶，再交换两个队列的角色。
// Time: O(n); storage O(n) for slice-backed queues, including possible append reallocations.
// 时间 O(n)；切片队列及可能的扩容需要 O(n) 存储，不能把底层队列空间也算成 O(1)。
//
// 把 main 中除最后一个之外的元素搬到 backup，剩下的最新元素就是栈顶；取出后交换两队列。
// Move all but main's newest element to backup; remove the newest as the top, then swap queues.
// 交换后 main 保存全部余项，backup 为空，恢复下次操作的前提；要求栈非空。
// After the swap main holds all remaining values and backup is empty again; require a nonempty stack.
func (s *MyStackTwoQueues) Pop() int {
	for len(s.main) > 1 {
		s.backup = append(s.backup, s.main[0])
		s.main = s.main[1:]
	}

	value := s.main[0]
	s.main = s.main[1:]

	s.main, s.backup = s.backup, s.main

	return value
}

// Top Reuse Pop then Push so the newest value is read without accessing the queue back.
// 复用 Pop 取出栈顶，再用 Push 加回，避免直接读队尾。
// Time: O(n); storage O(n) for slice-backed queues, including possible append reallocations.
// 时间 O(n)；切片队列及可能的扩容需要 O(n) 存储，不能把底层队列空间也算成 O(1)。
//
// Pop 后再 Push 将元素放回最新位置，恢复原顺序；依赖 Pop 恢复 backup 为空的不变量，要求非空。
// Pop then Push restores the newest value and original order; Pop also restores empty backup. Require a nonempty stack.
func (s *MyStackTwoQueues) Top() int {
	value := s.Pop()

	s.Push(value)

	return value
}

// Empty Report emptiness from main alone, because backup is empty between operations.
// 两次操作之间 backup 一定为空，所以只看 main 的长度。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (s *MyStackTwoQueues) Empty() bool {
	return len(s.main) == 0
}
