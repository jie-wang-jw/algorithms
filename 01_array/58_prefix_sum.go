package _1_array

/*
题目描述 / Problem Description
给定一个整数数组和多组闭区间查询 [left, right]，对每组查询输出该区间内所有元素的总和。
Given an integer array and multiple inclusive range queries [left, right], output the sum of all elements in each queried range.

解题思路 / Solution Approach
预先构造前缀和 prefix，其中 prefix[i] 表示前 i 个元素之和。每个区间和可用 prefix[right+1] - prefix[left] 在 O(1) 时间内得到。
Build a prefix-sum array where prefix[i] is the sum of the first i elements. Each range sum is then prefix[right+1] - prefix[left] in O(1) time.

关键逻辑：为什么这样做 / Why This Works
prefix[right+1] 包含下标 0 到 right，prefix[left] 包含下标 0 到 left-1。相减后公共的前半段抵消，恰好剩下闭区间 [left,right]。
例如 nums=[2,4,6]，查询 [1,2] 就是 (2+4+6)-2=10。prefix[0]=0 使 left=0 时也能直接套公式。
prefix[right+1] sums indices 0 through right, while prefix[left] sums 0 through left-1. Subtraction cancels the shared prefix,
leaving exactly [left,right]. For [2,4,6], query [1,2] gives (2+4+6)-2=10. prefix[0]=0 handles left=0 without a special case.

时间与空间复杂度 / Time and Space Complexity
n 为元素数，q 为查询数。预处理 O(n)，每次查询 O(1)，总时间 O(n+q)，按整数读写为常数成本计。辅助空间 O(n)，
当前实现保存 nums 和 prefix；输出逐条打印，不累计保存。
For n values and q queries, preprocessing takes O(n), each query O(1), and total time O(n+q),
treating integer I/O as constant cost. Auxiliary space O(n) stores nums and prefix; answers are printed without accumulation.
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
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	nums := make([]int, n)
	// prefix[i] stores the sum of the first i numbers; prefix[0] is 0.
	// prefix[i] 表示前 i 个数的和；prefix[0] 固定为 0。
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(in, &nums[i]); err != nil {
			return
		}
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
