package main

import "fmt"

func main() {

  //ch := make(chan int, 1)
  //fatal error: all goroutines are asleep - deadlock!
  //发送与接收默认是阻塞的。
  //当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，
  //直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
  //与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
