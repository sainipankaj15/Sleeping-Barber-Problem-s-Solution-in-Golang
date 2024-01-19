# Sleeping Barber Problem

The Sleeping Barber problem is a classic computer science dilemma that highlights the complexities arising in scenarios with multiple operating system processes. In this scenario, there exists a finite number of barbers, a limited number of seats in a waiting room, a fixed duration for the barbershop to remain open, and clients arriving at roughly regular intervals. The problem revolves around the actions of barbers and clients in a coordinated fashion.

## Rules and Conditions:

- **Initial State:**
  - If there are no customers, the barber falls asleep in the chair.
  - A customer must wake the barber if he is asleep.

- **Customer Arrivals:**
  - If a customer arrives while the barber is working, the customer:
    - Leaves if all chairs are occupied.
    - Sits in an empty chair if available.

- **Haircut Process:**
  - When the barber finishes a haircut, he checks the waiting room for waiting customers.
  - If there are no waiting customers, the barber falls asleep.

- **Closing Time:**
  - The shop can stop accepting new clients at closing time.
  - Barbers cannot leave until the waiting room is empty.

- **End of the Day:**
  - After the shop is closed and there are no clients left in the waiting area, the barber goes home.

# How to approch problem 
- The program is a classic barber shop scenario where customers and barbers interact. Concurrency is managed using Go channels.

## Key Features

### Goroutines and Channels
The program utilizes Goroutines to represent customers, barbers, and other entities within the simulated barber shop. Communication and synchronization are achieved through Go channels.

### Deadlock Prevention
The implementation is designed to avoid deadlocks by carefully managing channel operations and ensuring proper synchronization.

### Race Condition Detection
The program is developed with consideration for race conditions, and the Go race detector (`go run -race`) can be used to identify and address potential issues.


## How to run 
- Clone the repo.
- Run the command ```go mod tidy ```
- After all module successfully installed , run the command ```go run .```
- Output will command on the terminal with appropriate colors.



