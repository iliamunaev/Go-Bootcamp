package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("Int")
		default:
			fmt.Printf("Other: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(1.3)
	whatAmI("hi")
}
