package _3_hash

/*
题目描述 / Problem Description
给定整数数组 nums，找出所有和为 0 的不重复三元组。每个三元组必须使用三个不同下标的元素。
Given an integer array nums, return all unique triplets whose sum is zero.
Each triplet must use elements at three distinct indices.
*/

import "sort"

// 1. Sort + two pointers: fix the first value, then search the rest from both ends. Recommended.
// 1. 排序 + 双指针：固定第一个数，再在右侧用左右指针寻找另外两个数。推荐。
// Time: O(n²), Space: O(log n+r) including the sort stack and r triplets.
// 时间复杂度：O(n²)，空间复杂度：O(log n+r)，含排序栈与结果。
//
// 排序后固定 i，在其右侧双指针求和。和偏小时，当前 left 即使配最大 right 也不足，因此可排除 left；偏大同理排除 right。
// After sorting and fixing i, too small a sum excludes left even with its largest partner; too large excludes right similarly.
// nums[i]>0 时后续全为正，可停止；同层跳过相同 i 值，找到答案后跳过两端重复值，避免重复三元组。
// Stop when nums[i]>0; skip repeated fixed values and repeated endpoints after a match to avoid duplicate triples.
// 修改输入顺序；i<left<right 保证下标不同，整数求和须不溢出。
// Sorting mutates input; i<left<right prevents index reuse, and sums must fit in int.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]

		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			b, c := nums[left], nums[right]
			sum := a + b + c

			if sum == 0 {
				res = append(res, []int{a, b, c})

				for left < right && nums[left] == b {
					left++
				}

				for left < right && nums[right] == c {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}

// 2. Fix one value + hash: reuse Two Sum by looking up b=-a-c before inserting c.
// 2. 固定一个数 + 哈希：复用两数之和，先查 b=-a-c 再插入 c。
// Time: O(n²) expected, Space: O(n+r) auxiliary plus O(r) output.
// 时间复杂度：期望 O(n²)，空间复杂度：辅助 O(n+r)，结果另占 O(r)。
//
// 排序并固定 a；seen 只存当前扫描位置之前的后缀值，先查 b=-a-c 再存 c，保证三个下标不同。
// Sort and fix a; query b=-a-c before storing c so seen only supplies an earlier, distinct suffix position.
// 有序三元组作为 recorded 的键统一去重；排序修改输入，seen 需为每个 a 重建。
// Use sorted triples as recorded keys; sorting mutates input and seen must reset for each a.
func threeSumHash(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	recorded := make(map[[3]int]bool)

	for i, a := range nums {
		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}

		seen := make(map[int]bool)

		for _, c := range nums[i+1:] {
			b := -a - c

			if seen[b] {
				key := [3]int{a, b, c}
				if !recorded[key] {
					recorded[key] = true
					result = append(result, []int{a, b, c})
				}
			}

			seen[c] = true
		}
	}

	return result
}
