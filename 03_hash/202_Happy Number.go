package _3_hash

/*I repeatedly replace the number with the sum of the squares of its digits.
If it eventually becomes one, it is a happy number.
If I see the same number again, that means we are in a cycle, so it cannot reach one.*/

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

func getNext(n int) int {
	sum := 0

	/*
		I extract each digit using n mod 10, add its square to the sum, and then divide n by 10 to move to the next digit.
	*/
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
