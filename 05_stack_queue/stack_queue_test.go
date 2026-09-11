package _5_stack_queue

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

// bracketTestCases is shared by every 20 implementation.
// 20 题的所有实现共用这一组用例。
//
// Inputs contain only the six bracket characters, which is exactly what the problem guarantees.
// The two implementations classify an unknown character differently, so other characters
// are deliberately excluded as out of contract.
// 输入只包含六种括号字符，这正是题目给定的约束。
// 两种实现对未知字符的归类方式不同，因此这里故意不测试约定之外的字符。
func bracketTestCases() []struct {
	name string
	s    string
	want bool
} {
	return []struct {
		name string
		s    string
		want bool
	}{
		{name: "single pair", s: "()", want: true},
		{name: "different adjacent pairs", s: "()[]{}", want: true},
		{name: "nested pairs", s: "{[()]}", want: true},
		{name: "deeply nested", s: "[({})]", want: true},
		{name: "mismatched type", s: "(]", want: false},
		{name: "wrong closing order", s: "([)]", want: false},
		{name: "extra opening bracket", s: "(()", want: false},
		{name: "extra closing bracket", s: "())", want: false},
		// Even length with only openings: only the final nonempty-stack check can reject this.
		// 长度为偶数但全是左括号：只能靠最后的“栈非空”检查判错。
		{name: "only opening brackets", s: "((", want: false},
		// Even length with only closings: the empty-stack check must reject the first character.
		// 长度为偶数但全是右括号：必须在第一个字符处靠“栈为空”检查判错。
		{name: "only closing brackets", s: "))", want: false},
		{name: "closing before opening", s: "}{", want: false},
		{name: "empty string", s: "", want: true},
	}
}

func TestIsValid(t *testing.T) {
	// Cover valid nesting and the three invalid cases:
	// extra opening brackets, extra closing brackets, and mismatched brackets.
	// 覆盖有效嵌套以及三种无效情况：左括号多余、右括号多余和括号不匹配。
	//
	// The two methods translate between bracket types at opposite moments,
	// so they must still agree on every verdict.
	// 两种解法在不同时刻完成括号类型的转换，但对每个判定结果必须完全一致。
	implementations := []struct {
		name string
		fn   func(string) bool
	}{
		{name: "expected closing stack", fn: isValid},
		{name: "hash map", fn: isValidWithMap},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range bracketTestCases() {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.s); got != tt.want {
						t.Fatalf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
					}
				})
			}
		})
	}
}

// stackUnderTest is the common surface of the two 225 implementations.
// stackUnderTest 是 225 题两种实现共有的对外接口。
//
// Both satisfy it through pointer receivers, so the harness always stores a pointer.
// 两种实现都通过指针接收者满足它，因此测试统一持有指针。
type stackUnderTest interface {
	Push(x int)
	Pop() int
	Top() int
	Empty() bool
}

