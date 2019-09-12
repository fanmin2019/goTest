package main

import "fmt"

func sum(s []int, c chan int) {
  fmt.Println(s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}


//信道是 进程内 的通讯方式，是不支持跨进程通信的。如果需要跨进程间通讯的话，可以使用 Socket 等网络方式
//应该像对 map 和切片所做的那样，用 make 来定义信道
	c := make(chan int)
  fmt.Println("s[:len(s)/2", s[:len(s)/2])
	go sum(s[:len(s)/2], c) //7,2,8 = 17
//信道视为程序中的 FIFO（先进先出）数据队列

  fmt.Println("s[len(s)/2:]", s[len(s)/2:])
	go sum(s[len(s)/2:], c) //-9 4 0 = -5
	x, y := <-c, <-c // 从 c 中接收  //x + y = 12

	fmt.Println(x, y, x+y)
}
