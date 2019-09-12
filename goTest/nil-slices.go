package main

import "fmt"

func main() {

  //スライスのゼロ値は nil です。
//nil スライスは 0 の長さと容量を持っており、何の元となる配列も持っていません。
	var s []int
  //配列自体は要素が一つもないとしても、nilではない
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
