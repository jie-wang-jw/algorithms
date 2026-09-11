package _5_stack_queue

/*
232. 用栈实现队列 / Implement Queue Using Stacks

题目描述 / Problem Description
只使用栈的标准操作实现一个先进先出的队列，支持 Push、Pop、Peek 和 Empty。
Implement a first-in-first-out queue using only standard stack operations,
supporting Push, Pop, Peek, and Empty.

解法一：双栈惰性转移（推荐） / Method 1: Two Stacks, Lazy Transfer (Recommended)
使用 inStack 接收入队元素，使用 outStack 提供队首。当 outStack 为空时才把
inStack 的所有元素倒入其中，使最早入队的元素来到栈顶。
Use inStack for incoming elements and outStack for the queue front.
Only when outStack is empty, transfer all elements from inStack so the oldest element becomes its top.

关键逻辑：为什么这样做 / Why This Works
倒栈会反转先后顺序：按 1、2、3 入栈，依次弹出 3、2、1 再压入 outStack，
顶部就是最早来的 1。必须等 outStack 空了才转移，因为它剩余的元素都比 inStack 的新元素早来；
提前转移会把新元素压在旧元素上面，让后来者先出队，破坏先进先出。
Transfer reverses order: pushing 1,2,3 then moving popped values 3,2,1 puts oldest value 1 on top of outStack.
Transfer only when outStack is empty: its remaining items arrived before every item in inStack.
An early transfer would put newer values above older ones and violate FIFO.

MyQueue 由 Constructor 创建；它的零值也可以直接使用，因为两个 nil 切片就是空队列。
Pop 和 Peek 假定队列非空，与题目约定一致。
MyQueue is created by Constructor; its zero value is also usable, since two nil slices form an empty queue.
Pop and Peek assume a nonempty queue, matching the problem's guarantee.

解法二：入队时整理顺序 / Method 2: Eager Reordering on Push
MyQueueEager 只用一个主栈，并始终把最早入队的元素放在栈顶，于是 Pop 和 Peek 都是纯粹的栈顶操作。
Push 分三步：先把主栈里的旧元素全部弹入临时栈；再把新值压入空主栈，让它成为新的栈底；
最后按“从临时栈顶弹出、压回主栈”的顺序倒回旧元素。
倒出再倒回是两次反转，所以旧元素恢复原来的相对顺序，原本的最早值重新回到栈顶。
全过程只使用压栈、弹栈、读栈顶、查看长度，不从切片头部出队，否则就不是栈接口了。
MyQueueEager 没有构造函数，直接使用 Go 的零值即可：stack 为 nil 切片，Empty 立即返回 true。
Pop 和 Peek 同样假定队列非空。
MyQueueEager uses a single main stack and always keeps the oldest element on top,
so Pop and Peek are plain stack-top operations.
Push has three steps: drain every old element into a temporary stack; push the new value onto the now-empty
main stack so it becomes the new bottom; then restore the old elements by popping the temporary stack back.
Draining and restoring are two reversals, so the old elements regain their original relative order
and the previously oldest value returns to the top.
Only push, pop, top, and length are used, never dequeuing from the slice front, which would not be a stack interface.
MyQueueEager has no constructor and relies on its usable Go zero value: stack is a nil slice and Empty returns true.
Pop and Peek likewise assume a nonempty queue.

时间与空间复杂度 / Time and Space Complexity
n 为当前队列元素数。
解法一 MyQueue：Push 均摊 O(1)；Pop/Peek 均摊 O(1)、单次最坏 O(n)；Empty O(1)；存储 O(n)。
每个元素最多从 inStack 转移到 outStack 一次，所以一串操作的总工作量是线性的，
不存在把同一元素反复搬运的情形。
解法二 MyQueueEager：Push 每次都是最坏也是确定的 O(n)，两次搬运全部旧元素，
并额外分配 O(n) 的临时栈；Pop/Peek/Empty O(1)；存储 O(n)。
解法二不是均摊 O(1)：连续 n 次 Push 的总代价是 O(n²)，而解法一是 O(n)，因此默认优先解法一。
存储空间 O(n)；切片容量可保留到历史最大规模。
n is the current queue size.
Method 1 MyQueue: Push amortized O(1); Pop/Peek amortized O(1) with worst case O(n) per call;
Empty O(1); storage O(n). Each element transfers from inStack to outStack at most once,
making aggregate work linear, with no element ever moved twice.
Method 2 MyQueueEager: Push is a deterministic O(n) every time, moving all old elements twice
and allocating an O(n) temporary stack; Pop/Peek/Empty O(1); storage O(n).
Method 2 is not amortized O(1): n consecutive pushes cost O(n²) versus O(n) for Method 1,
so Method 1 is the default. Storage is O(n); slice capacity may remain at the historical peak size.
*/

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

