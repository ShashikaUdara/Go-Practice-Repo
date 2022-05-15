package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, city string) (bool, bool, bool, bool) {
	isValidUser := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidCity := city == "Colombo" || city == "Kandy"

	return isValidUser, isValidEmail, isValidTicketNumber, isValidCity
}
