package main

import (
	"fmt"
	"errors"
)


func main(){
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	
	switch{
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("Result: %v", result)
	default:
		fmt.Printf("Result: %v | Remainder: %v", result, remainder)
	}
}


func printMe(printValue string){
	fmt.Println(printValue)
}


func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Division by zero is not allowed\n")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}