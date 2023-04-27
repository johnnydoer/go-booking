package main

import (
	"booking-app/helper"
	"fmt"
)

// Set global variables such at booking name and total tickets.
const conferenceTickets uint = 50

var remainingTickets uint = 50
var conferenceName string = "Go Conference"

func main() {
	// Greet users to the booking system.
	greetUsers()

	// Get user data.
	firstName, lastName, email, userTickets := getUserData()

	// Validate the input.
	isValidName, isValidEmail, isValidTickets := helper.ValidateData(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {
		fmt.Println("Data is valid.")
	}

	// Book the tickets

	// Print booking information.

	// Email the booking information asynchronously.

	// Use switch case for city,
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend:\n", conferenceName, conferenceTickets, remainingTickets)
}

func getUserData() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Printf("Enter first name: ")
	fmt.Scanln(&firstName)

	fmt.Printf("Enter last name: ")
	fmt.Scanln(&lastName)

	fmt.Printf("Enter email ID: ")
	fmt.Scanln(&email)

	fmt.Printf("Enter number of tickets you want to purchase: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}
