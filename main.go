package main

import (
	"fmt"
	"strings"
)

func main() {
	// define variable/const and print to standard output
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string // slice (dynamic array)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		// taking user input
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thanks %v %v for ordering %v tickets. You will receive a confirmation email at %v.\n",
				firstName, lastName, userTickets, email)

			remainingTickets = remainingTickets - userTickets
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) // split by whitespace
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all the first name: %v\n", firstNames)
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n",
				remainingTickets, userTickets)
		}
	}
}
