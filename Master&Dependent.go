package main

import "fmt"

const N = 5
const FINE = -1

func ProducerForDep(outputChannel chan int, max int) {
	for i := 0; i <= max; i+=2 {
		outputChannel <- i
	}
	outputChannel <- FINE
}

func Dependent(idChan chan int, stopChannel chan bool ) {
	for {
		val := <-idChan
		if val == FINE {
			break
		}
		fmt.Printf("Begin of task %d", val)
		for i := 0; i < N; i++ {
			fmt.Printf("iteration %[1]d for task %[2]d", i, val)
		}
		fmt.Printf("End of task %d", val)
	}
	close(stopChannel)
}

func main() {
	fmt.Println("Insert Range limit: ")
	var rangeLimit int
	fmt.Scan(&rangeLimit)

	ch := make(chan int)
	stop := make(chan bool)
	go ProducerForDep(ch, rangeLimit)
	go Dependent(ch, stop)
	<-stop
}