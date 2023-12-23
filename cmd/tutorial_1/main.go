package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1,2,3} // [0, 0, 0]
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8 {"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("This profile is not valid!\n")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}
