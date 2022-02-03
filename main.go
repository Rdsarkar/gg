package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	// fmt.Println(conferenceName)
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v avaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var fristName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their nae
	fmt.Println("Enter your Frist Name: ")
	fmt.Scan(&fristName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of your tickets: ")
	fmt.Scan(&userTickets)

	// userName = "Tom"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets.\n", fristName, userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", fristName, lastName, userTickets, email)
}
