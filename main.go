package main

import (
	"fmt"
	"sync"
	"time"
)

type userData struct {
	firstName  string
	lastName   string
	email      string
	numTickets uint
}

var wg = sync.WaitGroup{}

const totalTickets = 50

var RemainingTickets uint = 50
var bookings = make([]userData, 0)

const conferenceName = "Go conference"

func main() {
	greet()
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		if RemainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}

	wg.Wait()

}
func greet() {
	fmt.Printf("Welcome to our %v booking application", conferenceName)
	fmt.Printf("We have a total of %v tickets, out of which \n %v tickets are left", totalTickets, RemainingTickets)
	fmt.Println("Please book your tickets here!")
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numTickets uint
	fmt.Println("Enter first name,last name, email and tickets required in that order")
	fmt.Scan(&firstName, &lastName, &email, &numTickets)
	return firstName, lastName, email, numTickets
}
func bookTickets(numTickets uint, firstName string, lastname string, email string) {
	RemainingTickets -= numTickets
	var userData = userData{
		firstName:  firstName,
		lastName:   lastname,
		email:      email,
		numTickets: numTickets,
	}
	bookings = append(bookings, userData)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
