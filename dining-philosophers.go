/**
 * You can find the problem here: https://en.wikipedia.org/wiki/Dining_philosophers_problem
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	name     string
	cutlery  chan bool
	neighbor *Philosopher
}

func makePhilosopher(name string, neighbor *Philosopher) *Philosopher {
	philosopher := &Philosopher{name, make(chan bool, 1), neighbor}
	philosopher.cutlery <- true
	return philosopher
}

func (philosopher *Philosopher) think() {
	fmt.Printf("%v is thinking.\n", philosopher.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (philosopher *Philosopher) eat() {
	fmt.Printf("%v is eating.\n", philosopher.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (philosopher *Philosopher) getCutlery() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	<-philosopher.cutlery
	select {
	case <-philosopher.neighbor.cutlery:
		fmt.Printf("%v got %v's cutlery.\n", philosopher.name, philosopher.neighbor.name)
		fmt.Printf("%v has two cutlery.\n", philosopher.name)
		return
	case <-timeout:
		philosopher.cutlery <- true
		philosopher.think()
		philosopher.getCutlery()
	}
}

func (philosopher *Philosopher) leaveCutlery() {
	philosopher.cutlery <- true
	philosopher.neighbor.cutlery <- true
}

func (philosopher *Philosopher) dine(announce chan *Philosopher) {
	philosopher.think()
	philosopher.getCutlery()
	philosopher.eat()
	philosopher.leaveCutlery()
	announce <- philosopher
}

func main() {
	names := []string{"Kant", "Heidegger", "Newton", "Hume", "Locke", "Marx", "Leibniz"}
	philosophers := make([]*Philosopher, len(names))
	var philosopher *Philosopher
	for i, name := range names {
		philosopher = makePhilosopher(name, philosopher)
		philosophers[i] = philosopher
	}
	philosophers[0].neighbor = philosopher
	fmt.Printf("There are %v philosophers sitting at a table.\n", len(philosophers))
	fmt.Println("They each have one cutlery, and must borrow from their neighbor to eat.")
	announce := make(chan *Philosopher)
	for _, philosopher := range philosophers {
		go philosopher.dine(announce)
	}
	for i := 0; i < len(names); i++ {
		philosopher := <-announce
		fmt.Printf("%v is done dining. \n", philosopher.name)
	}
}
