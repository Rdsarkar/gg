package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	// fmt.Println(conferenceName)
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v avaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{}
	var bookings [50]string

	var fristName string
	var lastName string
	var email string
	var userTickets uint
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

	remainingTickets = remainingTickets - userTickets
	bookings[0] = fristName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The frist value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fristName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
