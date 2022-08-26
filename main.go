package main

import "fmt"

func main() {
	// define variable/const and print to standard output
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	// taking user input
	var firstName string
	var lastName string
	var email string
	var tickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you need?")
	fmt.Scan(&tickets)

	fmt.Printf("Thanks %v %v for ordering %v tickets. You will receive your tickets in %v.\n",
		firstName, lastName, tickets, email)
}
