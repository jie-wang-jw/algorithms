package _1_array

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	for top <= bottom && left <= right {
		// 1. 从左到右，填 top 行
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++

		// 2. 从上到下，填 right 列
		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--

		// 3. 从右到左，填 bottom 行
		// 先判断是否还有行
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans[bottom][i] = num
				num++
			}
			bottom--
		}

		// 4. 从下到上，填 left 列
		// 先判断是否还有列
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