// 1. 双栈惰性转移：推荐
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

// 2. 入队时整理顺序：把 O(n) 成本放在 Push
type MyQueueEager struct {
	// stack keeps the queue inverted: top = oldest = queue front, bottom = newest = queue back.
	// stack 以倒置方式保存队列：栈顶 = 最旧 = 队首，栈底 = 最新 = 队尾。
	//
	// The zero value is ready to use: a nil slice is an empty queue.
	// 零值即可使用：nil 切片就是一个空队列，无需构造函数。
	stack []int
}

// Push adds x to the back of the queue.
// Push 将 x 加入队尾。
func (q *MyQueueEager) Push(x int) {
	// temp holds the existing elements while the new bottom is installed.
	// temp 在安放新栈底期间暂存原有元素。
	temp := make([]int, 0, len(q.stack))

	// Step 1: drain the main stack, which reverses the order once.
	// 第一步：倒空主栈，这一步把顺序反转了一次。
	for len(q.stack) > 0 {
		last := len(q.stack) - 1
		temp = append(temp, q.stack[last])
		q.stack = q.stack[:last]
	}

	// Step 2: the new value goes in first, so it becomes the bottom, that is the queue back.
	// 第二步：新值最先压入空主栈，于是成为栈底，也就是队尾。
	q.stack = append(q.stack, x)

	// Step 3: restore the old elements, reversing a second time to recover their order.
	// 第三步：把旧元素倒回来，第二次反转使它们恢复原来的相对顺序。
	for len(temp) > 0 {
		last := len(temp) - 1
		q.stack = append(q.stack, temp[last])
		temp = temp[:last]
	}

	// The previously oldest element is back on top; only an empty queue leaves x on top.
	// 原本最早的元素重新回到栈顶；只有队列原本为空时，栈顶才是 x。
}

// Pop removes and returns the front element.
// Pop 删除并返回队首元素。
func (q *MyQueueEager) Pop() int {
	// Push already arranged the oldest element on top, so no transfer is needed here.
	// Push 已经把最早的元素安排在栈顶，这里不需要任何转移。
	last := len(q.stack) - 1
	value := q.stack[last]

	// Removing the top preserves the invariant for the next-oldest element.
	// 删除栈顶后，次早的元素自然成为新的栈顶，不变量继续成立。
	q.stack = q.stack[:last]

	return value
}

// Peek returns the front element without removing it.
// Peek 返回队首元素，但不删除它。
func (q *MyQueueEager) Peek() int {
	// Reading the top does not modify the stack, so the invariant is untouched.
	// 读取栈顶不修改栈，因此不变量保持不变。
	return q.stack[len(q.stack)-1]
}

// Empty returns whether the queue contains no elements.
// Empty 判断队列中是否没有元素。
func (q *MyQueueEager) Empty() bool {
	// A single stack holds everything, so one length check is enough.
	// 所有元素都在一个栈里，所以只需检查一个长度。
	return len(q.stack) == 0
}
