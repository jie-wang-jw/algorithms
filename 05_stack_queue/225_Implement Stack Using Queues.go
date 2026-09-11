package _5_stack_queue

/*
225. 用队列实现栈 / Implement Stack Using Queues

题目描述 / Problem Description
只使用队列的标准操作实现一个后进先出的栈，支持 Push、Pop、Top 和 Empty。
Implement a last-in-first-out stack using only standard queue
operations, supporting Push, Pop, Top, and Empty.

解法一：入栈时旋转（推荐） / Method 1: Rotate on Push (Recommended)
使用一个队列。每次 Push 新元素后，把之前的所有元素从队首移动到队尾，
使新元素来到队首；这样 Pop 和 Top 都可以直接操作队首。
Use one queue. After each Push, rotate all older elements from
the front to the back so the new element becomes the front;
Pop and Top can then use the front directly.

关键逻辑：为什么这样做 / Why This Works
入栈前，队首到队尾已经是栈顶到栈底。新值追加到尾部后，只把 size 个旧元素逐个移到后面，
新值就来到队首，旧元素的相对顺序不变。例如 [2,1] 加入 3： [2,1,3] -> [1,3,2] -> [3,2,1]。
因此只能旋转旧元素数量，不能把新元素也转走。
Before Push, the queue runs from stack top to bottom. Append the new value, then rotate exactly size old items;
the new value reaches the front while old items retain their order.
Example: [2,1] plus 3 becomes [2,1,3] -> [1,3,2] -> [3,2,1]. Rotating the new item too would undo its placement.

MyStack 由 StackConstructor 创建；它的零值也可以直接使用，因为空队列就是空栈。
Pop 和 Top 假定栈非空，与题目约定一致，不对空栈返回哨兵值。
MyStack is created by StackConstructor; its zero value is also usable, since an empty queue is an empty stack.
Pop and Top assume a nonempty stack, matching the problem's guarantee, and return no sentinel for an empty stack.

解法二：把旋转推迟到出栈 / Method 2: Rotate on Pop
MyStackLazy.Push 仅追加，队列顺序始终是“最旧到最新”，队尾才是栈顶。
Pop 把前 size-1 个元素从队首移到队尾，此时队首就是最新入栈的值，直接出队即可；
剩下的元素因为整体旋转过一圈，相对顺序不变，仍然是“最旧到最新”。
Top 复用 Pop 取出该值，再用 Push 把它加回队尾，恰好还原原来的排列。
只使用队列的入队、出队、查看长度操作，不直接读队尾，否则就不是队列接口了。
MyStackLazy 没有构造函数，直接使用 Go 的零值即可：queue 为 nil 切片，
append 会在首次 Push 时分配，len(nil) 为 0 使 Empty 立即返回 true。
Pop 和 Top 同样假定栈非空。
Push appends only, so the queue always runs oldest-to-newest and the back is the stack top.
Pop rotates size-1 front elements to the back, leaving the newest value at the front to dequeue;
the remaining elements rotate a full cycle, so their relative order is unchanged.
Top reuses Pop to remove that value, then Push to re-enqueue it, restoring the original arrangement exactly.
It uses only enqueue, dequeue, and length, never reading the back directly, which would not be a queue interface.
MyStackLazy has no constructor and relies on its usable Go zero value: queue is a nil slice,
append allocates on the first Push, and len(nil) == 0 makes Empty return true immediately.
Pop and Top likewise assume a nonempty stack.

解法三：两个队列，备份队列中转 / Method 3: Two Queues with a Backup
MyStackTwoQueues 是代码随想录先介绍的写法：main 保存元素，backup 只在 Pop 期间充当中转区。
Pop 把 main 中除最新元素以外的全部元素依次转移到 backup，此时 main 只剩栈顶，直接出队；
最后交换两个队列的角色，让刚才的 backup 成为新的 main，而被掏空的旧 main 成为新的 backup。
它与解法二的区别不是复杂度，而是“元素去哪里”：解法二把出队的元素重新排到同一个队列的尾部，
解法三把它们放进另一个队列，因此任何时刻都不会出现“同一个队列里既有旧元素又有刚转移的元素”。
关键不变量：两次操作之间 backup 一定为空，所以 Empty 只需检查 main 的长度。
交换之后旧 main 已经被掏空，这个不变量自动维持，不需要额外清理。
MyStackTwoQueues is the form 代码随想录 introduces first: main holds the elements, and backup
serves only as a staging area during Pop.
Pop moves every element except the newest from main into backup, leaving just the stack top to dequeue,
then swaps the two roles so the former backup becomes the new main and the drained main becomes the new backup.
The difference from Method 2 is not complexity but where elements go: Method 2 re-enqueues them at the back
of the same queue, while Method 3 places them in a second queue, so a single queue never simultaneously
holds both untouched old elements and just-transferred ones.
Key invariant: backup is always empty between operations, so Empty needs to check only main's length.
The swap leaves the drained main empty, which maintains that invariant automatically with no extra cleanup.

MyStackTwoQueues 同样没有构造函数，零值可用：两个 nil 切片就是一个空栈。
MyStackTwoQueues likewise has no constructor and is usable as its zero value: two nil slices form an empty stack.

时间与空间复杂度 / Time and Space Complexity
n 为当前栈内元素数。
解法一 MyStack：Push 时间 O(n)，旋转所有旧元素；Pop、Top、Empty 时间 O(1)；存储空间 O(n)。
解法二 MyStackLazy：Push 均摊 O(1)，只做一次 append；Pop 时间 O(n)，旋转 n-1 个元素；
Top 时间 O(n)，等于一次 Pop 加一次 Push；Empty 时间 O(1)；存储空间 O(n)。
解法三 MyStackTwoQueues：Push 均摊 O(1)；Pop 时间 O(n)，转移 n-1 个元素；
Top 时间 O(n)，等于一次 Pop 加一次 Push；Empty 时间 O(1)；存储空间 O(n)，
两个队列的元素总数始终是 n，因为 backup 在操作之间为空，所以不会真的翻倍。
解法一与解法二的总成本相同，区别只是把 O(n) 放在入栈还是出栈；读多写少时用解法一，写多读少时用解法二。
解法二与解法三的复杂度完全一致，解法三多一个队列，胜在语义更贴近“只用队列操作”的教学演示。
Go append 的扩容成本按均摊分析计入；s.queue = s.queue[1:] 使出队后的容量无法复用，
底层数组随 append 迁移，峰值内存仍是 O(n)。
n is the current element count.
Method 1 MyStack: Push O(n) to rotate older elements; Pop, Top, and Empty O(1); storage O(n).
Method 2 MyStackLazy: Push amortized O(1) for a single append; Pop O(n) rotating n-1 elements;
Top O(n), equal to one Pop plus one Push; Empty O(1); storage O(n).
Method 3 MyStackTwoQueues: Push amortized O(1); Pop O(n) transferring n-1 elements;
Top O(n), equal to one Pop plus one Push; Empty O(1); storage O(n), because the two queues together
always hold n elements: backup is empty between operations, so nothing actually doubles.
Methods 1 and 2 have the same total cost and differ only in where the O(n) work sits; prefer Method 1
when reads dominate and Method 2 when writes dominate.
Methods 2 and 3 have identical complexity; Method 3 spends an extra queue to make the
"queue operations only" demonstration more explicit.
Go append growth costs are amortized; s.queue = s.queue[1:] forfeits reuse of the dequeued prefix,
so the backing array migrates on append while peak memory stays O(n).
*/

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

