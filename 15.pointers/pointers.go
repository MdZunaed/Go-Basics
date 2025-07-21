package main

func main() {
	 num:= 1
	 println("before num value", num)

	 // & to send memory reference
	 changeNum(&num)
	 println("after num value", num)

}

// * for memory location reference (pointer)
func changeNum(num *int){
	*num = 5
	println("in changeNum function", *num)
}