package _5_stack_queue

import "container/heap"

/*
239. 滑动窗口最大值 / Sliding Window Maximum

题目描述 / Problem Description
给定整数数组 nums 和窗口大小 k，窗口从数组最左侧每次向右移动一位，返回每个窗口中的最大值。
Given an integer array nums and a window size k, move the window one position at a
time from left to right and return the maximum value in every window.

解法一：单调递减队列（推荐） / Method 1: Monotonic Decreasing Deque (Recommended)
使用保存下标的单调递减队列。每轮删除队首过期下标，再删除队尾所有不大于当前值的下标并加入当前下标；
队首始终对应当前窗口最大值。
Use a decreasing monotonic deque of indices. Remove expired indices from the front,
remove values no greater than the current value from the back, then append the current index;
the front is always the maximum.

关键逻辑：为什么这样做 / Why This Works
队列同时保持两条性质：下标递增，所以过期元素只会在队首；对应值递减，所以有效候选中队首最大。
删掉较小旧值不会漏答案：只要旧值仍在今后的窗口里，更大且更晚加入的新值一定也在，
因此旧值永远不必胜出。相等时也保留更晚的下标，因为它能覆盖旧值剩下的全部有效窗口。
i-k+1>=0 即 i>=k-1 时才形成完整窗口。
The deque maintains increasing indices, so expiration occurs at the front, and decreasing values,
so its first valid candidate is largest. Discarding a smaller older value is safe:
every future window retaining it also retains the larger newer value. On equality keep the newer index,
which remains valid at least as long. A full window exists when i-k+1>=0, equivalently i>=k-1.

解法二：逐窗口暴力扫描 / Method 2: Brute Force per Window
maxSlidingWindowBruteForce 对每个左边界 left 重新扫描 [left,left+k-1] 求最大值，不复用上一窗口的信息。
循环条件写成 left+k<=len(nums) 而不是 left<len(nums)-k+1，避免 len(nums)-k+1 的减法在其他约定下溢出，
同时天然表达“窗口必须完整落在数组内”。
它是理解单调队列价值的基线：唯一浪费的就是相邻窗口重叠的 k-1 个元素被反复比较。
Rescan [left,left+k-1] for every left boundary, reusing nothing from the previous window.
The loop condition is left+k <= len(nums) rather than left < len(nums)-k+1, which avoids the subtraction
and directly expresses that a window must fit entirely inside the array.
This is the baseline that motivates the deque: its only waste is recomparing the k-1 overlapping elements.

解法三：惰性删除的大顶堆 / Method 3: Max-Heap with Lazy Deletion
maxSlidingWindowHeap 的堆元素是 [2]int{值, 下标} 对，每轮先插入当前元素，然后反复检查堆顶：
只要堆顶下标已经过期（下标 <= i-k）就弹出，直到堆顶有效。
有效堆顶是全堆最大值，而当前窗口的每个元素都还在堆里（被弹出的都已过期），所以它就是窗口最大值。
被压在堆内部的过期元素不主动删除，因为 container/heap 只能高效删除堆顶；
它们将来浮到堆顶时再被弹出，这就是“惰性删除”。
因此堆的规模可能长到 n 而不是 k，例如严格递减数组中任何元素都不会提前被挤掉。
循环不会掏空堆：刚插入的下标 i 满足 i>i-k（k>=1），永远不会被判为过期，所以它一定会终止。
值相等时 Less 的胜负不确定，但循环会持续弹出直到堆顶有效，因此结果不受并列影响。
Heap items are [2]int{value, index} pairs. Each round inserts the current element, then repeatedly pops
the root while its index has expired (index <= i-k), until the root is valid.
A valid root is the maximum over the whole heap, and every element of the current window is still in the heap
because only expired entries were ever popped, so the root is the window maximum.
Expired entries buried inside the heap are not removed eagerly, since container/heap can only delete the root
cheaply; they are popped later when they surface. That is the lazy deletion.
Heap size can therefore reach n rather than k, as in a strictly decreasing array where nothing is ever displaced.
The loop cannot empty the heap: the freshly inserted index i satisfies i > i-k for k >= 1 and can never expire,
which guarantees termination.
Ties in Less resolve arbitrarily, but the loop keeps popping until the root is valid, so ties do not affect results.

解法四：分块预处理前后缀最大值 / Method 4: Block Prefix and Suffix Maxima
maxSlidingWindowBlocks 把数组按 k 个元素切成块。
prefix[i] 是从 i 所在块的起点到 i 的最大值；suffix[i] 是从 i 到该块终点（或数组末尾）的最大值。
长度恰为 k 的窗口 [left,right]（right=left+k-1）最多跨越两个相邻块，
因此 max(suffix[left],prefix[right]) 正好覆盖 [left,right]：
suffix[left] 覆盖 left 到它所在块的末尾，prefix[right] 覆盖 right 所在块的开头到 right，两段相接且都不越界。
窗口与块恰好对齐时（left%k==0），right 与 left 同块，两项都代表该整块，结果仍然正确，也不会引入窗口外元素。
suffix 在数组末尾不足一整块时以 len(nums)-1 收尾；这不影响正确性，
因为任何完整窗口的 right 都不超过 len(nums)-1。
Split the array into blocks of k elements.
prefix[i] is the maximum from the start of i's block through i; suffix[i] is the maximum from i through
the end of that block, or the end of the array.
A window [left,right] of length exactly k, with right = left+k-1, spans at most two adjacent blocks,
so max(suffix[left],prefix[right]) covers [left,right] exactly: suffix[left] covers left to its block end,
prefix[right] covers right's block start to right, and the two ranges meet without exceeding the window.
When the window is block-aligned (left%k == 0), right shares left's block and both terms describe that whole block,
which is still correct and still introduces no outside element.
A trailing partial block ends suffix at len(nums)-1, which is harmless because every complete window's
right boundary is at most len(nums)-1.

时间与空间复杂度 / Time and Space Complexity
n = len(nums)，k 为窗口长度，输出长度恒为 n-k+1；四种解法都不修改 nums。
解法一 maxSlidingWindow：时间 O(n)，每个下标最多入队、出队一次，所以两个内层循环的总工作量是线性的；
辅助队列空间 O(k)，输出 O(n-k+1)，合计 O(n)。
解法二 maxSlidingWindowBruteForce：时间 O((n-k+1)k)，k 约为 n/2 时达到 O(n²) 峰值；
辅助空间 O(1)，输出 O(n-k+1)。
解法三 maxSlidingWindowHeap：时间 O(n log n)（用 O(n log(n+1)) 表述可涵盖 n=1）；
辅助空间 O(n)，因为惰性删除让堆规模由 n 而不是 k 决定，输出 O(n-k+1)。
解法四 maxSlidingWindowBlocks：时间 O(n)，三次线性遍历；
辅助空间 O(n)，prefix 与 suffix 各占 n 个整数，输出 O(n-k+1)。
四种解法都沿用题目约定 1<=k<=n；通常优先解法一，它是唯一同时做到 O(n) 时间和 O(k) 辅助空间的写法。
切片 append 的分配与复制按均摊计算。
n = len(nums), k is the window length, the output always has n-k+1 entries, and none of the four methods mutates nums.
Method 1 maxSlidingWindow: time O(n), since each index enters and leaves the deque at most once, giving the
inner loops linear aggregate work; auxiliary deque space O(k) plus output O(n-k+1), total O(n).
Method 2 maxSlidingWindowBruteForce: time O((n-k+1)k), peaking at O(n²) near k = n/2;
auxiliary space O(1) plus output O(n-k+1).
Method 3 maxSlidingWindowHeap: time O(n log n), written as O(n log(n+1)) to include n = 1;
auxiliary space O(n), because lazy deletion makes heap size depend on n rather than k, plus output O(n-k+1).
Method 4 maxSlidingWindowBlocks: time O(n) across three linear passes;
auxiliary space O(n) for the two n-element prefix and suffix arrays, plus output O(n-k+1).
All four assume the problem's 1 <= k <= n. Prefer Method 1: it is the only one achieving both O(n) time
and O(k) auxiliary space. Slice allocation and copying are amortized.
*/

