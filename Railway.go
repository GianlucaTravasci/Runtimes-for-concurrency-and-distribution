package main

import (
	"fmt"
	"time"
)

type station struct {
	id    int
	name  string
	queue chan traveler
}

type traveler struct {
	id                  int
	source, destination station
	delay               int
}

func fillChannels(travelers [7]traveler, stations [6]station) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	for i := 0; i < len(travelers); i++ {
		for j := 0; j < len(stations); i++ {
			if travelers[i].source.name == stations[j].name {
				stations[j].queue <- travelers[i]
			}
		}
	}
}

func getOff(train chan traveler, stationIn station) {
	auxChan := make(chan traveler)
	traveler := <-train
	if traveler.destination.name == stationIn.name {
		stationIn.queue <- traveler
	} else {
		auxChan <- traveler
	}

	train = auxChan
}

func getIn(train chan traveler, stationOff station) {
	traveler := <-stationOff.queue
	train <- traveler
}

func daylySchedule(train chan traveler, stations [6]station) {
	for {
		for i := 0; i < len(stations); i++ {
			getOff(train, stations[i])
			getIn(train, stations[i])
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	padovaChannel := make(chan traveler, 7)
	pordenoneChannel := make(chan traveler, 7)
	vicenzaChannel := make(chan traveler, 7)
	veneziaChannel := make(chan traveler, 7)
	rovigoChannel := make(chan traveler, 7)
	bassanoChannel := make(chan traveler, 7)

	Padova := station{1, "Padova", padovaChannel}
	Pordenone := station{2, "Pordenone", pordenoneChannel}
	Vicenza := station{3, "Vicenza", vicenzaChannel}
	Venezia := station{4, "Venezia", veneziaChannel}
	Rovigo := station{5, "Rovigo", rovigoChannel}
	Bassano := station{6, "Bassano", bassanoChannel}

	Sergio := traveler{1, Padova, Pordenone, 5}
	Luca := traveler{2, Pordenone, Vicenza, 10}
	Matteo := traveler{3, Padova, Bassano, 11}
	Marco := traveler{4, Vicenza, Venezia, 7}
	Luciana := traveler{5, Rovigo, Venezia, 20}
	Gianni := traveler{6, Venezia, Pordenone, 3}
	Caterina := traveler{7, Bassano, Rovigo, 13}

	stations := [6]station{Padova, Pordenone, Vicenza, Venezia, Rovigo, Bassano}
	travelers := [7]traveler{Sergio, Luca, Matteo, Marco, Luciana, Gianni, Caterina}
	train := make(chan traveler, 3)

	fillChannels(travelers, stations)

	fmt.Println(" Station ID  |  Passengers on Board  |  Passengers in Station ")
	go daylySchedule(train, stations)
}

