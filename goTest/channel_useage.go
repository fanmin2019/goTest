/**
package main

import  "fmt"

func  print() {
    fmt.Println("Hello world")
}

func  main() {
    for  i  :=  0; i <  10; i++ {
        go  print()
    }
}
⇒何も出力されない、print()は実行されない
**/

package main
import "fmt"

func print(ch chan int) {
    fmt.Println("Hello world")
    ch<- 1 //ブロッキング
}

func main() {
    chs := make([]chan int) //チャンネルのリストを作る
    for i := 0; i < 10; i++ {
        chs[i] = make(chan int) //ループの中でチャンネルを作る（空）
        go print(chs[i]) //チャンネルを渡す
    }

//循环 for i := range c 会不断从信道接收值，直到它被关闭
    for _, ch := range(chs){
        <-ch
    }
}
