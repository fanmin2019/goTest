package main

import (
	"fmt"
	"time"
)

//何とbreakを書く必要ない
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}



package main

import (
	"fmt"
	"time"
)


func farAwary(today time.Weekday) time.Weekday{
  var result time.Weekday = today + 3
  return result
}

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
  fmt.Printf("Type: %T Value: %v\n", today, today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case farAwary(today):
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
