package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

//これは最初の要素を含む half-open レンジを選択しますが、最後の要素は省かれます。
	var s []int = primes[1:4]  //[3, 5, 7]
	fmt.Println(s)
}
