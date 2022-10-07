package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets(uint)  = 50
 var bookings = make([]map[string]string,0)

func main() {

	for {

		greetUsers()

        firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

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
				
				firstNames = append(firstNames, booking["firstName"])

			}
 return firstNames

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


func bookTicket ( userTickets uint, bookings []map[string]string, firstName string, lastName string, email string){

	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)


	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking tickets. You will receive a confirmation email at %v and %v %v are still available.\n",firstName,lastName, email, remainingTickets, conferenceName)
	fmt.Printf("%v tickets  remaining  for this %v", remainingTickets, conferenceName)
	fmt.Printf("The list of %v", bookings)
}