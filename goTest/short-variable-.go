package main
//https://repl.it/languages/go
import "fmt"

//関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。
func main() {
	var i, j int = 1, 2
  //短い := の代入文を使い、暗黙的な型宣言ができます。
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
