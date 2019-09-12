package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
  //条件のないswitchは、 switch true と書くことと同じです
  // "if-then-else" のつながり
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
