package auth

import (
	"fmt"
)

// method name should start with Capital letter to access anywhere

func LoginWithCredential(username string, password string) {
	fmt.Println("Login user using", username, password)
}
