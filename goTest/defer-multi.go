package main

import "fmt"

func main() {
	fmt.Println("counting")

//defer へ渡した関数は LIFO(last-in-first-out) の順番
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
