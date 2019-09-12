package main

import "fmt"

func main() {
	i, j := 42, 2701

  //& オペレータは、そのオペランド( operand )へのポインタを引き出します
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer

  //* オペレータは、ポインタの指す先の変数を示します。
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
