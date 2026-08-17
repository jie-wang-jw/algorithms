package main

import (
	"fmt"
	"sync"
)

func main() {
	// 用来等待每个数字打印完成
	done := make(chan struct{})

	// 100 个协程对应 100 个 channel
	chMap := make(map[int]chan int, 100)

	for i := 1; i <= 100; i++ {
		chMap[i] = make(chan int)
	}

	var wg sync.WaitGroup
	wg.Add(100)

	// 开启 100 个协程
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

	// 顺序发送 1 到 1000
	for num := 1; num <= 1000; num++ {
		id := num % 100
		if id == 0 {
			id = 100
		}

		// 把数字发送给对应协程
		chMap[id] <- num

		// 等待这个数字打印完成
		<-done
	}

	// 关闭所有 channel，让协程退出
	for i := 1; i <= 100; i++ {
		close(chMap[i])
	}

	wg.Wait()
}
