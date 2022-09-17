package main

import (
	"fmt"
	"time"

	"strings"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

func main() {

	greetUser()
	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validation(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookingTicket(userTicket, firstName, lastName, email)
			go sendTicket(userTicket, firstName, lastName, email)

			firstNames := printFirstName()
			fmt.Printf("the first name of booking for %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("sold out")
				break
			}

		} else if !isValidName {
			fmt.Println("firstName or lastName is not valid")
		} else if isValidEmail {
			fmt.Println("email address you entered doesnt contain @")
		} else if isValidTicketNumber {
			fmt.Println("number of ticket entered is not valid")
		} else {
			fmt.Println("your input data is invalid! try again")
		}

	}
}
func greetUser() {
	fmt.Printf("conferenceTickets is %T, remainingTicket is %t\n", conferenceTickets, conferenceName, remainingTickets)

	fmt.Printf("welcme to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your ticket here to attend")
}
func printFirstName() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func validation(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("enter your email name")
	fmt.Scan(&email)

	fmt.Println("enter number of ticket")
	fmt.Scan(&userTicket)
	return firstName, lastName, email, userTicket
}

func bookingTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)
	fmt.Printf("thank you %v %v for booking %v tickets. you will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v ticket remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("sending ticket %v to email address %v", ticket, email)
	fmt.Println("##############")
}
