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
//
// 步骤与要点 / Steps and notes:
//  1. Even length with only openings: only the final nonempty-stack check can reject this.
//     长度为偶数但全是左括号：只能靠最后的“栈非空”检查判错。
//  2. Even length with only closings: the empty-stack check must reject the first character.
//     长度为偶数但全是右括号：必须在第一个字符处靠“栈为空”检查判错。
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
		{name: "only opening brackets", s: "((", want: false},
		{name: "only closing brackets", s: "))", want: false},
		{name: "closing before opening", s: "}{", want: false},
		{name: "empty string", s: "", want: true},
	}
}

// 步骤与要点 / Steps and notes:
//
//  1. Cover valid nesting and the three invalid cases:
//     extra opening brackets, extra closing brackets, and mismatched brackets.
//     覆盖有效嵌套以及三种无效情况：左括号多余、右括号多余和括号不匹配。
//
//     The two methods translate between bracket types at opposite moments,
//     so they must still agree on every verdict.
//     两种解法在不同时刻完成括号类型的转换，但对每个判定结果必须完全一致。
func TestIsValid(t *testing.T) {
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

// stackUnderTest is the common surface of the three 225 implementations.
// stackUnderTest 是 225 题三种实现共有的对外接口。
//
// All satisfy it through pointer receivers, so the harness always stores a pointer.
// 三种实现都通过指针接收者满足它，因此测试统一持有指针。
type stackUnderTest interface {
	Push(x int)
	Pop() int
	Top() int
	Empty() bool
}

// 步骤与要点 / Steps and notes:
//  1. External behaviour must be identical; only the placement of the O(n) rotation differs.
//     对外行为必须完全一致，区别只在于 O(n) 旋转发生在入栈还是出栈。
//  2. MyStackLazy has no constructor and must work straight from its zero value.
//     MyStackLazy 没有构造函数，必须直接从零值开始就能工作。
//  3. MyStackTwoQueues.Empty reads only main, so a leaked element in backup would go unnoticed;
//     interleaved operations catch a broken role swap between main and backup.
//     MyStackTwoQueues.Empty 只看 main，如果元素残留在 backup 里就不会被发现；
//     交错操作用于暴露 main 和 backup 的角色交换错误。
//  4. Verify LIFO order and state changes after interleaved operations.
//     验证后进先出顺序，以及交错操作后的状态变化。
//  5. Top must not consume the element it reports; the lazy version rebuilds the queue.
//     Top 不能消耗它返回的元素；惰性版本会重新排列队列，必须完整还原。
//  6. Pushing while elements remain must not disturb the existing order.
//     栈内仍有元素时继续入栈，不能打乱已有元素的顺序。
//  7. A single element exercises the degenerate rotation of zero items.
//     单个元素用于检验“旋转零次”这种退化情况。
//  8. A longer run stresses the rotation invariant beyond the hand-checked cases.
//     更长的操作序列用来检验旋转不变量，而不只是手算的小例子。
func TestMyStack(t *testing.T) {
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
			name: "rotate on pop",
			make: func() stackUnderTest {
				return &MyStackLazy{}
			},
		},
		{
			name: "two queues",
			make: func() stackUnderTest {
				return &MyStackTwoQueues{}
			},
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			stack := implementation.make()

			if !stack.Empty() {
				t.Fatal("a new stack should be empty / 新栈应该为空")
			}

			stack.Push(1)
			stack.Push(2)
			if got := stack.Top(); got != 2 {
				t.Fatalf("Top() = %d, want 2", got)
			}

			if got := stack.Top(); got != 2 {
				t.Fatalf("second Top() = %d, want 2", got)
			}

			if got := stack.Pop(); got != 2 {
				t.Fatalf("Pop() = %d, want 2", got)
			}
			if got := stack.Top(); got != 1 {
				t.Fatalf("Top() after Pop() = %d, want 1", got)
			}

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

// TestMyQueue
//
// 检查两种队列的 FIFO 顺序，以及 Peek 不移除元素。
// Check FIFO behavior and non-removing Peek for both queue implementations.
// 交错 Push/Pop 验证惰性版只在输出栈为空时转移，避免新元素插到尚未出队的旧元素前面。
// Interleaved Push/Pop checks that the lazy queue transfers only when its output stack is empty, preserving older entries first.
func TestMyQueue(t *testing.T) {
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
			name: "eager reordering",
			make: func() queueUnderTest {
				return &MyQueueEager{}
			},
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			queue := implementation.make()

			if !queue.Empty() {
				t.Fatal("a new queue should be empty / 新队列应该为空")
			}

			queue.Push(1)
			queue.Push(2)
			if got := queue.Peek(); got != 1 {
				t.Fatalf("Peek() = %d, want 1", got)
			}

			if got := queue.Peek(); got != 1 {
				t.Fatalf("second Peek() = %d, want 1", got)
			}

			if got := queue.Pop(); got != 1 {
				t.Fatalf("Pop() = %d, want 1", got)
			}

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

// 步骤与要点 / Steps and notes:
//  1. Both implementations apply the same cancellation rule; one appends to a separate
//     stack, the other reuses the front of a byte copy.
//     两种实现使用同一条消除规则：一种另建栈，另一种复用字节副本的前部。
//  2. These cases cover normal removal, chain reactions, no duplicates,
//     complete removal, and the smallest valid input.
//     这些用例覆盖普通删除、连锁删除、无重复、全部删除和最小有效输入。
//  3. An empty input must not read the stack top, which has no valid index yet.
//     空输入不能去读栈顶，因为此时还没有合法下标。
//  4. A three-deep nest cancels down to nothing, one pair at a time.
//     三层嵌套会逐对抵消，最终全部删除。
//  5. Three identical characters leave exactly one survivor.
//     三个相同字符最终只剩下一个。
//  6. Compare the actual result with the expected final string.
//     比较实际结果和预期的最终字符串。
func TestRemoveDuplicates(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(string) string
	}{
		{name: "byte slice stack", fn: removeDuplicates},
		{name: "two pointers", fn: removeDuplicatesTwoPointers},
	}

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
		{name: "empty string", s: "", want: ""},
		{name: "triple nesting", s: "abccba", want: ""},
		{name: "odd run length", s: "aaa", want: "a"},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.s); got != tt.want {
						t.Fatalf("removeDuplicates(%q) = %q, want %q", tt.s, got, tt.want)
					}
				})
			}
		})
	}
}

