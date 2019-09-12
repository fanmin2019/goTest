package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
    /*
    %e  科学的記数法、例: -1234.456e+78
%E  科学的記数法、例: -1234.456E+78
%f  指数なしの小数、例: 123.456
%g  %e、%fのどちらか出力の短い方
    */
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
