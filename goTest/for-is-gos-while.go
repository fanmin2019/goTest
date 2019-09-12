package main

import "fmt"

//C言語などにある while は、Goでは for だけを使います
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
