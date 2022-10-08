package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets(uint)  = 50
 var bookings = make([]UserData,0)
 var remainingTkt uint = 0

 type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
 }

func main() {

	for {

		greetUsers()

        firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail  && isValidTicketNumber{

			remainingTkt = bookTicket(userTickets, firstName,  lastName, email )
			sendTicket(userTickets, firstName, lastName, email)

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
				fmt.Println("email address you entered doesn't contain  that character @")
			}

			if !isValidTicketNumber{
				fmt.Printf("number of tickets you entered is invalid, There are %v remaining, but you are requesting %v ticket(s)\n", remainingTkt, userTickets )
			}
			
		}

	}

}


func greetUsers (){
	
	fmt.Printf("Welcome to our  %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend\n ")
}


func getFirstNames () [] string{

	firstNames := []string {}
			for _, booking := range bookings {
				
				firstNames = append(firstNames, booking.firstName)

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

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets :  ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}


func bookTicket ( userTickets uint, firstName string, lastName string, email string) uint {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,

	}


	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking tickets. You will receive a confirmation email at %v and %v %v are still available.\n",firstName,lastName, email, remainingTickets, conferenceName)
	fmt.Printf("%v tickets  remaining  for this %v\n", remainingTickets, conferenceName)
	fmt.Printf("The list of %v\n", bookings)

	return remainingTickets
}


func sendTicket(userTickets uint, firstName string, lastName string, email string){
	
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending tickets: %v to email address %v\n", ticket, email)
	fmt.Printf("########################\n")
}