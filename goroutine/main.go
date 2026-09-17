package main

import (
	"fmt"
	"sync"
)

// main
//
// 100 个协程各监听一个 channel，数字按 num%100 分配，余数 0 对应第 100 个协程。
// Each of 100 workers receives from its own channel; num%100 chooses the worker, with remainder zero mapped to worker 100.
// 主协程发送一个数字后等待 done，打印完成才能发送下一个，因此输出顺序确定为 1..1000。
// The sender waits for done after each number, so printing finishes before the next send and output stays ordered from 1 to 1000.
// 最后由发送方关闭全部输入 channel，再用 WaitGroup 等待所有协程退出。
// The sender closes every input channel after the final acknowledgment, then waits for all workers to exit.
func main() {
	done := make(chan struct{})

	chMap := make(map[int]chan int, 100)

	for i := 1; i <= 100; i++ {
		chMap[i] = make(chan int)
	}

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		id := i

		go func() {
			defer wg.Done()

			for num := range chMap[id] {
				fmt.Printf("goroutine %d: %d\n", id, num)
				done <- struct{}{}
			}
		}()
	}

	for num := 1; num <= 1000; num++ {
		id := num % 100
		if id == 0 {
			id = 100
		}

		chMap[id] <- num

		<-done
	}

	for i := 1; i <= 100; i++ {
		close(chMap[i])
	}

	wg.Wait()
}
