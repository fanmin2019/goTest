package main

import "fmt"

func main() {
  //スライスは、組み込みの make 関数を使用して作成することができます。 これは、動的サイズの配列を作成する方法
  const aa int = 3
  var yy = [aa]int {1, 3, 4}

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

//b := make([]int, 0, 5) // len(b)=0, cap(b)=5
//b = b[:cap(b)] // len(b)=5, cap(b)=5
//b = b[1:]      // len(b)=4, cap(b)=4
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
