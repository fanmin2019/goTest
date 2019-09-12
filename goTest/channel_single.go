package main

import "fmt"

func sendData(sendch chan<- int) {
    sendch <- 10 //10を受信
}

func main() {
  //送信専用 sendch := make(<-chan int) 
    sendch := make(chan<- int) //受信専用のチャンネル
    go sendData(sendch)
    fmt.Println(<-sendch)  //我们试图通过唯送信道接收数据，于是编译器报错
}
