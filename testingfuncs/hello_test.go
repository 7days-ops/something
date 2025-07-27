package testingfuncs

import (
	"fmt"
	h "mymodule/printhelloworld"
	"testing"
)

func TestHello(t *testing.T) {
	exp := "Hello Kirill"
	act , err := h.Hello("Kirill")
	if err != nil {
		fmt.Printf("some err: %v" , err)
	}
	if exp != act {
		t.Error()
	}
}