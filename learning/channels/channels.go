package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	fmt.Println("processing number", <- numChan)

}

func main() {

	numChannel := make(chan int)
	
	go processNum(numChannel)

	numChannel <- 5


	time.Sleep(time.Second * 2 )


	// messageChannel := make(chan string)

	// messageChannel <- "ping"   //sending data to message channel

	// msg:= <- messageChannel  // receiving data from message channel

	// fmt.Println(msg)
}