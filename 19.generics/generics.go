package main

import "fmt"

func main() {
	printItems([]int{1, 2, 3, 4})
	printItems([]float32{4.50, 3.82})
	printItems([]string{"Vincenzo", "Cassano"})
	printItems([]bool{true, false, true})
}

// any for all, can define custom too
func printItems[T any /*comparable | string | int*/](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
