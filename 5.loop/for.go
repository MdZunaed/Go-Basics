package main

import "fmt"

func main() {
	/// For loop ///

	for i := 1; i <= 5; i++ {
		if i == 5 {
			fmt.Print("Skipping")
			continue // To skip current iteration
		}
		fmt.Println(i)
	}

	/// While loop ///

	i := 5

	for i <= 12 {
		if i == 10 {
			fmt.Print("Break")
			break // To break the loop
		}
		fmt.Println(i)
		i++
	}

	/// Infinite loop ///

	// for {
	// 	fmt.Println("Loading..")
	// }
}
