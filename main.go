package main

import (
	"booking-app/helper"
	"fmt"
)

// Set global variables such at booking name and total tickets.
const conferenceTickets uint = 50

var remainingTickets uint = 50
var conferenceName string = "Go Conference"

var bookings = []UserData{}

type UserData struct {
	firstName, lastName, email string
	numberOfTickets            uint
}

func main() {
	// Greet users to the booking system.
	greetUsers()

	// Get user data.
	firstName, lastName, email, userTickets := getUserData()

	// Validate the input.
	isValidName, isValidEmail, isValidTickets := helper.ValidateData(firstName, lastName, email, userTickets, remainingTickets)

	if !(isValidName && isValidEmail && isValidTickets) {
		fmt.Println("Data is invalid.")
	}

	for {
		// Book the tickets.
		bookTickets(firstName, lastName, email, userTickets)

		// Print booking information.
		firstNames := getFirstNames()
		fmt.Printf("The first names of all the bookings are:\n%v", firstNames)

		// Email the booking information asynchronously.

		// Use switch case for city.

		// Break loop.
		if remainingTickets == 0 {
			break
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend:\n")
}

func getUserData() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Printf("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter email ID: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets you want to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName, lastName, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Println("All bookings:\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v remaining tickets.\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}
