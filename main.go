package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var conferenceName string = "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avairable!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 5 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := useTickets > 0 && useTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - useTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, useTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is bookied out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remeining, so you can't book %v tickets.\n", remainingTickets, useTickets)
			continue
		}

	}
}
