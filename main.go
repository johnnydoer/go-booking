package main

import "fmt"

// Set global variables such at booking name and total tickets.
const conferenceTickets uint = 50

var remainingTickets uint = 50
var conferenceName string = "Go Conference"

func main() {
	// Greet users to the booking system.
	greetUsers()

	// Get user data.

	// Validate the input.

	// Book the tickets

	// Print booking information.

	// Email the booking information asynchronously.

	// Use switch case for city,
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend:\n", conferenceName, conferenceTickets, remainingTickets)
}

// func getUserData() {
// 	var firstName, lastName, email string
// 	var userTickets uint

// 	fmt.Printf()
// }
