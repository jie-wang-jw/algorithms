package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定整数数组 nums 和整数 target，找出两个和为 target 的不同元素下标并返回。
题目保证恰好存在一个答案。
Given an integer array nums and an integer target, return the indices of two distinct elements whose sum equals target.
Exactly one answer is guaranteed.
*/

// 1. Hash map in one pass: store every seen value with its index and look up target-num. Recommended.
// 1. 哈希表一次遍历：把见过的数字连同下标存入哈希表，再查找 target-num。推荐。
// Time: O(n) expected, Space: O(n) for the hash map.
// 时间复杂度：期望 O(n)，空间复杂度：O(n)，由哈希表产生。
//
// seen 保存当前下标之前的“值→下标”；先查 target-num 再存 num，避免同一元素被使用两次。
// seen maps earlier values to indices; look up target-num before insertion to avoid using one element twice.
func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		need := target - num
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}

// 2. Brute-force double loop: enumerate every index pair i<j so no element is reused.
// 2. 暴力双循环：枚举所有 i<j 的下标对，保证不重复使用同一个元素。
// Time: O(n²), Space: O(1).
// 时间复杂度：O(n²)，空间复杂度：O(1)。
//
// j 从 i+1 开始，既保证下标不同，也避免重复检查同一无序对；找到即返回原下标。
// Start j at i+1 to keep indices distinct and check each unordered pair once; return original indices on a match.
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 3. Sort + two pointers: copy value-index pairs first, because sorting would otherwise destroy the original indices.
// 3. 排序 + 双指针：必须先把数值和原下标一起复制出来，否则排序会丢失要返回的下标。
// Time: O(n log n), Space: O(n) for the value-index copy.
// 时间复杂度：O(n log n)，空间复杂度：O(n)，由数值与下标的副本产生。
//
// pairs 同时保存值和原下标，排序仅改变 pairs；和偏小排除左端，偏大排除右端，命中返回原下标。
// pairs retain original indices without mutating nums; discard the left endpoint for a small sum or the right for a large sum.
func twoSumSorted(nums []int, target int) []int {
	pairs := make([][2]int, len(nums))
	for i, v := range nums {
		pairs[i] = [2]int{v, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	left, right := 0, len(pairs)-1

	for left < right {
		sum := pairs[left][0] + pairs[right][0]

		if sum == target {
			return []int{pairs[left][1], pairs[right][1]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
