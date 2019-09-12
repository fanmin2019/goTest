package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//もしcloseがないと
	//その他のgoroutineは全て書き込み待ちのままになります
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)k

	//循环 for i := range c 会不断从信道接收值，直到它被关闭
	//　<-はないけど、ブロッキングなんだ
	//您的主线程和goroutine线程不是同步进程.
	//你的主线程永远不会知道何时停止从goroutines调用通道.
	//由于主线程在通道上循环,它总是从通道调用值,
	//当goroutines完成并且通道停止发送值时,
	//主线程无法从通道获取任何值,因此条件变为死锁.
	for i := range c {
		fmt.Println(i)
	}
}
