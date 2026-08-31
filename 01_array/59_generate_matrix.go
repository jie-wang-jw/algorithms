package _1_array

func generateMatrix(n int) [][]int {
	// Create an n by n matrix initialized with zeros.
	// 创建一个 n x n 的二维数组，初始值都为 0。
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
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
