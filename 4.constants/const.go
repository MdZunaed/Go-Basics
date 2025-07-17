package main

import "fmt"

func main() {
	const name = "Vincenzo"
	/// Can't change const varaible
	// name = "Cassano"

	/// For multiple constants ///

	const (
		port = 3001
		host = "localhost"
	)

	fmt.Println(port)
	fmt.Println(host)
}