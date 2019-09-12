package main

import (
    "fmt"
)

func hello(done chan bool) {
    fmt.Println("Hello world goroutine")
    done <- true
    //もし「<-done」がなかったら、ブロッキングがないので、helloを実行する前に「main function」が実行される
}

func main() {
    done := make(chan bool)
    go hello(done)
    //ブロッキングがある、这一行代码发生了阻塞，除非有协程向 done 写入数据，否则程序不会跳到下一行代码。
    <-done //もし「done <- true」がなかったら、誰も送信してくれないpanic
    fmt.Println("main function")
}

//結果　先にHello world goroutine、後でmain function
