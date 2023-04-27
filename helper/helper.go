package helper

import "strings"

func ValidateData(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	var isValidName, isValidEmail, isValidTickets bool

	isValidName = len(firstName) > 2 && len(lastName) > 2
	isValidEmail = strings.Contains(email, "@")
	isValidTickets = remainingTickets >= userTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTickets
}
