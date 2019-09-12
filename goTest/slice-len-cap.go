package main

import "fmt"

func main() {

	s := []int{2, 3, 5, 7, 11, 13}
  //要素の数＝長さ=6
  //容量は2から数えて、元の配列の要素数＝6
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0] //{}
  //長さ=0,容量=2から数えて、元の配列の要素数＝6
	printSlice(s)

	// Extend its length.
	s = s[:4] //{2,3,5,7}
  //長さ=4,容量=2から数えて、元の配列の要素数＝6
	printSlice(s)

	// Drop its first two values.
  //スライスのスライス
	s = s[2:] //{5,7}
  //長さ=2,容量=5から数えて、元の配列の要素数＝4
	printSlice(s)
}

func printSlice(s []int) {
  //スライスは長さ( length )と容量( capacity )の両方
  //スライスの長さは、それに含まれる要素の数
  //スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
