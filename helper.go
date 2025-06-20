package main

import (
	"fmt"
	"os"
	"time"
)

// Booking struct: groups user booking info (structs, data structures)
type Booking struct {
	userName    string
	userTickets int
}

// greetUsers: prints welcome message (functions, input/output)
func greetUsers(conferenceName string, conferenceTickets, remainingTickets int) {
	fmt.Println("Welcome to our booking app!")
	fmt.Println("Please follow the instructions to book your appointment.")
	fmt.Println("Get ready for", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and", remainingTickets, "tickets available.")
}

// getUserInput: gets user input and handles input errors (functions, error handling, input/output)
func getUserInput() (string, int) {
	var userName string
	var userTickets int
	fmt.Print("Enter your name: ")
	_, err1 := fmt.Scan(&userName)
	if err1 != nil {
		fmt.Println("Error reading name. Please try again.")
		return "", 0
	}

	fmt.Print("Enter number of tickets: ")
	_, err2 := fmt.Scan(&userTickets)
	if err2 != nil {
		fmt.Println("Error reading ticket number. Please enter a valid number.")
		return userName, 0
	}
	return userName, userTickets
}

// bookTickets: processes booking and returns updated tickets (functions, logic)
func bookTickets(userName string, userTickets, remainingTickets int) int {
	remainingTickets -= userTickets
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email.\n", userName, userTickets)
	fmt.Printf("We have %v tickets remaining.\n", remainingTickets)
	return remainingTickets
}

// printBookings: prints all current bookings (functions, slices, structs)
func printBookings(bookings []Booking) {
	fmt.Println("Current bookings:")
	for _, b := range bookings {
		fmt.Printf("- %v booked %v tickets\n", b.userName, b.userTickets)
	}
}

// printNotEnoughTickets: error message for overbooking (functions, error handling)
func printNotEnoughTickets(userName string, remainingTickets int) {
	fmt.Printf("Sorry %v, we only have %v tickets left. Please try again.\n", userName, remainingTickets)
}

// printInvalidTicketNumber: error message for invalid ticket count (functions, error handling)
func printInvalidTicketNumber() {
	fmt.Println("Number of tickets must be at least 1.")
}

// printSoldOut: message when tickets are sold out (functions, control flow)
func printSoldOut() {
	fmt.Println("Sorry, we are sold out!")
}

// sendConfirmation: simulates sending confirmation in background (concurrency, goroutines, channels)
func sendConfirmation(userName string, userTickets int, channel chan string) {
	time.Sleep(2 * time.Second) // simulate delay
	confirmation := fmt.Sprintf("Confirmation sent to %v for %v tickets.", userName, userTickets)
	channel <- confirmation
}

// saveBookingToFile: saves booking info to a file (file I/O, error handling)
func saveBookingToFile(booking Booking) error {
	file, err := os.OpenFile("bookings.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	entry := fmt.Sprintf("%v booked %v tickets at %v\n", booking.userName, booking.userTickets, time.Now().Format(time.RFC1123))
	_, err = file.WriteString(entry)
	return err
}
