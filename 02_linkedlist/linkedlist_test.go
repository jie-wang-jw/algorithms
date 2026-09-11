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
