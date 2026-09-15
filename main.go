package main

import (
	"fmt"
)

// main prints a greeting and five integer-division examples.
// main 输出问候语，并演示 100/i 的五次整数除法结果。
func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
