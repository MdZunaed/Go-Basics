package main

func main() {
	//call function
	result := add(2, 6)
	println(result)

	// get multiple value
	println(getLanguages())

	// store function
	fn := increaseOne()
	println(fn(5))

}

func add(a int, b int) int {
	return a + b
}

// to return multiple value
func getLanguages() (string, string, string) {
	return "Go", "JS", "Dart"
}

// to return a function
func increaseOne() func(a int) int {
	return func(a int) int {
		return a + 1
	}
}
