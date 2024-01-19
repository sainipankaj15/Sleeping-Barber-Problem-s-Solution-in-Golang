// The sleeping barbar problem
package main

import (

	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variable
var seatingCapcity = 10
var arrivalRate = 100
var cuttingTime = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan("Sleeping Barber Problem Solution using Channel and Goroutine")
	color.Cyan("------------------------------------------------------------")

	clientChan := make(chan string, seatingCapcity) // Buffered channel
	doneChan := make(chan bool)

	shop  := BarberShop{
		ShopCapacity: seatingCapcity,
		HairCutDuration: cuttingTime,
		NumberOfBarbers: 0,
		BarberDoneChan: doneChan,
		ClientsChan: clientChan,
		shopOpen: true,
	}

	color.Green("The Shop is open Now!!")

	shop.AddBarber("Pankaj")

	time.Sleep(4 * time.Second)
}