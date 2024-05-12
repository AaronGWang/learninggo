package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var username string
	var userTickets int
	// ask user for their name

	username = "Tom"
	userTickets = 2
	fmt.Printf("User %v has booked %v tickets\n", username, userTickets)
}