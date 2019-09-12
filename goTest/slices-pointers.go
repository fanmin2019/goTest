package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

//スライスは配列への参照のようなもの
	a := names[0:2]
  //スライスの要素を変更すると、その元となる配列の対応する要素が変更され
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
  //同じ元となる配列を共有している他のスライスは、それらの変更が反映されます
	fmt.Println(names)
}
