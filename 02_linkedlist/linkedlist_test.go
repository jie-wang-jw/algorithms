package _2_linkedlist

import "testing"

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func requireList(t *testing.T, head *ListNode, want []int) {
	t.Helper()
	got := listToSlice(head)
	if !equalSlice(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func buildCycleList(vals []int, pos int) (*ListNode, *ListNode) {
	if len(vals) == 0 {
		return nil, nil
	}

	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
		return nodes[0], nodes[pos]
	}

	return nodes[0], nil
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty", in: nil, want: nil},
		{name: "single", in: []int{1}, want: []int{1}},
		{name: "two nodes", in: []int{1, 2}, want: []int{2, 1}},
		{name: "three nodes", in: []int{1, 2, 3}, want: []int{3, 2, 1}},
		{name: "five nodes", in: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "duplicate values", in: []int{7, 7, 8}, want: []int{8, 7, 7}},
	}

	reverseFuncs := []struct {
		name string
		fn   func(*ListNode) *ListNode
	}{
		{"iterative two pointers", reverseList},
		{"back to front recursion", reverseList2},
		{"front to back recursion", reverseListFromFront},
	}

	for _, rf := range reverseFuncs {
		t.Run(rf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					head := buildList(tt.in)
					got := rf.fn(head)
					requireList(t, got, tt.want)
				})
			}
		})
	}
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty", in: nil, want: nil},
		{name: "single", in: []int{1}, want: []int{1}},
		{name: "two nodes", in: []int{1, 2}, want: []int{2, 1}},
		{name: "even length", in: []int{1, 2, 3, 4}, want: []int{2, 1, 4, 3}},
		{name: "odd length", in: []int{1, 2, 3, 4, 5}, want: []int{2, 1, 4, 3, 5}},
		{name: "six nodes", in: []int{1, 2, 3, 4, 5, 6}, want: []int{2, 1, 4, 3, 6, 5}},
	}

	swapFuncs := []struct {
		name string
		fn   func(*ListNode) *ListNode
	}{
		{"dummy head iteration", swapPairs},
		{"recursion", swapPairsRecursive},
	}

	for _, sf := range swapFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					head := buildList(tt.in)
					got := sf.fn(head)
					requireList(t, got, tt.want)
				})
			}
		})
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		n    int
		want []int
	}{
		{name: "remove middle", in: []int{1, 2, 3, 4, 5}, n: 2, want: []int{1, 2, 3, 5}},
		{name: "remove head of five", in: []int{1, 2, 3, 4, 5}, n: 5, want: []int{2, 3, 4, 5}},
		{name: "remove tail of five", in: []int{1, 2, 3, 4, 5}, n: 1, want: []int{1, 2, 3, 4}},
		{name: "remove head", in: []int{1, 2}, n: 2, want: []int{2}},
		{name: "remove tail", in: []int{1, 2}, n: 1, want: []int{1}},
		{name: "remove only node", in: []int{1}, n: 1, want: nil},
	}

	removeFuncs := []struct {
		name string
		fn   func(*ListNode, int) *ListNode
	}{
		{"fast slow pointers", removeNthFromEnd},
		{"length then two passes", removeNthFromEndLength},
		{"stack", removeNthFromEndStack},
	}

	for _, rf := range removeFuncs {
		t.Run(rf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					head := buildList(tt.in)
					got := rf.fn(head, tt.n)
					requireList(t, got, tt.want)
				})
			}
		})
	}
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want int
	}{
		{name: "cycle starts at index 1", vals: []int{3, 2, 0, -4}, pos: 1, want: 2},
		{name: "cycle starts at head", vals: []int{1, 2}, pos: 0, want: 1},
		{name: "cycle starts at tail", vals: []int{1, 2, 3}, pos: 2, want: 3},
		{name: "single node cycle", vals: []int{1}, pos: 0, want: 1},
		{name: "no cycle", vals: []int{1, 2, 3}, pos: -1, want: 0},
		{name: "single node no cycle", vals: []int{1}, pos: -1, want: 0},
		{name: "empty", vals: nil, pos: -1, want: 0},
	}

	detectFuncs := []struct {
		name string
		fn   func(*ListNode) *ListNode
	}{
		{"floyd fast slow pointers", detectCycle},
		{"visited node set", detectCycleHash},
	}

	for _, df := range detectFuncs {
		t.Run(df.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					head, entry := buildCycleList(tt.vals, tt.pos)
					got := df.fn(head)

					if tt.pos == -1 {
						if got != nil {
							t.Fatalf("got node value %d, want nil", got.Val)
						}
						return
					}

					if got != entry {
						if got == nil {
							t.Fatalf("got nil, want entry value %d", tt.want)
						}
						t.Fatalf("got node value %d, want entry value %d", got.Val, tt.want)
					}
				})
			}
		})
	}
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		val  int
		want []int
	}{
		{name: "remove tail duplicates", in: []int{1, 2, 6, 3, 4, 5, 6}, val: 6, want: []int{1, 2, 3, 4, 5}},
		{name: "empty", in: nil, val: 1, want: nil},
		{name: "all removed", in: []int{7, 7, 7, 7}, val: 7, want: nil},
		{name: "remove head", in: []int{1, 2, 3}, val: 1, want: []int{2, 3}},
		{name: "none removed", in: []int{1, 2, 3}, val: 4, want: []int{1, 2, 3}},
		{name: "single kept", in: []int{1}, val: 2, want: []int{1}},
		{name: "single removed", in: []int{1}, val: 1, want: nil},
		{name: "alternating", in: []int{1, 2, 1, 2, 1}, val: 1, want: []int{2, 2}},
	}

	removeFuncs := []struct {
		name string
		fn   func(*ListNode, int) *ListNode
	}{
		{"original list", removeElementsDirect},
		{"dummy head", removeElements},
		{"recursion", removeElementsRecursive},
	}

	for _, rf := range removeFuncs {
		t.Run(rf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := rf.fn(buildList(tt.in), tt.val)
					requireList(t, got, tt.want)
				})
			}
		})
	}
}

