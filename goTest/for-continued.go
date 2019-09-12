package main

import "fmt"

//不思議だ
func main() {
	sum := 1

  //初期化と後処理ステートメントの記述は任意です。
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
