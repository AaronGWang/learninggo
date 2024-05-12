package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr[0])
	fmt.Println(&intArr[0])

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length of intSlice is %v, the capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length of intSlice is %v, the capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length of intSlice is %v, the capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 5, 8) // make([]type, length, capacity)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])
	var age, ok = myMap2["Sarah"]
	if ok{
		fmt.Println(age)
	}else{
		fmt.Println("Not in the map")
	}

	for name, age := range myMap2{
		fmt.Printf("Name: %v | Age: %v\n", name, age)
	}

	for i:=0; i<10; i++{
		fmt.Println(i)
	}
}