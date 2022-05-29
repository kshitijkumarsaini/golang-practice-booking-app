package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// defining array
	// var bookings [50]string
	// defining slices (Similar to array but has more flexibility)
	bookings := []string{}

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Println("Get your tickets here to attend", conferenceName)

	// fmt.Printf("conferenceName is of %T type,\n conferenceTickets is of %T type,\n remainingTickets is of %T type\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend", conferenceName)

	// this prints the value
	// fmt.Println(remainingTickets)
	// this prints the address. Here '&' is the pointer which points to the memory address of the variable
	// fmt.Println(&remainingTickets)

	// Arrays in go needs the size of array
	// var bookings = [50]string{}

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their firstName
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)

	// ask user for their lastName
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)

	// ask user for their email
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	// ask user for no. of tickets he/she needs to purchase
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v Tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings so far %v.", bookings)
}
