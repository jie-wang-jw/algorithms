package _1_array

/*
题目描述 / Problem Description
给定一个整数数组和多组闭区间查询 [left, right]，对每组查询输出该区间内所有元素的总和。
Given an integer array and multiple inclusive range queries [left, right], output the sum of all elements in each queried range.

解题思路 / Solution Approach
预先构造前缀和 prefix，其中 prefix[i] 表示前 i 个元素之和。每个区间和可用 prefix[right+1] - prefix[left] 在 O(1) 时间内得到。
Build a prefix-sum array where prefix[i] is the sum of the first i elements. Each range sum is then prefix[right+1] - prefix[left] in O(1) time.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func prefixSum() {
	// Use buffered input because the problem may contain many numbers and queries.
	// 使用缓冲输入，因为题目可能包含大量数字和查询。
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	// prefix[i] stores the sum of the first i numbers; prefix[0] is 0.
	// prefix[i] 表示前 i 个数的和；prefix[0] 固定为 0。
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
		// Add nums[i] to the previous prefix sum.
		// 在前一个前缀和的基础上加上 nums[i]。
		prefix[i+1] = prefix[i] + nums[i]
	}

	var left, right int
	// Read queries until the input ends.
	// 持续读取查询，直到输入结束。
	for {
		_, err := fmt.Fscan(in, &left, &right)
		if err != nil {
			break
		}

		// Sum of [left, right] = prefix before right+1 - prefix before left.
		// 闭区间 [left, right] 的和 = right+1 前缀和 - left 前缀和。
		sum := prefix[right+1] - prefix[left]
		fmt.Println(sum)
	}
}
