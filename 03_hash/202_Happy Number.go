package _3_hash

/*
I repeatedly replace the number with the sum of the squares of its digits.
不断把数字替换成“各位数字的平方和”。

If it eventually becomes one, it is a happy number.
如果最终变成 1，它就是快乐数。

If I see the same number again, that means we are in a cycle, so it cannot reach one.
如果某个中间结果再次出现，说明进入循环，之后不可能到达 1。
*/

func isHappy(n int) bool {
	// seen records every intermediate number we have processed.
	// seen 记录已经处理过的每个中间结果。
	seen := map[int]bool{}

	for n != 1 {
		// Seeing n again means the transformation is cycling.
		// 再次看到 n，表示后续变化开始循环。
		if seen[n] {
			return false
		}

		// Mark the current number before calculating the next one.
		// 计算下一个数字前，先把当前数字标记为已访问。
		seen[n] = true
		n = getNext(n)
	}
	return true
}

func getNext(n int) int {
	// sum accumulates the square of every extracted digit.
	// sum 累加每一位数字的平方。
	sum := 0

	/*
		I extract each digit using n mod 10, add its square to the sum, and then divide n by 10 to move to the next digit.
		使用 n % 10 取出个位数，把它的平方加入 sum，再用 n / 10 删除个位数。
	*/
	for n > 0 {
		// Example: 19 % 10 = 9, so digit is the last digit.
		// 例如 19 % 10 = 9，因此 digit 是当前个位数。
		digit := n % 10
		sum += digit * digit
		// Integer division removes the last digit: 19 / 10 = 1.
		// 整数除法会去掉个位数：19 / 10 = 1。
		n /= 10
	}
	return sum
}
