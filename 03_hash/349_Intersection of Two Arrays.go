package _3_hash

/*
题目描述 / Problem Description
给定两个整数数组 nums1 和 nums2，返回它们的交集。结果中的每个元素必须唯一，返回顺序不限。
Given two integer arrays nums1 and nums2, return their intersection. Every result value must be unique, and the order does not matter.

解题思路 / Solution Approach
先把 nums1 的元素放入哈希集合，再遍历 nums2。命中集合时加入结果并从集合删除，确保相同数字只加入一次。
Put all values from nums1 into a hash set, then scan nums2. When a value is found, append it and remove it from the set so it can appear only once.
*/

/*
I use a hash set to store all numbers from the first array.
我使用哈希集合保存第一个数组中的所有数字。
Then I iterate through the second array.
然后遍历第二个数组。
If a number exists in the set, I add it to the result and remove it from the set to avoid duplicates.
如果数字存在于集合中，就把它加入结果，并从集合删除以避免重复结果。
*/

func intersection(nums1 []int, nums2 []int) []int {
	// The map acts as a set; only key existence matters.
	// 这里把 map 当作集合使用，只关心 key 是否存在。
	set := map[int]bool{}
	result := []int{}

	for _, num := range nums1 {
		// Repeated values simply overwrite true, so the set remains unique.
		// 重复数字只会再次写入 true，因此集合中仍然只有一份。
		set[num] = true
	}

	for _, num := range nums2 {
		// A true value means num appeared in nums1 and has not been used yet.
		// true 表示 num 在 nums1 中出现过，并且还没有加入结果。
		if set[num] {
			result = append(result, num)
			//set[num] = false
			// Delete after use so duplicates in nums2 cannot be appended again.
			// 使用后删除，这样 nums2 中重复的 num 不会再次加入结果。
			delete(set, num)
		}
	}

	return result
}
