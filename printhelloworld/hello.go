package Printhelloworld

import (
	"errors"
	"fmt"
)

func Hello(user string) (string , error) {
	if len(user) >= 10 {
		return "",errors.New("name user must less than ten")
	}
	fmt.Printf("Hello %s\n", user)
	return "Hello " + user , nil
}