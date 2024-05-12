package main

import (
	"fmt"
)

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("The address of thing2 is: %p\n", thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}

func main(){
	var p *int32 = new(int32)
	var i int32
	p = &i
	*p = 10
	fmt.Printf("The value of p to is: %v\n", *p)
	fmt.Printf("The address of p is: %v\n", p)
	fmt.Printf("The value of i is: %v\n", i)
	fmt.Printf("The address of i is: %v\n", &i)

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing1 = [5]float64{1,2,3,4,5}
	var result [5]float64 = square(&thing1)
	fmt.Println(thing1)
	fmt.Println(result)
	fmt.Printf("The address of thing1 is: %p\n", &thing1)
}