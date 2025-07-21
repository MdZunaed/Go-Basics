package main

import (
	"fmt"
	"time"
)

func main() {
	newCustomer := customer{
		name:  "Vincenzo",
		phone: "017..",
	}
	myOrder := order{
		id:       "trx1234",
		ammount:  50.8,
		status:   "recieved",
		customer: newCustomer,
	}

	myOrder2 := order{
		id:        "trx5678",
		ammount:   100.2,
		status:    "pending",
		createdAt: time.Now(),
		customer: customer{
			name:  "Cassano",
			phone: "018..",
		},
	}

	orderWithFunc := newOrder("trx", 25, "canceled")

	myOrder.createdAt = time.Now()
	myOrder2.status = "paid"
	myOrder.changeStatus("Processing")

	fmt.Println(myOrder.id)
	fmt.Println(myOrder)
	fmt.Println(myOrder2)
	fmt.Println(orderWithFunc)

	// For one time use with in-line constructor
	testStruct := struct {
		id        int
		firstName string
		lastName  string
	}{450, "Vincenzo", "Cassano"}
	fmt.Println(testStruct.id, testStruct)
}

// / order struct
// If we don't set any value, default value is zero value
// zero value: int, float => 0, string => "", bool => false
type order struct {
	id        string
	ammount   float32
	status    string
	createdAt time.Time // nano-second precision
	customer  customer // embedding struct
}
type customer struct {
	name  string
	phone string
}

func newOrder(id string, ammount float32, status string) *order {
	//initial setup goes here....

	myOrder := order{
		id:      id,
		ammount: ammount,
		status:  status,
	}
	return &myOrder
}

func (ord *order) changeStatus(status string) {
	ord.status = status
}
