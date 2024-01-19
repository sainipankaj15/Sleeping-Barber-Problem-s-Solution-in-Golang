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

## Historical Context:

The Sleeping Barber problem was originally proposed in 1965 by computer science pioneer Edsger Dijkstra. The problem and its solution aim to emphasize that, in many cases, the use of semaphores (mutexes) may not be necessary, showcasing the elegance of its resolution.


