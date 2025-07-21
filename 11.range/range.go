package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 0}

	/// regular iteration
	// for i := 0; i < len(nums); i++ {
	// 	println(nums[i])
	// }

	/// iteration with range
	for i, num := range nums { // where i = index, num = item
		fmt.Printf("Index: %d, num: %d\n", i, num)
	}

	/// with map
	mapMap := map[string]string{"firstName": "Vincenzo", "lastName": "Cassano"}

	for k, v := range mapMap {
		println(k, v)
	}

	/// with string

	// i = starting byte of rune
	// c = unicode point rune
	for i, c := range "Vincenzo" {
		println(i, string(c))
	}
}
