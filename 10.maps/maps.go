package main

import "fmt"

func main() {
	/// Un-assigned
	mapObj := make(map[int]string)

	mapObj[197] = "Vincenzo"
	mapObj[198] = "Cassano"

	fmt.Println(mapObj)
	fmt.Println(mapObj[197])

	/// Assigned
	mapMap := map[int]string{1541: "Vincenzo", 1542: "Cassano"}
	println(mapMap[1542])

	mapKey, mapValue := mapMap[1541]
	mapKey2, mapValue2 := mapMap[1545]
	if mapValue {
		println("1541 Value found", mapValue)
	} else {
		println("1541 Not found", mapKey)
	}
	if mapValue2 {
		println("1545 Value found", mapValue2)
	} else {
		println("1545 Not found", mapKey2)
	}
}
