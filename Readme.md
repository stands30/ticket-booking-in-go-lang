## Ticket Booking System in Go

### Overview
This project implements a simple ticket booking system using the Go programming language. It allows users to book tickets for a conference, with basic validation and a concurrent ticket sending mechanism.

### Features
* User input validation for name, email, and ticket quantity.
* Ticket booking with remaining ticket count management.
* Concurrent ticket sending to users via goroutines.
* Basic user data storage in a slice.

### Getting Started
#### Prerequisites
* Go (version 1.18 or later)

#### Running the Application
1. Clone the repository:
   ```bash
   git clone https://github.com/stands30/ticket-booking-in-go-lang.git
   ```
2. Navigate to the project directory:
   ```bash
   cd ticket-booking-in-go-lang
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

### How it works
* The application initializes with a conference name, total tickets, and remaining tickets.
* It continuously prompts users for their name, email, and desired number of tickets.
* Input validation checks for name length, email format, and ticket availability.
* If valid, tickets are booked, and a goroutine is started to simulate sending a ticket confirmation email.
* The remaining ticket count is updated, and the process repeats until tickets are sold out.

### Limitations
* This is a basic implementation without persistent storage or error handling.
* The ticket sending mechanism is simulated and doesn't actually send emails.
* The user interface is basic command-line based.

### Potential Improvements
* Implement a database to store user and booking information.
* Add error handling for invalid inputs or unexpected conditions.
* Improve user interface with a web-based or GUI application.
* Enhance ticket sending with a real email service integration.
* Introduce features like ticket cancellations, refunds, and waiting lists.
* Implement unit and integration tests for code reliability.

### Contributing
Contributions are welcome! Please fork the repository and submit pull requests.