package main

import (
	"fmt"
)

// thread.
func sendVal(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i
	}
	close(myIntChannel)
}

// main thread.
func main() {

	myIntChannel := make(chan int)

	go sendVal(myIntChannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel)
	}

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel)
	}

}