func TestMyStack(t *testing.T) {
	// External behaviour must be identical; only the placement of the O(n) rotation differs.
	// 对外行为必须完全一致，区别只在于 O(n) 旋转发生在入栈还是出栈。
	implementations := []struct {
		name string
		make func() stackUnderTest
	}{
		{
			name: "rotate on push",
			make: func() stackUnderTest {
				stack := StackConstructor()
				return &stack
			},
		},
		{
			// This one has no constructor and must work straight from its zero value.
			// 这种实现没有构造函数，必须直接从零值开始就能工作。
			name: "rotate on pop",
			make: func() stackUnderTest {
				return &MyStackLazy{}
			},
		},
		{
			// Empty here reads only main, so a leaked element in backup would go unnoticed;
			// the interleaved sequences below are what catch a broken role swap.
			// 这种实现的 Empty 只看 main，如果元素残留在 backup 里就不会被发现；
			// 下面交错的操作序列正是用来暴露角色交换写错的情况。
			name: "two queues",
			make: func() stackUnderTest {
				return &MyStackTwoQueues{}
			},
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			// Verify LIFO order and state changes after interleaved operations.
			// 验证后进先出顺序，以及交错操作后的状态变化。
			stack := implementation.make()

			if !stack.Empty() {
				t.Fatal("a new stack should be empty / 新栈应该为空")
			}

			stack.Push(1)
			stack.Push(2)
			if got := stack.Top(); got != 2 {
				t.Fatalf("Top() = %d, want 2", got)
			}

			// Top must not consume the element it reports; the lazy version rebuilds the queue.
			// Top 不能消耗它返回的元素；惰性版本会重新排列队列，必须完整还原。
			if got := stack.Top(); got != 2 {
				t.Fatalf("second Top() = %d, want 2", got)
			}

			if got := stack.Pop(); got != 2 {
				t.Fatalf("Pop() = %d, want 2", got)
			}
			if got := stack.Top(); got != 1 {
				t.Fatalf("Top() after Pop() = %d, want 1", got)
			}

			// Pushing while elements remain must not disturb the existing order.
			// 栈内仍有元素时继续入栈，不能打乱已有元素的顺序。
			stack.Push(3)
			if got := stack.Pop(); got != 3 {
				t.Fatalf("Pop() after Push(3) = %d, want 3", got)
			}
			if got := stack.Pop(); got != 1 {
				t.Fatalf("final Pop() = %d, want 1", got)
			}
			if !stack.Empty() {
				t.Fatal("stack should be empty after all pops / 所有元素弹出后栈应该为空")
			}

			// A single element exercises the degenerate rotation of zero items.
			// 单个元素用于检验“旋转零次”这种退化情况。
			stack.Push(7)
			if got := stack.Top(); got != 7 {
				t.Fatalf("Top() of single-element stack = %d, want 7", got)
			}
			if got := stack.Pop(); got != 7 {
				t.Fatalf("Pop() of single-element stack = %d, want 7", got)
			}
			if !stack.Empty() {
				t.Fatal("stack should be empty again / 栈应该再次为空")
			}

			// A longer run stresses the rotation invariant beyond the hand-checked cases.
			// 更长的操作序列用来检验旋转不变量，而不只是手算的小例子。
			for value := 1; value <= 5; value++ {
				stack.Push(value)
			}
			for want := 5; want >= 1; want-- {
				if got := stack.Top(); got != want {
					t.Fatalf("Top() during drain = %d, want %d", got, want)
				}
				if got := stack.Pop(); got != want {
					t.Fatalf("Pop() during drain = %d, want %d", got, want)
				}
			}
			if !stack.Empty() {
				t.Fatal("stack should be empty after draining / 全部弹出后栈应该为空")
			}
		})
	}
}

// queueUnderTest is the common surface of the two 232 implementations.
// queueUnderTest 是 232 题两种实现共有的对外接口。
type queueUnderTest interface {
	Push(x int)
	Pop() int
	Peek() int
	Empty() bool
}

