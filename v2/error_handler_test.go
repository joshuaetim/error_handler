package error_handler_test

import (
	"fmt"
	"testing"

	"github.com/joshuaetim/error_handler/v2"
)

type Name struct {
	Value string
}

func TestDeferError(t *testing.T) {
	handler := error_handler.NewErrorHandler(&error_handler.Config{
		PreferredColor: error_handler.ColorRed,
	})
	defer handler.HandleError()
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
