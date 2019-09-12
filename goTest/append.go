package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
  //append への最初のパラメータ s は、追加元となる T 型のスライスです。
  //残りの vs は、追加する T 型の変数群です。
	s = append(s, 0)
  //append の戻り値は、追加元のスライスと追加する変数群を合わせたスライス
  //つまり、毎回のポインタは変わっている
  //新しい割当先を示す
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
