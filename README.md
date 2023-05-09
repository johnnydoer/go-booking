# Booking App

This is a booking application written in Go. Users can book tickets for a conference and receive a confirmation email.

## Getting Started

To get started with the application, follow these steps:

1. Clone the repository:

```
git clone https://github.com/username/booking-app.git
```

2. Navigate to the root directory of the application:

```
cd booking-app
```

3. Run the application:

```
go run main.go
```

4. Follow the prompts to book your tickets.

## Features

### Book Tickets

Users can book tickets for the conference by providing their name, email address, and number of tickets they wish to purchase.

### Validation

The application validates user input to ensure that it is valid before booking tickets.

### Remaining Tickets

The application displays the number of remaining tickets after each booking.

### Email Confirmation

Users receive a confirmation email after booking their tickets. The email is sent asynchronously.

## Built With

* Go - The programming language used
* Sync - Package for synchronization primitives
* Time - Package for time-related functions

## Author

* John Doe - john.doe@example.com

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.