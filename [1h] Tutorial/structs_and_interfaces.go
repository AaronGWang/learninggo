package main

import (
	"fmt"
)

type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
	owner
}

type owner struct{
	name string
}

type engine interface{
	milesLeft() uint8
}


func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles<=e.milesLeft(){
		fmt.Println("You can make it!")
	}else{
		fmt.Println("You can't make it!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 30, gallons: 10, owner: owner{name: "Aaron"}}
	myEngine.mpg = 35
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner.name)

	var myEgine2 = struct{
		mpg uint8
		gallons uint8
		owner
	}{25, 15, owner{"Sarah"}}
	fmt.Println(myEgine2.mpg, myEgine2.gallons, myEgine2.owner.name)

	fmt.Printf("Total miles left on %v's engine: %v\n", myEngine.owner.name, myEngine.milesLeft())

	var myElectricEngine electricEngine = electricEngine{mpkwh: 4, kwh: 20, owner: owner{name: "Jason"}}
	canMakeIt(myElectricEngine, 81)
}