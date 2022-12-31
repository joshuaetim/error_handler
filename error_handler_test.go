package error_handler

import (
	"fmt"
	"testing"
)

type Name struct {
	Value string
}

func TestDeferError(t *testing.T) {
	defer HandleError()
	ok := false
	GetOrders()
	if ok = true; ok {
		t.Error("this part of code should be unreachable")
	}
}

func DivideByZero() {
	num := 0
	main := 4 / num
	fmt.Println(main)
}

func GetOrders() {
	var name *Name
	fmt.Println(name.Value)
}

func MockIndexError() {
	fmt.Println([]string{"cool"}[4])
}
