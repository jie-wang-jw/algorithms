package _3_hash

/*
题目描述 / Problem Description
给定四个整数数组 nums1、nums2、nums3 和 nums4，统计满足 nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0 的下标四元组数量。
Given four integer arrays nums1, nums2, nums3, and nums4, count index tuples satisfying nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0.

解题思路 / Solution Approach
将四个数组分成两组。用哈希表统计前两个数组所有两数和的出现次数，再遍历后两个数组的两数和，累加其相反数在哈希表中的频次。
Split the arrays into two pairs. Count every sum from the first pair in a hash map, then scan sums from the second pair and add the frequency of each opposite sum.

时间与空间复杂度 / Time and Space Complexity
四个数组等长 n 时，平均时间 O(n²)，前两组配对各遍历 n² 次；辅助空间 O(n²)，保存不同两数和及频次。若长度分别为 a,b,c,d，则时间 O(ab+cd)，空间 O(ab)。
For four arrays of length n, average time is O(n²) and auxiliary space O(n²) for pair-sum frequencies. With lengths a,b,c,d, time is O(ab+cd) and space O(ab).
*/

/*
Split the four arrays into two pairs.
把四个数组分成两组处理。

First, I count all possible sums from nums1 and nums2 using a hash map.
先用哈希表统计 nums1 和 nums2 所有两数之和出现的次数。

Then, for each sum from nums3 and nums4, I look for its negative value in the map.
再遍历 nums3 和 nums4 的两数之和，并在哈希表中查找它的相反数。

The frequency in the map tells me how many valid tuples we can form.
哈希表中的频次表示当前组合能够配出多少个有效四元组。
*/

func fourSumCount(A []int, B []int, C []int, D []int) int {
	// Key: a + b; value: how many index pairs produce that sum.
	// key 是 a+b，value 是产生这个和的下标组合数量。
	sumCount := map[int]int{}
	count := 0

	for _, a := range A {
		for _, b := range B {
			// Do not store only true: duplicate pairs must all be counted.
			// 不能只保存 true，因为不同下标产生的重复组合也必须计数。
			sumCount[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			// We need (a+b) + (c+d) = 0, so a+b must equal -(c+d).
			// 要满足总和为 0，a+b 必须等于 -(c+d)。
			need := -(c + d)
			// A missing key has value 0 in Go, so it adds nothing automatically.
			// Go 中不存在的 map key 读取结果为 0，因此无需单独判断是否存在。
			count += sumCount[need]
		}
	}

	return count
}
