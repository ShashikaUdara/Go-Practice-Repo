package main

import (
	"fmt"
	"strings"
)

func main() {
	// can't do this for const values
	conferenceName := "Go Conference" // error because of not using the variable

	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// checking types
	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("welcome to %v booking app\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "still available")
	fmt.Println("Get your ticket here")

	// var bookings = [50]string{}
	// var bookings [50]string // this is a simmilar convention for the above statment

	// slices insterd of arrays
	var bookings []string // or
	// var bookings = []string{} // or
	// bookings := []string[] // this

	// function

	for {
		// getting user inputs
		firstName, lastName, email, city, userTickets := getUserInput()

		switch city {
		case "Colombo", "Panadura":
			// execute code for both colombo and panadura
			fmt.Println("Colombo or Panadura")
		case "Kandy", "Matale":
			// execute code for both colombo and panadura
			fmt.Println("Kandy or Matale")
		default:
			fmt.Println("You enterd a wrong city")
		}

		// validating the user inputs
		isValidUser, isValidEmail, isValidTicketNumber, isValidCity := validateUserInput(firstName, lastName, email, userTickets, remainingTickets, city)

		if isValidUser && isValidEmail && isValidTicketNumber && isValidCity {
			// booking ticket
			bookings, remainingTickets = bookTicket(bookings, firstName, lastName, remainingTickets, userTickets, conferenceName, email)

			// greeting to user
			greetUser(conferenceName)

			// userName = "Udara"
			// fmt.Printf("Thank you %v %v for booking %v tickets. You will revceve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			// fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			// var noTicketsRemaining bool = remainingTickets == 0 // or

			// calling only by first name
			// firstNames [] string
			firstNames := getFirstNames(bookings)
			fmt.Printf("First names in Bookings %v\n\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				// end the program
				fmt.Println("Come next year")

				break
			}
		} else {
			if !isValidUser {
				fmt.Println("The input first name or the last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("The input email is not in the correct order!")
			}
			if !isValidTicketNumber {
				fmt.Printf("You can't buy %v tickets, we have only %v tickets left!\n", userTickets, remainingTickets)
			}
			if !isValidCity {
				fmt.Println("You are in a wrong city!")
			}
			// fmt.Println("Somthing is not right")
			// fmt.Printf("You can't have %v amount of tickets, the number of remaining tickets are %v\n\n", userTickets, remainingTickets)
			continue
		}
	}
}

func greetUser(confName string) {
	fmt.Printf("Welcome to our %v", confName)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	// fmt.Printf("First names in Bookings %v\n\n", firstNames)
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint, city string) (bool, bool, bool, bool) {
	isValidUser := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidCity := city == "Colombo" || city == "Kandy"

	return isValidUser, isValidEmail, isValidTicketNumber, isValidCity
}

func getUserInput() (string, string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var city string
	var userTickets uint
	// ask for user name

	// fmt.Println(conferenceName)  // actual value
	// fmt.Println(&conferenceName) // pointer: address of the variable

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Print("Enter your city : ")
	fmt.Scan(&city)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, city, userTickets
}

func bookTicket(bookings []string, firstName string, lastName string, remainingTickets uint, userTickets uint, confName string, email string) ([]string, uint) {
	bookings = append(bookings, firstName+" "+lastName)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will revceve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, confName)

	return bookings, remainingTickets
}