// 步骤与要点 / Steps and notes:
//  1. The forward scan and the backward recursion must produce identical values,
//     which in particular pins down the operand order for - and /.
//     从前向后扫描与从后向前递归必须给出相同结果，这尤其固定了减法和除法的操作数顺序。
//  2. Cover nested operations, operand order, and division truncation toward zero.
//     覆盖嵌套运算、操作数顺序以及除法向零截断。
//  3. Reversing the operands would yield 5 instead of -5, so this pins down the direction.
//     操作数颠倒会得到 5 而不是 -5，因此这一用例固定了减法方向。
//  4. Go truncates toward zero rather than flooring, so -7/3 is -2, not -3.
//     Go 的整数除法向零截断而不是向下取整，所以 -7/3 是 -2，不是 -3。
//  5. A lone number is a complete expression and must consume exactly one token.
//     单个数字本身就是完整表达式，必须恰好消耗一个 token。
//  6. A left-leaning chain makes the recursive version nest as deeply as the token count.
//     左倾的表达式链会让递归版本的嵌套深度接近 token 数量。
func TestEvalRPN(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([]string) int
	}{
		{name: "explicit stack", fn: evalRPN},
		{name: "reverse recursion", fn: evalRPNRecursive},
	}

	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{name: "addition then multiplication", tokens: []string{"2", "1", "+", "3", "*"}, want: 9},
		{name: "division then addition", tokens: []string{"4", "13", "5", "/", "+"}, want: 6},
		{name: "complex expression", tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, want: 22},
		{name: "subtraction operand order", tokens: []string{"8", "3", "-"}, want: 5},
		{name: "subtraction yielding negative", tokens: []string{"3", "8", "-"}, want: -5},
		{name: "negative division truncates toward zero", tokens: []string{"7", "-3", "/"}, want: -2},
		{name: "negative dividend truncates toward zero", tokens: []string{"-7", "3", "/"}, want: -2},
		{name: "multiplication with negative", tokens: []string{"-4", "2", "*"}, want: -8},
		{name: "single number", tokens: []string{"42"}, want: 42},
		{name: "single negative number", tokens: []string{"-13"}, want: -13},
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

// TestMaxSlidingWindow
//
// 覆盖最大值过期、重复值、单调序列和尾部不足整块；同时检查输入未被修改。
// Cover expired maxima, duplicates, monotone inputs and a partial trailing block, also checking input preservation.
// 递增输入会使惰性堆积累非堆顶的过期小值；单调队列及时淘汰候选，最多保存 k 个下标。
// Increasing input lets a lazy heap retain expired smaller entries below its root; the deque removes dominated candidates and keeps at most k indices.
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
		{name: "maximum leaves the window", nums: []int{9, 1, 2, 3}, k: 2, want: []int{9, 2, 3}},
		{name: "strictly decreasing", nums: []int{5, 4, 3, 2, 1}, k: 2, want: []int{5, 4, 3, 2}},
		{name: "strictly increasing", nums: []int{1, 2, 3, 4, 5}, k: 3, want: []int{3, 4, 5}},
		{name: "duplicate maximums", nums: []int{1, 3, 3, 2}, k: 2, want: []int{3, 3, 3}},
		{name: "all equal", nums: []int{7, 7, 7, 7}, k: 2, want: []int{7, 7, 7}},
		{name: "all negative", nums: []int{-4, -2, -8, -1}, k: 2, want: []int{-2, -2, -1}},
		{name: "unaligned trailing block", nums: []int{1, 5, 2, 8, 3, 4, 9}, k: 3, want: []int{5, 8, 8, 8, 9}},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					nums := append([]int(nil), tt.nums...)

					got := implementation.fn(nums, tt.k)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
					}

					if !reflect.DeepEqual(nums, tt.nums) {
						t.Fatalf("input modified to %v, want %v unchanged", nums, tt.nums)
					}
				})
			}
		})
	}
}

