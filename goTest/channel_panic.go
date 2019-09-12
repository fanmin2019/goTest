package main

func main() {
    ch := make(chan int)

    //对于本程序，没有其他的协程从 ch 接收数据。于是程序触发 panic，出现如下运行时错误
    ch <- 5
}
