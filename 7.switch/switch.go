package main

import (
	"fmt"
	"time"
)

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
		// break // No need break in Go
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}
}
