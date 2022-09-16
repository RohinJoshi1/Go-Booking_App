package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, numTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := numTickets > 0 && numTickets <= RemainingTickets
	return isValidName, isValidEmail, isValidTicket
}
