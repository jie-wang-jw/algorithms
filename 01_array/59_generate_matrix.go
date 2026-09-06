package _1_array

/*
题目描述 / Problem Description
给定正整数 n，生成一个 n x n 矩阵，按照顺时针螺旋顺序填入从 1 到 n² 的所有整数。
Given a positive integer n, generate an n x n matrix filled with the integers from 1 to n² in clockwise spiral order.

解题思路 / Solution Approach
维护上、下、左、右四条未填充边界。每轮依次填充上边、右边、下边和左边，然后将四条边界向内收缩。
Maintain top, bottom, left, and right boundaries for the unfilled area. Fill the four sides clockwise, then move all boundaries inward.

关键逻辑：为什么这样做 / Why This Works
四条边围住尚未填写的区域。每填完一条边就立即收缩该边界，而不是等四条边都写完才收缩：例如 top++ 后，右边从下一行开始，右上角不会重复填写；
其他角同理。后两条边先检查剩余行/列，避免处理已经消失的区域。n=3 填完外圈后四个边界都为 1，下一轮只填中心的 9，随后边界交错，循环结束。
The four boundaries enclose unfilled cells. Shrink each boundary immediately after filling that side, not after all four sides:
top++ makes the right column start one row lower, avoiding a second write to the top-right corner. Other corners follow the same rule.
Check remaining rows/columns before the last two sides. For n=3, all boundaries become 1 after the outer ring;
the next round fills only the center with 9, then the boundaries cross.

时间与空间复杂度 / Time and Space Complexity
n 为矩阵边长。时间 O(n²)，每个格子填入一次。除返回矩阵外辅助空间 O(1)；返回矩阵占 O(n²)，包含结果的总空间为 O(n²)。
n is the matrix side length. Time O(n²), filling each cell once. Auxiliary space excluding the output is O(1);
the returned matrix and total space take O(n²).
*/

func generateMatrix(n int) [][]int {
	// Create an n by n matrix initialized with zeros.
	// 创建一个 n x n 的二维数组，初始值都为 0。
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	// These four boundaries describe the unfilled rectangle.
	// 这四个边界围住当前尚未填充的矩形区域。
	top, bottom := 0, n-1
	left, right := 0, n-1
	// num is the next value to write into the matrix.
	// num 是下一个要写入矩阵的数字。
	num := 1

	// Fill one outer layer clockwise, then move all boundaries inward.
	// 每轮顺时针填完一层，再把四条边界向内收缩。
	for top <= bottom && left <= right {
		// 1. Fill the top row from left to right. / 从左到右填充 top 行。
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++

		// 2. Fill the right column from top to bottom. / 从上到下填充 right 列。
		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--

		// 3. Fill the bottom row from right to left. / 从右到左填充 bottom 行。
		// Check first because the top row may have used the last remaining row.
		// 先判断是否还有行，因为 top 行可能已经填掉最后一行。
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans[bottom][i] = num
				num++
			}
			bottom--
		}

		// 4. Fill the left column from bottom to top. / 从下到上填充 left 列。
		// Check first because the right column may have used the last remaining column.
		// 先判断是否还有列，因为 right 列可能已经填掉最后一列。
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans[i][left] = num
				num++
			}
			left++
		}
	}

	return ans
}
