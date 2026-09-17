package _3_hash

/*
题目描述 / Problem Description
将正整数 n 反复替换为其各位数字的平方和。如果最终得到 1，则 n 是快乐数；
如果进入不包含 1 的循环，则不是。判断 n 是否为快乐数。
Repeatedly replace a positive integer n with the sum of the squares of its digits.
Return whether this process eventually reaches 1 rather than entering a cycle.
*/

// 1. Hash set of seen values: a repeated intermediate number means the sequence is cycling.
// 1. 哈希集合判重：中间结果再次出现即进入循环。
// Time: O(d) for d decimal digits, Space: O(d) for the hash set.
// 时间复杂度：O(d)（d 为十进制位数），空间复杂度：O(d)，由哈希集合产生。
//
// 每个整数的下一状态唯一；若重复遇到同一状态而尚未到 1，就会永远重复该环，因此返回 false。
// Each value has one successor; revisiting a state before reaching 1 traps the sequence in a nonhappy cycle.
// 平方和很快落入有限范围，所以最终必到 1 或进入环；题目要求 n 为正整数。
// Digit-square sums soon enter a finite range, forcing 1 or a cycle; the problem assumes positive n.
func isHappy(n int) bool {
	seen := map[int]bool{}

	for n != 1 {
		if seen[n] {
			return false
		}

		seen[n] = true
		n = getNext(n)
	}
	return true
}

// Digit-square sum: the shared transform used by both solutions.
// 各位数字平方和：两种解法共用的变换。
// Time: O(d) for a d-digit argument, Space: O(1).
// 时间复杂度：O(d)（d 为参数的十进制位数），空间复杂度：O(1)。
//
// 每次用 n%10 取末位、n/=10 去末位，累加每位平方；此辅助函数按非负整数使用。
// Take the last digit with n%10, remove it with n/=10, and sum its square; this helper assumes nonnegative input.
func getNext(n int) int {
	sum := 0

	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// 2. Floyd's slow and fast pointers: treat getNext as the unique successor; meeting at 1 is happy. No hash set needed.
// 2. 快慢指针判环：把 getNext 当作唯一后继，相遇在 1 即快乐数；不需要哈希集合。
// Time: O(d) for d decimal digits, Space: O(1).
// 时间复杂度：O(d)（d 为十进制位数），空间复杂度：O(1)。
//
// 把“数位平方和”看成唯一后继，快慢指针在最终的环中相遇；幸福数的环只有固定点 1。
// Treat the digit-square sum as a successor; Floyd pointers meet in the eventual cycle, which is the fixed point 1 for happy numbers.
// fast 先走一步，避免初始化时 slow==fast 直接跳过循环；结束后判断相遇值是否为 1。
// Start fast one step ahead to avoid immediate false termination; success means the meeting value is 1.
func isHappyFloyd(n int) bool {
	slow, fast := n, getNext(n)

	for slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return slow == 1
}
