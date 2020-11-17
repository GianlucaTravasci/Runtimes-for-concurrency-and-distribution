/**
@Title Railway Station - Exercise
@authors: Pozzan Paolo, Salvadore Nicola, Sciacco Mariano, Travasci Gianluca
@Description:
"
	Realize a circular line metro service simulator

	M > 1 train stations (along a circular line)
	N > M commuters who forever revolve around their duty cycle
	1 commuter train with capacity C < N (no prebooking)

	Commuter duty cycle is the following:

	1. Home --> Nearest train station
	2. --> Jump on first possible train
	3. Train --> Work (working..)
	4. Work --> Nearest train station
	5. --> Jump on first possible train
	6. Train --> Home (rest..)
"
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Stazione struct {
	nome string            // Id Stazione
	coda chan *Viaggiatore // Coda dei viaggiatori in attesa
}

type Viaggiatore struct {
	nome             string    // Nome del viaggiatore
	partenza, arrivo string    // Nomi delle stazioni di partenza e arrivo
	attesa           int       // EXTRA: Attesa per lavoro e per casa
	notifica         chan bool // canale di notifica per salita/discesa
	colore           string    // EXTRA: solo per il terminale
}

type Treno struct {
	capienza int
	posti    []*Viaggiatore
}

func RoutineViaggiatore(persona *Viaggiatore, stazioni []*Stazione) {
	for {
		// Trova la stazione dove metterti in coda
		p := 0
		for i := 0; i < len(stazioni); i++ {
			if stazioni[i].nome == persona.partenza {
				p = i
				break
			}
		}
		fmt.Printf("%v %v raggiunge %v \n", persona.colore, persona.nome, stazioni[p].nome)
		stazioni[p].coda <- persona

		// Mettiti in coda stazione partenza
		<-persona.notifica
		fmt.Printf("%v %v sale a %v \n", persona.colore, persona.nome, stazioni[p].nome)

		// Attende una notifica di arrivo dal treno
		<-persona.notifica
		fmt.Printf("%v %v scende a %v \n", persona.colore, persona.nome, persona.arrivo)

		// Eseguo il lavoro o dormo (simulato con uno sleep)
		time.Sleep(time.Second * time.Duration(persona.attesa))
		// [...]

		// Scambio partenza - arrivo e ripeto il ciclo
		persona.partenza, persona.arrivo = persona.arrivo, persona.partenza
	}
}
