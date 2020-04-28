package main

import "fmt"

func main() {
	fmt.Printf(
		"Sum : %v",
		multipleParamsExample(1, 1),
	)

	sum, isGreater := multipleReturnExample(1, 1)
	// ignore second returned value
	sum2, _ := multipleReturnExample(1, 1)
	fmt.Println(sum, isGreater, sum2)
	// Sum : 22 false 2
}

func multipleParamsExample(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

// This method returns the sum of "f" and "s" as first returned value
// and true if f > s on the second returned value.
// Multiple params same type can be defined in a short way
func multipleReturnExample(f, s int) (int, bool) {
	isGreater := false
	if f > s {
		isGreater = true
	}
	return f + s, isGreater
}