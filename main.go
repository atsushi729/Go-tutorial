package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avairable!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var useTickets uint
	
	firstName = "tom"
	useTickets = 2
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets:  ")
	fmt.Scan(&useTickets)
	
	remainingTickets = remainingTickets - useTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, useTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)
}