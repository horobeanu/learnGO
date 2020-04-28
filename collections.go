package main

import (
	"fmt"
)

func main() {
	// ### 1. Array - fixed sized collection of similar types
	// long syntax
	var arr [3]int
	fmt.Println(arr)
	//[0 0 0]
	arr[0] = 65
	arr[1] = 5
	arr[2] = 45
	fmt.Println(arr)
	//[65 5 45]

	// implicit initialization and initialization
	arr2 := [3]int{3, 5, 7}
	fmt.Println(arr2)
	//[3 5 7]

	// ### 2. Slice - pointing to the original value arr[startIndex:endIndex]
	slice := arr2[:]
	slice[2] = 99
	fmt.Println(arr2)
	//[3 5 99]

	// ### 2. Slice - dynamic size
	sliceArrDynamicSize := []int{}
	sliceArrDynamicSize = append(sliceArrDynamicSize, 9, 7)
	fmt.Println(sliceArrDynamicSize)
	//[9 7]

	// ### 3. Map (not safe for concurrent use - https://blog.golang.org/maps)
	m := map[string]int{"key_1": 20}
	fmt.Println(m["key_1"])

	// eliminate a key from map
	delete(m, "key_1")
	// add a key to map
	m["key_7"] = 90

	// Iterate over maps (the iteration order is not specified and is not guaranteed to be the same from one iteration to the next.)
	for key, value := range m {
		fmt.Printf("key:%v value:%v \n", key, value)
	}

	// ### 3. Struct
	type User struct {
		ID int // starts with capital letter => public
		// 0 default value for int
		password string // starts with low letter => private
		// empty string default value for strings
		FirstName string
	}
	var u User
	fmt.Println(u)
	// {0  }

	// create a slice of pointers to User type variable
	var users []*User
	// a new User pointer to the slice
	users = append(users, &u)
}