/*滑动窗口 + 单调队列
Sliding window + monotonic deque

左边 = 队头 front
右边 = 队尾 back
新下标从右边加入：
过期下标从左边删除：

三个固定步骤
第一步：删除已经离开窗口的队头
第二步：删除队尾所有不大于当前值的元素
第三步：加入当前下标

暴力法对每个窗口重新检查 k 个数字，时间 O((n-k+1)*k)；下面的单调队列复用候选，不会逐窗口重新扫描。
Brute force scans k values per window in O((n-k+1)*k); the deque below reuses candidates instead of rescanning each window.

每轮的三个步骤 / Three Steps per Iteration
第一步：删除过期下标 / Remove Expired Indices
当前窗口的有效下标范围是：
The current window’s valid index range is:
[i-k+1, i]
如果队首下标小于 i-k+1，说明它已经离开窗口，需要删除。
If the front index is smaller than i-k+1, it has left the window and must be removed.
等价条件：
Equivalent condition:
deque[0] <= i-k
第二步：维护单调递减 / Maintain Decreasing Order
加入 nums[i] 前，从队尾删除所有小于或等于 nums[i] 的元素。
Before adding nums[i], remove every value from the back that is less than or equal to nums[i].
原因是这些元素以后不可能成为窗口最大值：
Those elements can never become a future window maximum because:
- nums[i] 大于或等于它们；
- nums[i] is greater than or equal to them;
- nums[i] 比它们更晚离开窗口。
- nums[i] leaves the window later.
因此可以永久删除这些较弱的候选元素。
Therefore, these weaker candidates can be discarded permanently.
第三步：加入当前下标 / Add the Current Index
把 i 加入队尾。
Append i to the back.
当第一个完整窗口形成后，将队首对应的数字加入结果。
Once a complete window has formed, append the value represented by the deque front to the result.


I use a monotonic deque to store indices whose corresponding values are in decreasing order.
Before adding a new index, I remove expired indices from the front and remove smaller values from the back.
Therefore,the front of the deque always represents the maximum value in the current window.
*/

