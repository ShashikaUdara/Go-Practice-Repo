package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// package level variables (Not global) : normaly this is a bad practice
// define variables as much as possible in locally, not in publicaly
const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings = []string // criating a empty slize of string
// var bookings = make([]map[string]string, 0) // criating a empty slize of map
var bookings = make([]UserData, 0) // criating a empty slize of map

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// function level variables
	// can't do this for const values
	// var conferenceName string= "Go Conference" // error because of not using the variable
	// const conferenceTickets uint = 50
	// var remainingTickets uint = 50
	// var bookings []string // or
	// var bookings = []string{} // or
	// bookings := []string[] // this

	// checking types
	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("welcome to %v booking app\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "still available")
	fmt.Println("Get your ticket here")

	// var bookings = [50]string{}
	// var bookings [50]string // this is a simmilar convention for the above statment

	// slices insterd of arrays
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
		isValidUser, isValidEmail, isValidTicketNumber, isValidCity := helper.ValidateUserInput(firstName, lastName, email, userTickets, city, remainingTickets)

		if isValidUser && isValidEmail && isValidTicketNumber && isValidCity {
			// booking ticket
			bookings, remainingTickets = bookTicket(firstName, lastName, userTickets, email)
			go sendTicket(userTickets, firstName, lastName, email)
			// greeting to user
			greetUser(conferenceName)

			// userName = "Udara"
			// fmt.Printf("Thank you %v %v for booking %v tickets. You will revceve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			// fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			// var noTicketsRemaining bool = remainingTickets == 0 // or

			// calling only by first name
			// firstNames [] string
			firstNames := getFirstNames()
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

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}

	// fmt.Printf("First names in Bookings %v\n\n", firstNames)
	return firstNames
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

func bookTicket(firstName string, lastName string, userTickets uint, email string) ([]UserData, uint) {
	// storing user data in a map, used to store similar type of data. (this case strings only)
	// var myMap map[string]string // or
	// var userData = make(map[string]string) // make stands to create the empty map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// we need to cast userTickets to add into userData map
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// insterd of map here we use structure
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will revceve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	// fmt.Printf("\n\nList of bookings is %v\n\n", bookings[0]["firstName"]) // accessing only the first name of the firast element
	fmt.Printf("\n\nList of bookings is %v\n\n", bookings)
	return bookings, remainingTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// adding a time dealy
	time.Sleep(10 * time.Second)
	// creating a formated string
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sendint ticket:\n\t%vTo email address %v\n", ticket, email)
	fmt.Println("#######################")
}
