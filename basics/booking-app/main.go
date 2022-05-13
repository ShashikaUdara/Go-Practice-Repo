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

	for {
		var firstName string
		var lastName string
		var email string
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

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets = remainingTickets - userTickets

		// userName = "Udara"
		fmt.Printf("Thank you %v %v for booking %v tickets. You will revceve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("First names in Bookings %v\n\n", firstNames)
	}
}
