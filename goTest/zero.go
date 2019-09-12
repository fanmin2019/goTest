package main

import "fmt"

func main() {
	var i int //0
	var f float64 //0
	var b bool //false
	var s string //""
  //%qは文字列にクォーテーションで囲むこと
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
