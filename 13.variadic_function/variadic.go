package main

import "fmt"

func main() {
	nums := []int{1, 5, 4}
	//sumResult := sum(1, 5, 4) // or
	sumResult := sum(nums...)

	fmt.Println(sumResult)

	printAnything("Vincenzo", 45, "Cassano", true)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func printAnything(anything ...interface{}) {
	fmt.Println(anything...)
}
