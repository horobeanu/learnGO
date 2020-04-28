// navigate to this folder and execute "go run data_types.go"

// The package “main” tells the Go compiler that the package
// should compile as an executable program instead of a shared library.
// The main function in the package “main” will be the entry point of our executable program.
// When you build shared libraries, you will not have any main package and main function in the package.

// The naming convention for Go package is to use
// the name of the system directory where we are putting our Go source files.
// Within a single folder, the package name will be same for the all source files which belong to that directory.

// To download and install third-party Go packages use "go get" command eg. "go get gopkg.in/mgo.v2"

// When you write Go packages, you can provide a function “init”
// that will be called at the beginning of the execution time.
// The init method can be used for adding initialization logic into the package.

// If we imported a package and are not using the package identifier in the program,
// Go compiler will show an error. In such a situation, we can use a blank identifier ( _ ) as the package alias name,
// so the compiler ignores the error of not using the package identifier, but will still invoke the init function.

// Build Go package for reusing with other packages using "go install" command
package main

// import fmt package
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

	// string is alias on []byte
	fmt.Println([]byte(lastName))
	//[72 111 114 111 98 101 97 110 117]

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