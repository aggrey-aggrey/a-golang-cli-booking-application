package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets(uint)  = 50
 var bookings = []string{}

func main() {

	for {

		greetUsers()

        firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail  && isValidTicketNumber{

 
			//bookTicket function
			bookTicket( userTickets, bookings, firstName,  lastName, email )

            firstNames := getFirstNames()
			fmt.Printf("The first names of booking are : %v\n", firstNames)
			

			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year.")

				break
			}
		
		} else {

			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail{
				fmt.Println("email address you entered doesn't contain  that c @")
			}

			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
			
		}

	}

}


func greetUsers (){
	
	fmt.Printf("Welcome to our  %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend ")
}


func getFirstNames () [] string{

	firstNames := []string {}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
 return firstNames

}


func validateUserInput (firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){

	isValidName := len(firstName) >= 2 && len(lastName) >=2
		isValidEmail := strings.Contains(email, "@")
		
		isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)

		return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput () (string, string, string, uint){

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:  ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}


func bookTicket ( userTickets uint, bookings [] string, firstName string, lastName string, email string){

	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking tickets. You will receive a confirmation email at %v and %v %v are still available.\n",firstName,lastName, email, remainingTickets, conferenceName)
	fmt.Printf("%v tickets  remaining  for this %v", remainingTickets, conferenceName)
	fmt.Printf("The list of %v", bookings)
}