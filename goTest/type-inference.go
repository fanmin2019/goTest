package main

import "fmt"

func main() {

/**

var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

**/
  //明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、変数の型は右側の変数から型推論されます
	v := 42 // change me!
  v := 42.1 // change me!
	fmt.Printf("v is of type %T\n", v)
}
