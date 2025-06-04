package main

import "fmt"

func main() {
	a := 1
	b := 2

	if a < b {
		fmt.Println("a less b")
	}

	if 2/2 == 1 || a > b {
		fmt.Println("At least one contition is true")
	}

	if 2/2 == 1 && a < b {
		fmt.Println("Both conditions are true")
	}

	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("0 < num < 10 ")
	} else {
		fmt.Println("num > 10")
	}
}
