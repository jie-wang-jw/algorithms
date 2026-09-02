package _5_stack_queue

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	// Cover valid nesting and the three invalid cases:
	// extra opening brackets, extra closing brackets, and mismatched brackets.
	// 覆盖有效嵌套以及三种无效情况：左括号多余、右括号多余和括号不匹配。
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "single pair", s: "()", want: true},
		{name: "different adjacent pairs", s: "()[]{}", want: true},
		{name: "nested pairs", s: "{[()]}", want: true},
		{name: "mismatched type", s: "(]", want: false},
		{name: "wrong closing order", s: "([)]", want: false},
		{name: "extra opening bracket", s: "(()", want: false},
		{name: "extra closing bracket", s: "())", want: false},
		{name: "empty string", s: "", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Fatalf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestMyStack(t *testing.T) {
	// Verify LIFO order and state changes after interleaved operations.
	// 验证后进先出顺序，以及交错操作后的状态变化。
	stack := StackConstructor()

	if !stack.Empty() {
		t.Fatal("a new stack should be empty / 新栈应该为空")
	}

	stack.Push(1)
	stack.Push(2)
	if got := stack.Top(); got != 2 {
		t.Fatalf("Top() = %d, want 2", got)
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
}

func TestMyQueue(t *testing.T) {
	// Verify FIFO order, including pushes performed while outStack still has data.
	// 验证先进先出顺序，包括 outStack 尚有元素时继续入队的情况。
	queue := Constructor()

	if !queue.Empty() {
		t.Fatal("a new queue should be empty / 新队列应该为空")
	}

	queue.Push(1)
	queue.Push(2)
	if got := queue.Peek(); got != 1 {
		t.Fatalf("Peek() = %d, want 1", got)
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
}

func TestRemoveDuplicates(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Compare the actual result with the expected final string.
			// 比较实际结果和预期的最终字符串。
			if got := removeDuplicates(tt.s); got != tt.want {
				t.Fatalf("removeDuplicates(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestMaxSlidingWindow(t *testing.T) {
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
		{name: "strictly decreasing", nums: []int{5, 4, 3, 2, 1}, k: 2, want: []int{5, 4, 3, 2}},
		{name: "duplicate maximums", nums: []int{1, 3, 3, 2}, k: 2, want: []int{3, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
