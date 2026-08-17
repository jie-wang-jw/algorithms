package _1_array

import (
	"bufio"
	"fmt"
	"os"
)

func prefixSum() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
		prefix[i+1] = prefix[i] + nums[i]
	}

	var left, right int
	for {
		_, err := fmt.Fscan(in, &left, &right)
		if err != nil {
			break
		}

		sum := prefix[right+1] - prefix[left]
		fmt.Println(sum)
	}
}
