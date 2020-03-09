// navigate to this folder and execute "go run main.go"
package main

import (
	"fmt"
)
// package LEVEL - SCOPE
const (
	piPackage = 2.15
	someOtherConst = "Cannot change this on runtime"
	// start from 0 and incr +1 on each use
	autoIncrConst = iota
	// second use of iota => iota = 1
	autoIncrCont2 = iota + 6
)

// second const block
const (
	autoIncrConstSecondBlock = iota
	autoIncrConstSecondBlock2
	autoIncrConstSecondBlock3
)

func main()  {
	// ### 1. Constants
	fmt.Println(piPackage, someOtherConst, autoIncrConst, autoIncrCont2)
	//2.15 Cannot change this on runtime 2 9
	fmt.Println(autoIncrConstSecondBlock, autoIncrConstSecondBlock2, autoIncrConstSecondBlock3)
	//0 1 2

	// implicit definition of a constant in main() SCOPE
	const pi = 3.1415
	fmt.Println(pi)
	//3.1415
	// explicit definition of a constant
	const piExplicit float32 = 3.1415

	// ### 2. verbose declaration
	var intTypeVariable int
	intTypeVariable = 42
	fmt.Println(intTypeVariable)
	//42

	var float32Variable float32 = 3.14
	fmt.Println(float32Variable)
	//3.14

	// ### 3. implicit declaration - compiler assign data type
	lastName := "Horobeanu"
	fmt.Println(lastName)
	//Horobeanu

	complexNumber := complex(1, 8)
	fmt.Println(complexNumber)
	//(1+8i)

	realPart, imaginaryPart := real(complexNumber), imag(complexNumber)
	fmt.Println(realPart, imaginaryPart)
	//1 8

	// ### 4. Pointers
	// pointer to nil
	var firstName *string
	fmt.Println(firstName)
	//<nil>

	// pointer to allocated memory address
	var firstName2 *string = new(string)
	// dereferencing => set the value not the address memory
	*firstName2 = "Daniel"
	fmt.Println(firstName2)
	//0xc000010210
	// this is memory address pointer value
	// to print the value we have to dereference the pointer
	fmt.Println(*firstName2)
	//Daniel

	// get a pointer to a variable eg. float32Variable defined above
	pointerToFloat32Variable := &float32Variable
	fmt.Println(float32Variable, *pointerToFloat32Variable)
	//3.14 3.14

	float32Variable = 0.11
	fmt.Println(float32Variable, *pointerToFloat32Variable)
	//0.11 0.11

}