func TestMyQueue(t *testing.T) {
	// Both implementations must agree; one defers the O(n) work to Pop, the other pays it on Push.
	// 两种实现的外部行为必须一致：一种把 O(n) 推迟到 Pop，另一种在 Push 时就付出代价。
	implementations := []struct {
		name string
		make func() queueUnderTest
	}{
		{
			name: "lazy transfer",
			make: func() queueUnderTest {
				queue := Constructor()
				return &queue
			},
		},
		{
			// This one has no constructor and must work straight from its zero value.
			// 这种实现没有构造函数，必须直接从零值开始就能工作。
			name: "eager reordering",
			make: func() queueUnderTest {
				return &MyQueueEager{}
			},
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			// Verify FIFO order, including pushes performed while outStack still has data.
			// 验证先进先出顺序，包括 outStack 尚有元素时继续入队的情况。
			queue := implementation.make()

			if !queue.Empty() {
				t.Fatal("a new queue should be empty / 新队列应该为空")
			}

			queue.Push(1)
			queue.Push(2)
			if got := queue.Peek(); got != 1 {
				t.Fatalf("Peek() = %d, want 1", got)
			}

			// Peek must not consume the front element.
			// Peek 不能消耗队首元素。
			if got := queue.Peek(); got != 1 {
				t.Fatalf("second Peek() = %d, want 1", got)
			}

			if got := queue.Pop(); got != 1 {
				t.Fatalf("Pop() = %d, want 1", got)
			}

			// Pushing here leaves the out-stack nonempty, so a premature refill would
			// bury the older value 2 underneath the newer value 3 and break FIFO.
			// 此时 outStack 非空，如果提前转移，就会把较新的 3 压在较旧的 2 上面，破坏先进先出。
			queue.Push(3)
			if got := queue.Peek(); got != 2 {
				t.Fatalf("Peek() after Push(3) = %d, want 2", got)
			}
			if got := queue.Pop(); got != 2 {
				t.Fatalf("second Pop() = %d, want 2", got)
			}
			if got := queue.Pop(); got != 3 {
				t.Fatalf("third Pop() = %d, want 3", got)
			}
			if !queue.Empty() {
				t.Fatal("queue should be empty after all pops / 所有元素出队后队列应该为空")
			}

			// A single element exercises the smallest nonempty transfer.
			// 单个元素用于检验最小规模的转移。
			queue.Push(4)
			if got := queue.Peek(); got != 4 {
				t.Fatalf("Peek() of single-element queue = %d, want 4", got)
			}
			if got := queue.Pop(); got != 4 {
				t.Fatalf("Pop() of single-element queue = %d, want 4", got)
			}
			if !queue.Empty() {
				t.Fatal("queue should be empty again / 队列应该再次为空")
			}

			// Interleave pushes and pops so a refill happens with items still pending on both sides.
			// 交错入队和出队，让转移发生在两侧都还有元素的时刻。
			for value := 1; value <= 4; value++ {
				queue.Push(value)
			}
			if got := queue.Pop(); got != 1 {
				t.Fatalf("Pop() during interleaving = %d, want 1", got)
			}
			queue.Push(5)
			for want := 2; want <= 5; want++ {
				if got := queue.Peek(); got != want {
					t.Fatalf("Peek() during drain = %d, want %d", got, want)
				}
				if got := queue.Pop(); got != want {
					t.Fatalf("Pop() during drain = %d, want %d", got, want)
				}
			}
			if !queue.Empty() {
				t.Fatal("queue should be empty after draining / 全部出队后队列应该为空")
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	// Both implementations apply the same cancellation rule; one appends to a separate
	// stack, the other reuses the front of a byte copy.
	// 两种实现使用同一条消除规则：一种另建栈，另一种复用字节副本的前部。
	implementations := []struct {
		name string
		fn   func(string) string
	}{
		{name: "byte slice stack", fn: removeDuplicates},
		{name: "two pointers", fn: removeDuplicatesTwoPointers},
	}

	// These cases cover normal removal, chain reactions, no duplicates,
	// complete removal, and the smallest valid input.
	// 这些用例覆盖普通删除、连锁删除、无重复、全部删除和最小有效输入。
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "example with chain reaction", s: "abbaca", want: "ca"},
		{name: "second chain reaction", s: "azxxzy", want: "ay"},
		{name: "no adjacent duplicates", s: "abc", want: "abc"},
		{name: "all characters removed", s: "aabbcc", want: ""},
		{name: "nested chain reaction", s: "abba", want: ""},
		{name: "single character", s: "a", want: "a"},
		// An empty input must not read the stack top, which has no valid index yet.
		// 空输入不能去读栈顶，因为此时还没有合法下标。
		{name: "empty string", s: "", want: ""},
		// A three-deep nest cancels down to nothing, one pair at a time.
		// 三层嵌套会逐对抵消，最终全部删除。
		{name: "triple nesting", s: "abccba", want: ""},
		// Three identical characters leave exactly one survivor.
		// 三个相同字符最终只剩下一个。
		{name: "odd run length", s: "aaa", want: "a"},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// Compare the actual result with the expected final string.
					// 比较实际结果和预期的最终字符串。
					if got := implementation.fn(tt.s); got != tt.want {
						t.Fatalf("removeDuplicates(%q) = %q, want %q", tt.s, got, tt.want)
					}
				})
			}
		})
	}
}

