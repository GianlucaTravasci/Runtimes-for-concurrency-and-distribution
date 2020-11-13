package main

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

func scendere(train chan traveler, stazione chan traveler) {
	//TODO
}

func salire(train chan traveler, stazione chan traveler) {
	//TODO
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

	fillChannels(travelers, stations)

}
