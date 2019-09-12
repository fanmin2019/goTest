package main

import "fmt"

//定義とかを纏める必要があるよね
//あと制御をコントロールするときに
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
