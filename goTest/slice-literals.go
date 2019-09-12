package main

import "fmt"

func main() {
  //スライスのリテラルは長さのない配列リテラルのようなもの
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

//スライスのリテラルは長さのない配列リテラルのようなもの
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

/**
type test struct {
		i int
		b bool
	}
	s := []test {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
*/
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