// 步骤与要点 / Steps and notes:
//  1. Push frequencies in an unsorted order and verify that heap.Pop
//     returns them from lowest to highest frequency.
//     按无序频率入堆，并验证 heap.Pop 会按频率从低到高弹出。
//  2. The root must already be the minimum before it is popped.
//     弹出之前，堆顶就必须是频率最低的元素。
func TestFrequencyHeap(t *testing.T) {
	minHeap := &frequencyHeap{}
	heap.Init(minHeap)

	heap.Push(minHeap, [2]int{1, 3})
	heap.Push(minHeap, [2]int{2, 1})
	heap.Push(minHeap, [2]int{3, 2})

	wantFrequencies := []int{1, 2, 3}
	for i, want := range wantFrequencies {
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

// 步骤与要点 / Steps and notes:
//  1. 239's heap must order by descending value, the opposite of 347's min-heap.
//     Pushing out of order confirms that Less points the right way for a max-heap.
//     239 的堆按数值降序排列，与 347 的小顶堆方向相反。
//     乱序入堆可以确认 Less 的方向确实构成大顶堆。
//  2. Each pair carries its original index, which lazy deletion relies on for expiry checks.
//     每一项都携带原始下标，惰性删除正是依靠它判断是否过期。
func TestWindowMaxHeap(t *testing.T) {
	maxHeap := &windowMaxHeap{}
	heap.Init(maxHeap)

	heap.Push(maxHeap, [2]int{3, 0})
	heap.Push(maxHeap, [2]int{9, 1})
	heap.Push(maxHeap, [2]int{5, 2})

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

// repeatedValues
//
// 生成测试数据：数值 i+1 重复 counts[i] 次，便于直接指定每个值的频率。
// Generate value i+1 exactly counts[i] times so tests can specify frequencies directly.
func repeatedValues(counts ...int) []int {
	nums := []int{}
	for i, count := range counts {
		for range count {
			nums = append(nums, i+1)
		}
	}
	return nums
}

// TestTopKFrequent
//
// 结果顺序不限，因此排序后比较；用例选取确定的前 k 项集合，并检查输入未修改。
// Sort results before comparison because order is unspecified; cases use an unambiguous top-k set and check input preservation.
// 多频率层级覆盖不同 k 的选择，但随机枢轴仍可能一次命中，不能保证执行多轮快速选择。
// Several frequency levels exercise different k values, but a random pivot may hit immediately and does not guarantee multiple quickselect rounds.
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

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "standard example", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "single element", nums: []int{1}, k: 1, want: []int{1}},
		{name: "all identical", nums: []int{9, 9, 9}, k: 1, want: []int{9}},
		{name: "negative numbers", nums: []int{-1, -1, -1, -2, -2, -3}, k: 2, want: []int{-1, -2}},
		{name: "one most frequent value", nums: []int{5, 5, 5, 1, 2, 3}, k: 1, want: []int{5}},
		{name: "two equal top frequencies", nums: []int{4, 1, -1, 2, -1, 2, 3}, k: 2, want: []int{-1, 2}},
		{name: "k equals distinct count", nums: []int{4, 4, 5, 6}, k: 3, want: []int{4, 5, 6}},
		{name: "k equals whole array", nums: []int{7, 8, 9}, k: 3, want: []int{7, 8, 9}},
		{name: "zero and duplicates", nums: []int{0, 0, 0, 4, 4, 6}, k: 2, want: []int{0, 4}},
		{name: "many frequency levels middle k", nums: repeatedValues(8, 7, 6, 5, 4, 3, 2, 1), k: 4, want: []int{1, 2, 3, 4}},
		{name: "many frequency levels high k", nums: repeatedValues(8, 7, 6, 5, 4, 3, 2, 1), k: 7, want: []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					nums := append([]int(nil), tt.nums...)

					got := implementation.fn(nums, tt.k)
					sort.Ints(got)

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
