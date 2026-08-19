package _3_hash

import (
	"reflect"
	"sort"
	"testing"
)

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
