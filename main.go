package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	remainingTickets = -1
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avairable!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	var userName string
	var useTickets int
	// ask user ofr their name

	userName = "tom"
	useTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, useTickets)
}