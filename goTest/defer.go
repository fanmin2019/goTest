package main

import "fmt"

func main() {

  //defer へ渡した関数の引数は、すぐに評価されますが、
  //その関数自体は呼び出し元の関数がreturnするまで実行されません
	defer fmt.Println("world")
  defer fmt.Println("world2")
  defer fmt.Println("world3")
  defer fmt.Println("world4")

	fmt.Println("hello")
  fmt.Println("hello2")
  fmt.Println("hello3")
  //実行順番 hello⇒hello2⇒hello3⇒world4⇒world3⇒world2⇒world //スタックだからか!
}