func TestEvalRPN(t *testing.T) {
	// The forward scan and the backward recursion must produce identical values,
	// which in particular pins down the operand order for - and /.
	// 从前向后扫描与从后向前递归必须给出相同结果，这尤其固定了减法和除法的操作数顺序。
	implementations := []struct {
		name string
		fn   func([]string) int
	}{
		{name: "explicit stack", fn: evalRPN},
		{name: "reverse recursion", fn: evalRPNRecursive},
	}

	// Cover nested operations, operand order, and division truncation toward zero.
	// 覆盖嵌套运算、操作数顺序以及除法向零截断。
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{name: "addition then multiplication", tokens: []string{"2", "1", "+", "3", "*"}, want: 9},
		{name: "division then addition", tokens: []string{"4", "13", "5", "/", "+"}, want: 6},
		{name: "complex expression", tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, want: 22},
		{name: "subtraction operand order", tokens: []string{"8", "3", "-"}, want: 5},
		// Reversing the operands would yield 5 instead of -5, so this pins down the direction.
		// 操作数颠倒会得到 5 而不是 -5，因此这一用例固定了减法方向。
		{name: "subtraction yielding negative", tokens: []string{"3", "8", "-"}, want: -5},
		{name: "negative division truncates toward zero", tokens: []string{"7", "-3", "/"}, want: -2},
		// Go truncates toward zero rather than flooring, so -7/3 is -2, not -3.
		// Go 的整数除法向零截断而不是向下取整，所以 -7/3 是 -2，不是 -3。
		{name: "negative dividend truncates toward zero", tokens: []string{"-7", "3", "/"}, want: -2},
		{name: "multiplication with negative", tokens: []string{"-4", "2", "*"}, want: -8},
		// A lone number is a complete expression and must consume exactly one token.
		// 单个数字本身就是完整表达式，必须恰好消耗一个 token。
		{name: "single number", tokens: []string{"42"}, want: 42},
		{name: "single negative number", tokens: []string{"-13"}, want: -13},
		// A left-leaning chain makes the recursive version nest as deeply as the token count.
		// 左倾的表达式链会让递归版本的嵌套深度接近 token 数量。
		{name: "left leaning chain", tokens: []string{"1", "1", "+", "1", "+", "1", "+"}, want: 4},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.tokens); got != tt.want {
						t.Fatalf("evalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
					}
				})
			}
		})
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{name: "monotonic deque", fn: maxSlidingWindow},
		{name: "brute force", fn: maxSlidingWindowBruteForce},
		{name: "lazy max-heap", fn: maxSlidingWindowHeap},
		{name: "block prefix suffix", fn: maxSlidingWindowBlocks},
	}

	// Cover the standard case, boundary window sizes, decreasing input,
	// and duplicate maximum values.
	// 覆盖标准情况、窗口边界、递减数组和重复最大值。
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "standard example", nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, want: []int{3, 3, 5, 5, 6, 7}},
		{name: "window size one", nums: []int{4, 2, 12}, k: 1, want: []int{4, 2, 12}},
		{name: "window covers entire input", nums: []int{2, 1, 5, 3}, k: 4, want: []int{5}},
		{name: "single element", nums: []int{8}, k: 1, want: []int{8}},
		// The maximum leaves the window, so front eviction must actually happen.
		// 最大值会离开窗口，因此队首过期删除必须真正生效。
		{name: "maximum leaves the window", nums: []int{9, 1, 2, 3}, k: 2, want: []int{9, 2, 3}},
		// Nothing is ever displaced from the back, so the deque and the lazy heap both grow.
		// 队尾元素从不被挤掉，因此单调队列和惰性堆都会不断增长。
		{name: "strictly decreasing", nums: []int{5, 4, 3, 2, 1}, k: 2, want: []int{5, 4, 3, 2}},
		// Every push clears the entire back, so the deque never holds more than one index.
		// 每次入队都会清空整个队尾，因此单调队列始终只保存一个下标。
		{name: "strictly increasing", nums: []int{1, 2, 3, 4, 5}, k: 3, want: []int{3, 4, 5}},
		// Back popping uses <=, so an equal newer value replaces the older one.
		// 队尾弹出条件是 <=，因此数值相等时用更晚的下标替换更早的下标。
		{name: "duplicate maximums", nums: []int{1, 3, 3, 2}, k: 2, want: []int{3, 3, 3}},
		{name: "all equal", nums: []int{7, 7, 7, 7}, k: 2, want: []int{7, 7, 7}},
		// All-negative input rules out a zero-valued neutral initial maximum.
		// 全负数输入可以排除把初值误设为 0 的写法。
		{name: "all negative", nums: []int{-4, -2, -8, -1}, k: 2, want: []int{-2, -2, -1}},
		// A window length that does not divide the input exercises the trailing partial block.
		// 窗口长度不能整除输入长度，用来检验末尾不足一整块的情况。
		{name: "unaligned trailing block", nums: []int{1, 5, 2, 8, 3, 4, 9}, k: 3, want: []int{5, 8, 8, 8, 9}},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// Give each implementation its own copy so a mutation cannot leak between them.
					// 每种实现使用独立副本，避免某种实现误改输入后影响其他实现。
					nums := append([]int(nil), tt.nums...)

					got := implementation.fn(nums, tt.k)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
					}

					// None of the four methods may modify the input array.
					// 四种解法都不允许修改输入数组。
					if !reflect.DeepEqual(nums, tt.nums) {
						t.Fatalf("input modified to %v, want %v unchanged", nums, tt.nums)
					}
				})
			}
		})
	}
}

