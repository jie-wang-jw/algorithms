package _3_hash

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeIntGroups(groups [][]int) [][]int {
	for _, group := range groups {
		sort.Ints(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		for k := 0; k < len(groups[i]) && k < len(groups[j]); k++ {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return len(groups[i]) < len(groups[j])
	})
	return groups
}
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "example", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "middle pair", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "same value different indices", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{name: "no solution", nums: []int{1, 2, 3}, target: 100, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "one is happy", n: 1, want: true},
		{name: "example happy number", n: 19, want: true},
		{name: "cycle number", n: 2, want: false},
		{name: "another happy number", n: 7, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Fatalf("isHappy(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 19, want: 82},
		{n: 82, want: 68},
		{n: 100, want: 1},
	}

	for _, tt := range tests {
		if got := getNext(tt.n); got != tt.want {
			t.Fatalf("getNext(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "same letters different order", s: "anagram", t: "nagaram", want: true},
		{name: "different counts", s: "rat", t: "car", want: false},
		{name: "different lengths", s: "ab", t: "a", want: false},
		{name: "same repeated letters", s: "aabb", t: "baba", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name+" isAnagram1", func(t *testing.T) {
			if got := isAnagram1(tt.s, tt.t); got != tt.want {
				t.Fatalf("isAnagram1(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})

		t.Run(tt.name+" isAnagram2", func(t *testing.T) {
			if got := isAnagram2(tt.s, tt.t); got != tt.want {
				t.Fatalf("isAnagram2(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{name: "deduplicate result", nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, want: []int{2}},
		{name: "multiple common values", nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{4, 9}},
		{name: "no intersection", nums1: []int{1, 3, 5}, nums2: []int{2, 4, 6}, want: []int{}},
		{name: "negative numbers", nums1: []int{-1, 0, 1}, nums2: []int{1, -1, -1}, want: []int{-1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.nums1, tt.nums2)
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "example with duplicates", nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "all zeros deduplicated", nums: []int{0, 0, 0, 0}, want: [][]int{{0, 0, 0}}},
		{name: "no triplet", nums: []int{1, 2, -2, -1}, want: [][]int{}},
		{name: "less than three numbers", nums: []int{0, 1}, want: [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeIntGroups(threeSum(append([]int(nil), tt.nums...)))
			want := normalizeIntGroups(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("threeSum(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{name: "example", nums: []int{1, 0, -1, 0, -2, 2}, target: 0, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{name: "all same values deduplicated", nums: []int{2, 2, 2, 2, 2}, target: 8, want: [][]int{{2, 2, 2, 2}}},
		{name: "negative target", nums: []int{-3, -1, 0, 2, 4, 5}, target: 2, want: [][]int{{-3, -1, 2, 4}}},
		{name: "no quadruplet", nums: []int{1, 2, 3}, target: 6, want: [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeIntGroups(fourSum(append([]int(nil), tt.nums...), tt.target))
			want := normalizeIntGroups(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("fourSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, want)
			}
		})
	}
}

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		c    []int
		d    []int
		want int
	}{
		{name: "example", a: []int{1, 2}, b: []int{-2, -1}, c: []int{-1, 2}, d: []int{0, 2}, want: 2},
		{name: "all zeros", a: []int{0, 0}, b: []int{0}, c: []int{0}, d: []int{0, 0}, want: 4},
		{name: "no tuples", a: []int{1}, b: []int{1}, c: []int{1}, d: []int{1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.a, tt.b, tt.c, tt.d); got != tt.want {
				t.Fatalf("fourSumCount(%v, %v, %v, %v) = %d, want %d", tt.a, tt.b, tt.c, tt.d, got, tt.want)
			}
		})
	}
}
