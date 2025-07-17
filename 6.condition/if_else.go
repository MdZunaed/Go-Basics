package main

import "fmt"

func main() {
	age := 180

	if age < 12 {
		fmt.Println("Person is kid")
	} else if age > 12 && age < 18 {
		fmt.Println("Person is Not adult")
	} else if age >= 18 && age < 150 {
		fmt.Println("Person is Adult")
	} else {
		fmt.Println("Invalid age")
	}

	/// Go doesn't have any Ternery operator (!)
	/// We have to use if else instead
}