// 1. 单调递减队列：推荐
func maxSlidingWindow(nums []int, k int) []int {
	// deque stores indices instead of values.
	// deque 保存数组下标，而不是直接保存数字。
	//
	// The corresponding values remain in decreasing order.
	// 这些下标对应的数字保持单调递减。
	deque := make([]int, 0, k)

	// There are len(nums)-k+1 complete windows.
	// 一共有 len(nums)-k+1 个完整窗口。
	result := make([]int, 0, len(nums)-k+1)

	// i is the right boundary of the current window.
	// i 是当前窗口的右边界。
	for i := range nums {
		// Step 1: remove indices that have left the window.
		// 第一步：删除已经离开当前窗口的下标。
		//
		// The valid window starts at i-k+1, so indices <= i-k are expired.
		// 当前窗口从 i-k+1 开始，因此下标 <= i-k 的元素已经过期。
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// Step 2: remove weaker candidates from the back.
		// 第二步：从队尾删除不可能成为最大值的候选元素。
		//
		// nums[i] is at least as large and will leave the window later.
		// nums[i] 不小于这些元素，而且会更晚离开窗口。
		for len(deque) > 0 &&
			nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Step 3: append the current index.
		// 第三步：将当前下标加入队尾。
		deque = append(deque, i)

		// The first complete window is formed when i reaches k-1.
		// 当 i 到达 k-1 时，第一个完整窗口形成。
		if i >= k-1 {
			// The front always represents the current maximum.
			// 队首始终对应当前窗口的最大值。
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

// 2. 逐窗口暴力扫描：理解单调队列价值的基线
func maxSlidingWindowBruteForce(nums []int, k int) []int {
	// There are len(nums)-k+1 complete windows.
	// 一共有 len(nums)-k+1 个完整窗口。
	result := make([]int, 0, len(nums)-k+1)

	// left is the window's left boundary; left+k <= len(nums) keeps the window inside the array.
	// left 是窗口左边界；left+k <= len(nums) 保证窗口完整落在数组内。
	for left := 0; left+k <= len(nums); left++ {
		// Seed with the leftmost value so no neutral element is needed.
		// 用窗口最左元素作为初值，就不必设定一个人为的极小初值。
		best := nums[left]

		// Rescan the remaining k-1 values, reusing nothing from the previous window.
		// 重新扫描剩下的 k-1 个元素，完全不复用上一个窗口的结果。
		for j := left + 1; j < left+k; j++ {
			best = max(best, nums[j])
		}

		result = append(result, best)
	}

	return result
}

// 3. 惰性删除的大顶堆：只在堆顶过期时弹出

// windowMaxHeap is a max-heap ordered by value, not a fully sorted slice.
// windowMaxHeap 是按数值维护父子大小关系的大顶堆，不是完全有序的切片。
//
// Each pair stores a value and its original index.
// 每一项保存数值及其原始下标。
type windowMaxHeap [][2]int

// Len reports the heap size, which container/heap uses to locate the last slot.
// Len 返回堆中的元素数量，container/heap 借它定位末尾位置。
func (h windowMaxHeap) Len() int {
	return len(h)
}

// Less places the larger value closer to the root, making this a max-heap.
// Less 让数值较大的元素更接近堆顶，因此这是一个大顶堆。
func (h windowMaxHeap) Less(i, j int) bool {
	// Only the value is compared; equal values may sit in either order.
	// 只比较数值，数值相等时两者先后顺序不确定。
	return h[i][0] > h[j][0]
}

// Swap exchanges two heap elements; container/heap calls it while sifting.
// Swap 交换堆中的两个元素；container/heap 在上浮和下沉时调用它。
func (h windowMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a new element to the end of the underlying slice.
// heap.Push calls this first, then sifts the new element up using Less and Swap.
// Push 将新元素加入底层切片末尾。
// heap.Push 先调用它追加元素，再用 Less 和 Swap 把该元素上浮到正确位置。
func (h *windowMaxHeap) Push(x any) {
	// The pointer receiver is required because appending may reallocate the slice.
	// 必须使用指针接收者，因为 append 可能重新分配底层数组。
	*h = append(*h, x.([2]int))
}

// Pop removes and returns the last element of the underlying slice.
// container/heap moves the root to the end and repairs the rest before calling this method,
// so Pop never needs to search for the maximum itself.
// Pop 删除并返回底层切片末尾的元素。
// container/heap 调用它之前，会先把原堆顶移到切片末尾并调整剩余部分，
// 所以 Pop 自己不需要去寻找最大值。
func (h *windowMaxHeap) Pop() any {
	old := *h
	last := len(old) - 1

	// Save the element that will be returned.
	// 保存即将返回的元素。
	value := old[last]

	// Remove the last element from the slice.
	// 从切片中删除最后一个元素。
	*h = old[:last]

	return value
}

func maxSlidingWindowHeap(nums []int, k int) []int {
	// An empty literal is already a valid heap, so heap.Init is unnecessary.
	// 空字面量本身就是合法的堆，因此不需要调用 heap.Init。
	h := &windowMaxHeap{}

	result := make([]int, 0, len(nums)-k+1)

	// i is the right boundary of the current window.
	// i 是当前窗口的右边界。
	for i, value := range nums {
		// Insert the current element together with its index.
		// 把当前元素和它的下标一起插入堆中。
		heap.Push(h, [2]int{value, i})

		// Drop expired roots only; expired entries buried deeper stay until they surface.
		// 只弹出过期的堆顶；埋在堆内部的过期元素等浮到堆顶时再处理。
		//
		// The index just inserted satisfies i > i-k for k >= 1, so the heap cannot be emptied.
		// 刚插入的下标满足 i > i-k（k >= 1），所以这个循环不会掏空堆。
		for (*h)[0][1] <= i-k {
			heap.Pop(h)
		}

		// The first complete window is formed when i reaches k-1.
		// 当 i 到达 k-1 时，第一个完整窗口形成。
		if i >= k-1 {
			// Every window element is still in the heap, so a valid root is the window maximum.
			// 窗口内的元素都还在堆中，因此有效的堆顶就是窗口最大值。
			result = append(result, (*h)[0][0])
		}
	}

	return result
}

// 4. 分块预处理前后缀最大值：用两次线性扫描代替队列
func maxSlidingWindowBlocks(nums []int, k int) []int {
	// prefix[i] is the maximum from the start of i's block through i.
	// suffix[i] is the maximum from i through the end of i's block.
	// prefix[i] 是 i 所在块的起点到 i 的最大值。
	// suffix[i] 是 i 到该块终点的最大值。
	prefix, suffix := make([]int, len(nums)), make([]int, len(nums))

	for i, value := range nums {
		if i%k == 0 {
			// A block start has no earlier element in its own block.
			// 块起点在本块内没有更早的元素，因此直接取自身。
			prefix[i] = value
		} else {
			// Extend the running maximum within the same block.
			// 在同一个块内继续累积最大值。
			prefix[i] = max(prefix[i-1], value)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || (i+1)%k == 0 {
			// A block end, or the array end for a trailing partial block, starts a new run.
			// 块终点，或末尾不足一整块时的数组末端，都要重新开始累积。
			suffix[i] = nums[i]
		} else {
			// Extend the running maximum backward within the same block.
			// 在同一个块内向左继续累积最大值。
			suffix[i] = max(suffix[i+1], nums[i])
		}
	}

	result := make([]int, 0, len(nums)-k+1)

	for left := 0; left+k <= len(nums); left++ {
		// A length-k window spans at most two adjacent blocks.
		// 长度恰为 k 的窗口最多跨越两个相邻块。
		right := left + k - 1

		// suffix[left] covers left to its block end, prefix[right] covers right's block start to right.
		// The two ranges meet exactly at the block boundary and never reach outside the window.
		// suffix[left] 覆盖 left 到它所在块的末尾，prefix[right] 覆盖 right 所在块的开头到 right。
		// 两段正好在块边界相接，也都不会越出窗口。
		result = append(result, max(suffix[left], prefix[right]))
	}

	return result
}