func TestMyLinkedList(t *testing.T) {
	list := Constructor()
	if got := list.Get(0); got != -1 {
		t.Fatalf("empty Get(0) = %d, want -1", got)
	}

	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	requireList(t, list.dummy.Next, []int{1, 2, 3})
	if got := list.Get(1); got != 2 {
		t.Fatalf("Get(1) = %d, want 2", got)
	}

	list.DeleteAtIndex(1)
	requireList(t, list.dummy.Next, []int{1, 3})
	if got := list.Get(1); got != 3 {
		t.Fatalf("Get(1) = %d, want 3", got)
	}

	list.AddAtIndex(0, 0)
	requireList(t, list.dummy.Next, []int{0, 1, 3})
	list.AddAtIndex(3, 4)
	requireList(t, list.dummy.Next, []int{0, 1, 3, 4})
	list.AddAtIndex(10, 9)
	requireList(t, list.dummy.Next, []int{0, 1, 3, 4})
	list.AddAtIndex(-1, -1)
	requireList(t, list.dummy.Next, []int{-1, 0, 1, 3, 4})

	list.DeleteAtIndex(0)
	requireList(t, list.dummy.Next, []int{0, 1, 3, 4})
	list.DeleteAtIndex(3)
	requireList(t, list.dummy.Next, []int{0, 1, 3})
	list.DeleteAtIndex(100)
	requireList(t, list.dummy.Next, []int{0, 1, 3})
	if got := list.Get(-1); got != -1 {
		t.Fatalf("Get(-1) = %d, want -1", got)
	}
}

func TestGetIntersectionNode(t *testing.T) {
	intersectFuncs := []struct {
		name string
		fn   func(*ListNode, *ListNode) *ListNode
	}{
		{"align by length", getIntersectionNode},
		{"switch heads", getIntersectionNodeTwoPointers},
	}

	for _, inf := range intersectFuncs {
		t.Run(inf.name, func(t *testing.T) {
			t.Run("intersect in the middle", func(t *testing.T) {
				shared := buildList([]int{8, 4, 5})
				headA := buildList([]int{4, 1})
				headB := buildList([]int{5, 6, 1})
				tail(headA).Next = shared
				tail(headB).Next = shared
				if got := inf.fn(headA, headB); got != shared {
					t.Fatalf("want shared node 8")
				}
			})
			t.Run("intersect at head", func(t *testing.T) {
				headA := buildList([]int{1, 2, 3})
				if got := inf.fn(headA, headA); got != headA {
					t.Fatalf("same list should intersect at head")
				}
			})
			t.Run("no intersection", func(t *testing.T) {
				headA := buildList([]int{2, 6, 4})
				headB := buildList([]int{1, 5})
				if got := inf.fn(headA, headB); got != nil {
					t.Fatalf("got node %d, want nil", got.Val)
				}
			})
			t.Run("both empty", func(t *testing.T) {
				if got := inf.fn(nil, nil); got != nil {
					t.Fatalf("got %v, want nil", got)
				}
			})
			t.Run("one empty", func(t *testing.T) {
				headA := buildList([]int{1})
				if got := inf.fn(headA, nil); got != nil {
					t.Fatalf("got %v, want nil", got)
				}
			})
		})
	}
}

func tail(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}
