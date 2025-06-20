package main

import "fmt"

// Entry point of the application
func main() {
	// -------------------------------
	// Variables and Constants Section
	// -------------------------------
	conferenceTickets := 50           // Total number of tickets (constant for this run)
	remainingTickets := 50            // Tickets left for booking (will decrease)
	conferenceName := "Go conference" // Name of the conference

	// -------------------------------
	// Data Structures: Slice of Structs
	// -------------------------------
	bookings := []Booking{} // Stores all bookings as a slice of Booking structs

	// -------------------------------
	// Welcome Message (Functions)
	// -------------------------------
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// -------------------------------
	// Concurrency: Channel for Confirmations
	// -------------------------------
	confirmations := make(chan string) // Channel to receive confirmation messages from goroutines

	// -------------------------------
	// Main Booking Loop (Control Flow: for loop)
	// -------------------------------
	for remainingTickets > 0 {
		// -------------------------------
		// User Input (Functions, Input/Output)
		// -------------------------------
		userName, userTickets := getUserInput()

		// -------------------------------
		// Input Validation (if/else)
		// -------------------------------
		if userTickets > remainingTickets {
			printNotEnoughTickets(userName, remainingTickets)
		} else if userTickets <= 0 {
			printInvalidTicketNumber()
		} else {
			// -------------------------------
			// Booking Logic (Functions, Structs)
			// -------------------------------
			remainingTickets = bookTickets(userName, userTickets, remainingTickets)
			booking := Booking{userName, userTickets}
			bookings = append(bookings, booking)
			printBookings(bookings)

			// -------------------------------
			// File I/O: Save Booking to File
			// -------------------------------
			err := saveBookingToFile(booking)
			if err != nil {
				fmt.Println("Error saving booking:", err)
			}

			// -------------------------------
			// Concurrency: Send Confirmation in Background
			// -------------------------------
			go sendConfirmation(userName, userTickets, confirmations)
		}

		// -------------------------------
		// Sold Out Check
		// -------------------------------
		if remainingTickets == 0 {
			printSoldOut()
		}

		// -------------------------------
		// Non-blocking Receive: Print Confirmation if Ready
		// -------------------------------
		select {
		case msg := <-confirmations:
			fmt.Println(msg)
		default:
			// No confirmation ready yet; continue
		}
	}
}
