package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	shopOpen        bool
}

func (barberShop *BarberShop) AddBarber(barber string) {

	barberShop.NumberOfBarbers++
	fmt.Println(barber, "is in the house and ready to cut the hairs for customors")

	go func() {
		isSleeping := false
		color.Yellow("%s is waiting for the customoers", barber)

		for {

			// Very first we will check , is there any client who is waiting in the queue if not then barber will go to sleep
			if len(barberShop.ClientsChan) == 0 {
				isSleeping = true
				color.Yellow("There was nothing to do so %s went to sleep", barber)
			}

			client, shopOpen := <- barberShop.ClientsChan
			// second variable only will tell , do we get anything from above channel : This is in Genral in channel
			// If second variable (in above case shopOpen) is true then it means we got something from the channel
			// If second variable (in above case shopOpen) is false then it means we didn't get anything from the channel

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}

				// Now Barber will cut the hair of the customor
				barberShop.cutHair(barber, client)
			} else {
				// Shop is closed , barber is going to home
				barberShop.barberGoesHome(barber)
				return
			}
		}
	}()
}

func (barberShop *BarberShop) cutHair(barber string, client string) {
	color.Green("%s is cuttign the hair of %s", barber, client)
	time.Sleep(barberShop.HairCutDuration)
	color.Green("%s finished the cutting of %s hair", barber, client)
}

func (barberShop *BarberShop) barberGoesHome(barber string) {
	color.Blue("%s is going home", barber)
	barberShop.BarberDoneChan <- true
}

func (barberShop *BarberShop) CloseShopForTheDay() {
	color.Red("Shop is going to close day and Not accepting any new client")

	// So basiclly now we have to close the channel of client
	close(barberShop.ClientsChan)
	barberShop.shopOpen = false

	//Now here we will wait to wait to go home for all barbers

	for i := 1; i <= barberShop.NumberOfBarbers; i++ {
		// This code will block until and unless something comes on BarberDoneChannel
		<-barberShop.BarberDoneChan
	}

	close(barberShop.BarberDoneChan)
	color.Cyan("--------------------------------------------------------------")
	color.Green("The barber Shop is now closed and all barber has gone to home")
	color.Cyan("--------------------------------------------------------------")
}

func (barberShop *BarberShop) AddClient(clientName string) {
	
	color.Green("%s has come in the shop" , clientName)

	if barberShop.shopOpen {
		select {
		case barberShop.ClientsChan <- clientName :
			color.Blue("%s takes a seat in a waiting room" , clientName)
		default :
			color.Red("The waiting room is full , so %s leaves", clientName)
		}
	}else{
		color.Red("Shop is closed so you can go Mr. %s ", clientName)
	}

}