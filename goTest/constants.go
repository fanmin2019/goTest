package main

import "fmt"

//定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます
//定数は := を使って宣言できません
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