func TestFrequencyHeap(t *testing.T) {
	// Push frequencies in an unsorted order and verify that heap.Pop
	// returns them from lowest to highest frequency.
	// 按无序频率入堆，并验证 heap.Pop 会按频率从低到高弹出。
	minHeap := &frequencyHeap{}
	heap.Init(minHeap)

	heap.Push(minHeap, [2]int{1, 3})
	heap.Push(minHeap, [2]int{2, 1})
	heap.Push(minHeap, [2]int{3, 2})

	wantFrequencies := []int{1, 2, 3}
	for i, want := range wantFrequencies {
		// The root must already be the minimum before it is popped.
		// 弹出之前，堆顶就必须是频率最低的元素。
		if got := (*minHeap)[0][1]; got != want {
			t.Fatalf("root before pop %d has frequency %d, want %d", i+1, got, want)
		}

		item := heap.Pop(minHeap).([2]int)
		if got := item[1]; got != want {
			t.Fatalf("pop %d returned frequency %d, want %d", i+1, got, want)
		}
	}

	if got := minHeap.Len(); got != 0 {
		t.Fatalf("Len() after all pops = %d, want 0 / 全部弹出后堆应该为空", got)
	}
}

func TestWindowMaxHeap(t *testing.T) {
	// 239's heap must order by descending value, the opposite of 347's min-heap.
	// Pushing out of order confirms that Less points the right way for a max-heap.
	// 239 的堆按数值降序排列，与 347 的小顶堆方向相反。
	// 乱序入堆可以确认 Less 的方向确实构成大顶堆。
	maxHeap := &windowMaxHeap{}
	heap.Init(maxHeap)

	heap.Push(maxHeap, [2]int{3, 0})
	heap.Push(maxHeap, [2]int{9, 1})
	heap.Push(maxHeap, [2]int{5, 2})

	// Each pair carries its original index, which lazy deletion relies on for expiry checks.
	// 每一项都携带原始下标，惰性删除正是依靠它判断是否过期。
	wantPairs := [][2]int{{9, 1}, {5, 2}, {3, 0}}
	for i, want := range wantPairs {
		if got := (*maxHeap)[0]; got != want {
			t.Fatalf("root before pop %d = %v, want %v", i+1, got, want)
		}

		if got := heap.Pop(maxHeap).([2]int); got != want {
			t.Fatalf("pop %d returned %v, want %v", i+1, got, want)
		}
	}

	if got := maxHeap.Len(); got != 0 {
		t.Fatalf("Len() after all pops = %d, want 0 / 全部弹出后堆应该为空", got)
	}
}