// 1. 入栈时旋转：推荐
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
	for range size {
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

// 2. 出栈时旋转：把 O(n) 成本推迟到 Pop
type MyStackLazy struct {
	// queue stores elements in arrival order:
	// front = oldest = stack bottom, back = newest = stack top.
	// queue 按照入栈先后保存元素：
	// 队首 = 最旧 = 栈底，队尾 = 最新 = 栈顶。
	//
	// The zero value is ready to use: a nil slice is an empty stack.
	// 零值即可使用：nil 切片就是一个空栈，无需构造函数。
	queue []int
}

// Push adds x onto the top of the stack.
// Push 将 x 压入栈顶。
func (s *MyStackLazy) Push(x int) {
	// Appending is all that is needed because the back represents the top.
	// 队尾就代表栈顶，所以只需追加，不必立刻整理顺序。
	s.queue = append(s.queue, x)
}

// Pop removes and returns the top element.
// Pop 删除并返回栈顶元素。
func (s *MyStackLazy) Pop() int {
	// Rotate every element except the newest one to the back.
	// 把除最新元素以外的所有元素移动到队尾。
	//
	// A negative count cannot occur here: an empty stack is out of contract,
	// and Go's range over an int simply performs zero iterations.
	// 这里不会出现负数轮次：空栈不在约定范围内，
	// 而且 Go 对负整数的 range 直接执行零次循环。
	for range len(s.queue) - 1 {
		// Read the front element.
		// 读取队首元素。
		front := s.queue[0]

		// Remove it from the front.
		// 将它从队首删除。
		s.queue = s.queue[1:]

		// Re-enqueue it behind the newest element.
		// 把它重新排到最新元素之后。
		s.queue = append(s.queue, front)
	}

	// After a full rotation of the older elements, the newest one is at the front.
	// 旧元素整体旋转一圈后，最新入栈的元素来到队首。
	value := s.queue[0]
	s.queue = s.queue[1:]

	// The survivors rotated as a block, so they remain ordered oldest to newest.
	// 剩下的元素整体旋转，相对顺序不变，仍然是“最旧到最新”。
	return value
}

// Top returns the top element without removing it.
// Top 返回栈顶元素，但不删除它。
func (s *MyStackLazy) Top() int {
	// A queue cannot read its back directly, so remove the top the same way Pop does.
	// 队列接口不能直接读队尾，所以用与 Pop 相同的方式把栈顶取出来。
	value := s.Pop()

	// Re-enqueueing it restores the exact arrangement Pop consumed.
	// 再入队一次，恰好还原 Pop 之前的排列。
	s.Push(value)

	return value
}

// Empty returns whether the stack contains no elements.
// Empty 判断栈中是否没有元素。
func (s *MyStackLazy) Empty() bool {
	return len(s.queue) == 0
}

// 3. 两个队列：备份队列中转，不在同一个队列里回排
type MyStackTwoQueues struct {
	// main holds every element in arrival order:
	// front = oldest = stack bottom, back = newest = stack top.
	// main 按入栈先后保存全部元素：
	// 队首 = 最旧 = 栈底，队尾 = 最新 = 栈顶。
	main []int

	// backup is a staging area used only inside Pop, and is empty between operations.
	// backup 只在 Pop 内部充当中转区，两次操作之间一定为空。
	backup []int

	// The zero value is ready to use: two nil slices are an empty stack.
	// 零值即可使用：两个 nil 切片就是一个空栈，无需构造函数。
}

// Push adds x onto the top of the stack.
// Push 将 x 压入栈顶。
func (s *MyStackTwoQueues) Push(x int) {
	// The back of main represents the top, so appending is enough.
	// main 的队尾就代表栈顶，所以只需追加。
	s.main = append(s.main, x)
}

// Pop removes and returns the top element.
// Pop 删除并返回栈顶元素。
func (s *MyStackTwoQueues) Pop() int {
	// Move everything except the newest element into the backup queue.
	// 把除最新元素以外的所有元素转移到备份队列。
	for len(s.main) > 1 {
		// Dequeue from the front of main and enqueue at the back of backup.
		// 从 main 队首出队，再从 backup 队尾入队。
		s.backup = append(s.backup, s.main[0])
		s.main = s.main[1:]
	}

	// main now holds exactly one element, which is the stack top.
	// 此时 main 中只剩一个元素，它就是栈顶。
	value := s.main[0]
	s.main = s.main[1:]

	// Swap roles: backup already holds the survivors in their original order,
	// and the drained main becomes the empty backup for the next call.
	// 交换角色：backup 已经按原顺序保存了剩余元素，
	// 而被掏空的 main 成为下一次调用的空 backup。
	s.main, s.backup = s.backup, s.main

	return value
}

// Top returns the top element without removing it.
// Top 返回栈顶元素，但不删除它。
func (s *MyStackTwoQueues) Top() int {
	// A queue cannot read its back directly, so take the top out the way Pop does.
	// 队列接口不能直接读队尾，所以用与 Pop 相同的方式把栈顶取出来。
	value := s.Pop()

	// Re-enqueueing puts it back at the top, restoring the arrangement Pop consumed.
	// 重新入队会把它放回栈顶，恰好还原 Pop 之前的排列。
	s.Push(value)

	return value
}

// Empty returns whether the stack contains no elements.
// Empty 判断栈中是否没有元素。
func (s *MyStackTwoQueues) Empty() bool {
	// backup is empty between operations, so main's length alone decides.
	// 两次操作之间 backup 一定为空，因此只看 main 的长度就够了。
	return len(s.main) == 0
}
