package main

import "fmt"

func main() {
	var nums [5]int // Max 5 length
	ids := []int{1, 2, 3, 5} // Not fixed length
	var names [6]string


	twoD := [2][2]int{{1, 2}, {4, 5}}

	nums[1] = 2
	names[0] = "Vincenzo"
	names[1] = "Casssano"
	ids[3] = 6

	fmt.Println(nums)
	fmt.Println(names)
	fmt.Println(ids)
	fmt.Println(twoD)
	// fmt.Println(len(nums)) To get length
}
