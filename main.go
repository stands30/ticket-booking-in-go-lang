package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {
	conferenceName := "Go Conference"
	const confrenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = make([]UserData, 0)

	greetUsers(conferenceName, confrenceTickets, remainingTickets)
	for {
		userData := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userData.firstName, userData.lastName, userData.email, userData.tickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets, bookings := bookTickets(remainingTickets, bookings, userData, conferenceName)
			wg.Add(1)
			go sendTicket(userData)

			firstNames := printFirstNames(bookings)
			fmt.Printf("These are all our bookings: %v", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else if userData.tickets == remainingTickets {
			fmt.Printf("We have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userData.tickets)
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

			fmt.Println("Your input is invalid, try again!")
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		}

	}

}

func greetUsers(conferenceName string, confrenceTickets int, remainingTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() UserData {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	return userData
}

func bookTickets(remainingTickets uint, bookings []UserData, userData UserData, conferenceName string) (uint, []UserData) {
	remainingTickets = remainingTickets - userData.tickets

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", userData.firstName, userData.lastName, userData.tickets, userData.email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userData.tickets, userData.firstName, userData.lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket %v to email %v.\n", ticket, userData.email)
	fmt.Println("###################")
	wg.Done()
}
