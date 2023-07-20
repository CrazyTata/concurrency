package main

import (
	"fmt"
	"time"
)

// 四个worker 轮询打印 1 2 3 4
func newWorker(id int, ch chan int, nextCh chan int) {
	for {
		token := <-ch         // 取得令牌
		fmt.Println((id + 1)) // id从1开始
		fmt.Println(token)    //
		time.Sleep(time.Second)
		nextCh <- token
	}
}
func main() {
	chs := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	//首先把令牌交给第一个worker
	chs[0] <- 1

	select {}
}
