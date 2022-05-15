package helper

import "strings"

// we can make the function global by changing the function name first letter
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, city string, remainingTickets uint) (bool, bool, bool, bool) {
	isValidUser := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidCity := city == "Colombo" || city == "Kandy"

	return isValidUser, isValidEmail, isValidTicketNumber, isValidCity
}
