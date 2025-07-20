package main

import (
	"fmt"
	"slices"
)

func main() {
	/// un-initialized slice is nil/null // nil in Go
	var nums []int
	fmt.Println("nums:", nums)

	/// default value
	var numsTwo = make([]int, 2)                          // initial length will be 2
	var numsThree = make([]int, 2, 5)                     // max length will be 5
	fmt.Printf("numThree capacity: %d\n", cap(numsThree)) // To check capacity
	fmt.Printf("numsTwo before append: %d\n", numsTwo)

	/// To add element
	// nums = append(nums, 8)
	// fmt.Println(nums)

	fmt.Printf("numsTwo capacity before append: %d\n", cap(numsTwo))
	numsTwo = append(numsTwo, 3)
	numsTwo = append(numsTwo, 4)
	fmt.Printf("numsTwo after append: %d\n", numsTwo)
	fmt.Printf("numsTwo capacity after append: %d\n", cap(numsTwo))

	/// Slice range

	numsFour := []int{1, 3, 5, 7, 9, 11}

	fmt.Println("\nnumFour:", numsFour)
	fmt.Println("numFour ranged[1:3]:", numsFour[1:4])
	fmt.Println("numFour ranged[2:]:", numsFour[2:])
	fmt.Println("numFour ranged[:3]:", numsFour[:4])

	var sliceOne = []int{1,5,8}
	var sliceTwo = []int{1,2,7}
	var sliceThree = []int{1,5,8}

	fmt.Println("\nsliceOne & sliceTwo equal? :", slices.Equal(sliceOne, sliceTwo))
	fmt.Println("sliceOne & sliceThree equal? :", slices.Equal(sliceOne, sliceThree))

}
