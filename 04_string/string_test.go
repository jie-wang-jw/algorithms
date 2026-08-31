package _4_string

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	// Both implementations should remove extra spaces and reverse the word order.
	// 两种实现都应该删除多余空格，并反转单词顺序。
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "basic sentence", s: "the sky is blue", want: "blue is sky the"},
		{name: "leading and trailing spaces", s: "  hello world  ", want: "world hello"},
		{name: "multiple spaces between words", s: "a good   example", want: "example good a"},
		{name: "single word", s: "algorithm", want: "algorithm"},
		{name: "only spaces", s: "     ", want: ""},
	}

	reverseWordsFuncs := []struct {
		name string
		fn   func(string) string
	}{
		{name: "strings.Fields solution", fn: reverseWords},
		{name: "manual in-place solution", fn: reverseWords2},
	}

	for _, rf := range reverseWordsFuncs {
		t.Run(rf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := rf.fn(tt.s); got != tt.want {
						t.Fatalf("%s(%q) = %q, want %q", rf.name, tt.s, got, tt.want)
					}
				})
			}
		})
	}
}

func TestStrStr(t *testing.T) {
	// Run the same cases against the brute-force and KMP implementations.
	// 使用同一组用例验证暴力匹配和 KMP 两种实现。
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{name: "match at beginning", haystack: "sadbutsad", needle: "sad", want: 0},
		{name: "match in middle", haystack: "hello", needle: "ll", want: 2},
		{name: "overlapping characters", haystack: "mississippi", needle: "issip", want: 4},
		{name: "not found", haystack: "leetcode", needle: "leeto", want: -1},
		{name: "needle longer than haystack", haystack: "a", needle: "aa", want: -1},
		{name: "empty needle", haystack: "abc", needle: "", want: 0},
		{name: "both empty", haystack: "", needle: "", want: 0},
	}

	strStrFuncs := []struct {
		name string
		fn   func(string, string) int
	}{
		{name: "brute-force solution", fn: strStr},
		{name: "KMP solution", fn: strStr2},
	}

	for _, sf := range strStrFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := sf.fn(tt.haystack, tt.needle); got != tt.want {
						t.Fatalf("%s(%q, %q) = %d, want %d", sf.name, tt.haystack, tt.needle, got, tt.want)
					}
				})
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	// next[i] stores the longest equal proper prefix and suffix length for s[0:i+1].
	// next[i] 保存 s[0:i+1] 的最长相等真前缀和真后缀长度。
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{name: "alternating pattern", s: "abab", want: []int{0, 0, 1, 2}},
		{name: "fallback after partial match", s: "aabaaf", want: []int{0, 1, 0, 1, 2, 0}},
		{name: "long prefix broken at end", s: "aaaaab", want: []int{0, 1, 2, 3, 4, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			next := make([]int, len(tt.s))
			getNext(next, tt.s)
			if !reflect.DeepEqual(next, tt.want) {
				t.Fatalf("getNext(%q) = %v, want %v", tt.s, next, tt.want)
			}
		})
	}
}

func TestRepeatedSubstringPattern(t *testing.T) {
	// All three solutions should agree on repeated and non-repeated strings.
	// 枚举、双倍字符串和 KMP 三种实现应该得到相同结果。
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "two copies", s: "abab", want: true},
		{name: "three copies", s: "abcabcabc", want: true},
		{name: "same character repeated", s: "aaaa", want: true},
		{name: "longer repeated block", s: "abaababaab", want: true},
		{name: "equal prefix suffix but incomplete blocks", s: "aba", want: false},
		{name: "no repeated pattern", s: "abac", want: false},
		{name: "single character", s: "a", want: false},
	}

	repeatedFuncs := []struct {
		name string
		fn   func(string) bool
	}{
		{name: "enumeration solution", fn: repeatedSubstringPattern},
		{name: "doubled-string solution", fn: repeatedSubstringPattern2},
		{name: "KMP solution", fn: repeatedSubstringPattern3},
	}

	for _, rf := range repeatedFuncs {
		t.Run(rf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := rf.fn(tt.s); got != tt.want {
						t.Fatalf("%s(%q) = %v, want %v", rf.name, tt.s, got, tt.want)
					}
				})
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{name: "odd length", input: []byte("hello"), want: []byte("olleh")},
		{name: "even length", input: []byte("abcd"), want: []byte("dcba")},
		{name: "single character", input: []byte("a"), want: []byte("a")},
		{name: "empty", input: []byte(""), want: []byte("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := append([]byte(nil), tt.input...)
			reverseString(got)

			if string(got) != string(tt.want) {
				t.Fatalf("reverseString(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestReverseStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{name: "example", s: "abcdefg", k: 2, want: "bacdfeg"},
		{name: "exact 2k", s: "abcd", k: 2, want: "bacd"},
		{name: "less than k remaining", s: "abc", k: 5, want: "cba"},
		{name: "between k and 2k remaining", s: "abcdef", k: 4, want: "dcbaef"},
		{name: "single character groups", s: "abc", k: 1, want: "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.s, tt.k); got != tt.want {
				t.Fatalf("reverseStr(%q, %d) = %q, want %q", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func TestReplaceNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "example", s: "a1b2c3", want: "anumberbnumbercnumber"},
		{name: "no digits", s: "abc", want: "abc"},
		{name: "all digits", s: "123", want: "numbernumbernumber"},
		{name: "digit at both ends", s: "1abc2", want: "numberabcnumber"},
		{name: "zero digit", s: "a0b", want: "anumberb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceNumber(tt.s); got != tt.want {
				t.Fatalf("replaceNumber(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
