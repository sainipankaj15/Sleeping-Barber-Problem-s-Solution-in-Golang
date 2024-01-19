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

			client, shopOpen := <-barberShop.ClientsChan
			// second variable only will tell , do we get anything from above channel : This is in Genral in channel
			// If second variable (in above case shopOpen) is true then it means we got something from the channel
			// If second variable (in above case shopOpen) is false then it means we didn't get anything from the channel

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = true
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
