package main

import ("fmt"
		
		"booking-app/helper"
		"strconv"

)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets int = 50
var bookings = make([]map [string]string, 0) //initializing an empty list of maps

func main(){

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber:= helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else { 
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ ")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid ")
			}
				continue
		}

	}
	
}


func greetUsers() {
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining. \n", conferenceTickets, remainingTickets)
	fmt.Println ("Get your tickets here to attend")
}

func getFirstNames() [] string {
	firstNames :=[]string{}
			for _, booking :=range bookings { 
				firstNames = append(firstNames, booking["firstName"])
			}
			return firstNames
}

func getUserInput()(string, string, string, int) {
	var firstName string
		var lastName string
		var email string
		var userTickets int
		
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets) 

		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	
	// create a map for a user

	var userData = make(map [string]string)
	userData ["firstName"] = firstName
	userData ["last Name"] = lastName
	userData ["email"] = email
	userData ["number of Tickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)
			
	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v. \n", remainingTickets, conferenceName)
}