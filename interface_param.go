package main

import (
	"fmt"
)

func checkType(param interface{}) {
	// Type switch to check the underlying type of param
	switch v := param.(type) {
	default:
		fmt.Printf("Unexpected type %T\n", v)
	case uint64:
		fmt.Println("Integer type")
	case string:
		fmt.Println("String type")
	}
}

func main() {
	// Testing with different types
	var a uint64 = 100
	var b string = "Hello"
	var c float64 = 99.99

	checkType(a) // Should print "Integer type"
	checkType(b) // Should print "String type"
	checkType(c) // Should print "Unexpected type float64"
}
