package _3_hash

/*
题目描述 / Problem Description
将正整数 n 反复替换为其各位数字的平方和。如果最终得到 1，则 n 是快乐数；
如果进入不包含 1 的循环，则不是。判断 n 是否为快乐数。
Repeatedly replace a positive integer n with the sum of the squares of its digits.
Return whether this process eventually reaches 1 rather than entering a cycle.

解题思路 / Solution Approach
使用哈希集合记录所有出现过的中间结果。结果变成 1 时返回 true；
某个结果重复出现时说明进入循环，返回 false。
Store every intermediate value in a hash set. Return true upon reaching 1;
if a value repeats, the process has entered a cycle, so return false.

关键逻辑：为什么这样做 / Why This Works
同一个数字经过 getNext 总会得到同一个结果，所以一旦重复出现，后面的整条变化路径也会重复，
不可能突然走出循环到达 1。循环条件已经排除了 n=1，因此检测到的重复是非快乐循环。
平方和把大数压到有限范围内，最终必然到达 1 或重复，不会无限产生不同数字。
getNext is deterministic: the same number always has the same successor.
Revisiting a number repeats the entire future path, so it cannot newly escape to 1.
The loop condition already excludes 1, making any detected repetition a non-happy cycle.
Digit-square sums eventually enter a bounded range, so the process must reach 1 or repeat.

时间与空间复杂度 / Time and Space Complexity
d 为初始数字的十进制位数。首次平方和计算 O(d)，之后数值至多 81d 并继续快速缩小到固定范围，因此时间 O(d)，即 O(log(n+1))。
哈希集合空间可保守记为 O(d)，不是 O(n)。单次 getNext(x) 时间与 x 的位数成正比，辅助空间 O(1)。哈希操作按平均 O(1) 计。
Let d be the initial decimal digit count. The first digit-square sum costs O(d);
the value then becomes at most 81d and rapidly shrinks to a fixed range, giving O(d), or O(log(n+1)), time.
A conservative hash-set space bound is O(d), not O(n). getNext(x) takes time proportional to its digit count
and O(1) auxiliary space. Hash operations are average O(1).

补充解法：快慢指针 / Alternative: Floyd Cycle Detection
isHappyFloyd 把每个整数看作节点，平方和变换 getNext 就是唯一的 Next。
慢指针变换一次、快指针两次。相遇在 1 表示进入 1->1，否则进入不含 1 的环。
正整数经过变换会进入有限范围，因此一定到 1 或进入环；不需要显式集合。
d 为输入十进制位数，时间 O(d)，辅助空间 O(1)，复用原 getNext。
Treat each integer as a node and getNext as its unique successor.
Advance one and two steps; meeting at 1 means success, any other cycle means failure.
Digit-square sums eventually enter a bounded range. Time O(d) for d input digits, O(1) auxiliary space.
*/

/*
I repeatedly replace the number with the sum of the squares of its digits.
不断把数字替换成“各位数字的平方和”。

If it eventually becomes one, it is a happy number.
如果最终变成 1，它就是快乐数。

If I see the same number again, that means we are in a cycle, so it cannot reach one.
如果某个中间结果再次出现，说明进入循环，之后不可能到达 1。
*/

// 1. 哈希集合判重
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

// 各位数字平方和：两种解法共用的变换
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

// 2. 快慢指针判环：不需要哈希集合
func isHappyFloyd(n int) bool {
	// Every number has exactly one successor, so the sequence behaves like a linked list.
	// 每个数字只有一个后继，因此整条变换序列相当于一个链表。
	// fast starts one step ahead so the loop condition is not satisfied immediately.
	// fast 先走一步，这样循环条件不会在开始时就成立。
	slow, fast := n, getNext(n)

	// fast moves twice as quickly, so it catches slow inside any cycle.
	// fast 的速度是 slow 的两倍，只要存在环就一定会追上 slow。
	for slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	// 1 maps to itself, so meeting at 1 means the number is happy.
	// 1 的后继还是 1，因此在 1 处相遇就说明是快乐数；在其他值相遇则是不含 1 的环。
	return slow == 1
}
