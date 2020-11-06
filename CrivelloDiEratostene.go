package main

import (
	"fmt"
)
var primeNumbersCount = 0
const STOP = -1

/**
	At the end of all the numbers is sent the stop guard to close the goroutines(not so elegant)
	@param: channel to fill with all the odds integer to be processed
	@param: maximum limit picked by the user at the beginning of the program
	@return: void
 */
func Producer(outputChannel chan int, max int)  {
	for i := 3; i <= max; i+=2 {
		outputChannel <- i
	}
	outputChannel <- STOP
}

/**
	@param: channel with the integers given by producer
	@param: channel to stop the execution
	@return: void
 */
func Consumer(inputChannel chan int, stopChannel chan bool)  {
	for {
		toBeCheckedNum := <- inputChannel
		if toBeCheckedNum == STOP {
			break
		}
		fmt.Printf("Found the prime number %d\n", toBeCheckedNum)
		primeNumbersCount++

		outputChannel := make(chan int)

		//I use a separate goroutine in order to remove form my channel all the values that can be divided by the
		//current number so the next i know for sure that the next value is a prime number.
		go filter(toBeCheckedNum, inputChannel, outputChannel)
		inputChannel = outputChannel
	}
	close(stopChannel)
}

/**
	@param: int, prime number for Consumer func
	@param: channels to exchange data with Consumer func
 */
func filter(prime int, inputChannel, outputChannel chan int) {
	for {
		value := <-inputChannel
		if value == STOP {
			outputChannel <- value
			return
		}
		if (value % prime) != 0 {
			outputChannel <- value
		}
	}
}

func main() {
	fmt.Println("Insert Range limit: ")
	var rangeLimit int
	fmt.Scan(&rangeLimit)
	fmt.Println(rangeLimit)

	ch := make(chan int)
	stop := make(chan bool)
	go Producer(ch, rangeLimit)
	go Consumer(ch, stop)
	<-stop

	fmt.Printf("I have found %d prime numbers\n", primeNumbersCount)
	fmt.Printf("The number of goroutines generated is: %d\n", primeNumbersCount+2)
}
