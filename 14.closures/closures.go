package main

import "fmt"

func main() {
	increament := counter()

	fmt.Println(increament())
	fmt.Println(increament())
	fmt.Println(increament())
}

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}