// repeatedValues builds an input in which value i+1 occurs counts[i] times.
// repeatedValues 构造输入：第 i 个值恰好出现 counts[i] 次。
//
// Distinct counts give every value a distinct rank, so the expected answer set
// stays uniquely determined no matter where the k boundary falls.
// 各值出现次数互不相同，因此每个值的名次唯一，无论 k 取在哪里，期望答案集合都唯一确定。
func repeatedValues(counts ...int) []int {
	nums := []int{}
	for i, count := range counts {
		for range count {
			nums = append(nums, i+1)
		}
	}
	return nums
}

func TestTopKFrequent(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{name: "min-heap of size k", fn: topKFrequent},
		{name: "full sort", fn: topKFrequentSorted},
		{name: "bucket sort", fn: topKFrequentBucket},
		{name: "quickselect", fn: topKFrequentQuickselect},
	}

	// Result order is not part of the contract, so each result and expected
	// slice is sorted before comparison.
	// 题目不要求结果顺序，因此比较前分别对实际结果和预期结果排序。
	//
	// Every case has a uniquely determined answer set: no tied frequency straddles the
	// k boundary, which would make several different return values equally correct
	// and leave nothing for the test to compare against.
	// 每个用例的答案集合都唯一确定：没有并列频次跨越第 k 名边界，
	// 否则多个不同的返回值都算正确，测试就失去了可比较的期望值。
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "standard example", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "single element", nums: []int{1}, k: 1, want: []int{1}},
		// One distinct value repeated: the bucket index reaches len(nums), the highest slot.
		// 只有一个不同的值且重复出现：桶下标会取到 len(nums)，也就是最高的那一格。
		{name: "all identical", nums: []int{9, 9, 9}, k: 1, want: []int{9}},
		{name: "negative numbers", nums: []int{-1, -1, -1, -2, -2, -3}, k: 2, want: []int{-1, -2}},
		{name: "one most frequent value", nums: []int{5, 5, 5, 1, 2, 3}, k: 1, want: []int{5}},
		{name: "two equal top frequencies", nums: []int{4, 1, -1, 2, -1, 2, 3}, k: 2, want: []int{-1, 2}},
		{name: "k equals distinct count", nums: []int{4, 4, 5, 6}, k: 3, want: []int{4, 5, 6}},
		// k covers the whole array, so every distinct value must be returned.
		// k 覆盖整个数组，因此所有不同的值都必须返回。
		{name: "k equals whole array", nums: []int{7, 8, 9}, k: 3, want: []int{7, 8, 9}},
		{name: "zero and duplicates", nums: []int{0, 0, 0, 4, 4, 6}, k: 2, want: []int{0, 4}},
		// Eight distinct frequency levels force quickselect through several partition rounds
		// and make the bucket scan walk down from len(nums) to the highest occupied bucket.
		// 八个互不相同的频次会让快速选择经过多轮划分，
		// 也让桶扫描从 len(nums) 一直向下走到真正有元素的最高桶。
		{name: "many frequency levels middle k", nums: repeatedValues(8, 7, 6, 5, 4, 3, 2, 1), k: 4, want: []int{1, 2, 3, 4}},
		// A high k puts the target rank near the least frequent end, so quickselect must
		// advance past the equal band into the lower-frequency side.
		// k 较大时目标名次靠近最低频端，快速选择必须越过等频段进入低频一侧。
		{name: "many frequency levels high k", nums: repeatedValues(8, 7, 6, 5, 4, 3, 2, 1), k: 7, want: []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// Quickselect reorders its own working slice, so verify the input survives.
					// 快速选择会重排自己的工作切片，因此要确认输入本身没有被改动。
					nums := append([]int(nil), tt.nums...)

					got := implementation.fn(nums, tt.k)
					sort.Ints(got)

					// Copy want so sorting does not modify the table entry.
					// 复制 want，避免排序修改表格中的原始测试数据。
					want := append([]int(nil), tt.want...)
					sort.Ints(want)

					if !reflect.DeepEqual(got, want) {
						t.Fatalf("topKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, want)
					}

					if !reflect.DeepEqual(nums, tt.nums) {
						t.Fatalf("input modified to %v, want %v unchanged", nums, tt.nums)
					}
				})
			}
		})
	}
}
