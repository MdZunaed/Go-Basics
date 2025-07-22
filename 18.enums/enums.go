package main

import "fmt"

func main() {
	changeOrderStatus(Prepared)
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changed status to", status)
}

type OrderStatus string

const (
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
	Recieved  OrderStatus = "recieved"
)

/// With int,iota
// type OrderStatus int

// const (
// 	Confirmed OrderStatus = iota
// 	Prepared
// 	Delivered
// 	Recieved
// )
