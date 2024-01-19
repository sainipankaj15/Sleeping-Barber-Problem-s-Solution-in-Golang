// The sleeping barbar problem
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapcity = 3 // this is the no of total capicaity of the shop : Seating capacity
var arrivalRate = 100 // this time is for how much time a client will take to come : indirectly
var cuttingTime = 1000 * time.Millisecond // this time is for how much time a barber takes to cut the hair
var timeOpen = 5 * time.Second // this time is for how much time shop will be open

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan("Sleeping Barber Problem Solution using Channel and Goroutine")
	color.Cyan("------------------------------------------------------------")

	clientChan := make(chan string, seatingCapcity) // Buffered channel
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapcity,
		HairCutDuration: cuttingTime,
		NumberOfBarbers: 0,
		BarberDoneChan:  doneChan,
		ClientsChan:     clientChan,
		shopOpen:        true,
	}

	color.Green("The Shop is open Now!!")

	// Adding Barber in the Shop
	shop.AddBarber("Pankaj")

	// Run barber shop in a go rotuine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		// Sleep this goroutine : For the timings of shop opens
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForTheDay()
		closed <- true

	}()
	
	// Now we will add client : This will be also in go routine 
	clientNo := 1

	go func(){
		for {
			// taking an random number 
			randMillSeconds := rand.Int() % (2 *arrivalRate)
			fmt.Println("Random no is ", randMillSeconds)
			select{
			case <- shopClosing :
				return 
			case <- time.After(time.Millisecond * time.Duration(randMillSeconds)):
				clientName := fmt.Sprintf("Client%d",clientNo)
				clientNo++
				shop.AddClient(clientName)
			}
		}
	}()

	<- closed
	close(closed)
}
