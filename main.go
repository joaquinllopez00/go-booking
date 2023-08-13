package main

import "fmt"

func main() {
		var conferenceName string = "Go Conference"
		const conferenceTickets int = 50
		var remainingTickets int = 50
		fmt.Println("Welcome to the go booking application for the", conferenceName)
		fmt.Println("There are", conferenceTickets, " total. There are", remainingTickets, "tickets remaining for the", conferenceName)
		fmt.Println("Booking isn't free but it's worth it")


	}