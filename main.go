package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	// fmt.Println(conferenceName)
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v avaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// var bookings = [50]string{}
		// var bookings [50]string   //Array
		var bookings []string //Slice
		// var bookings =[]string{}    //Slice Alternative
		// bookings := []string{}     //Slice Alternative

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

		isValidName := len(fristName) >= 2 && len(lastName) >=2

		isValidEmail := strings.Contains(email, "@")

		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		
		
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = fristName + " " + lastName						//Array
		bookings = append(bookings, fristName+" "+lastName) //Slice

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The frist value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fristName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of boolings are: %v\n", firstNames)

		if  !isValidName && !isValidEmail && !isValidTicketNumber{
			fmt.Println("Your input data is invalid, try again")
			break
		}

		if remainingTickets == 0{
			fmt.Println("Our conference is booked out, Come back next year.")
			break
		}
	}

}


//1:33:50