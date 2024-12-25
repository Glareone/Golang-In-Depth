package main

import (
	"errors"
	"fmt"
	"time"
)

func FuncAsync(text string, isDoneChannel chan bool, errorChannel chan error) {
	time.Sleep(2 * time.Second)
	fmt.Println(text)
	isDoneChannel <- true
}

func FuncErrorAsync(text string, isDoneChannel chan bool, errorChannel chan error) {
	time.Sleep(2 * time.Second)
	fmt.Println(text, "ERROR HAPPENED!")
	var errorRaised = errors.New("an error occurred in FuncErrorAsync")

	if errorRaised != nil {
		errorChannel <- errorRaised
		// we use return here to skip next operations because error happened
		return
	}

	fmt.Println("Will never be written because of return above")
	isDoneChannel <- true
}

func main() {
	var isDoneChannels = make([]chan bool, 2)
	isDoneChannels[0] = make(chan bool)
	isDoneChannels[1] = make(chan bool)

	var errorsChannels = make([]chan error, 2)
	errorsChannels[0] = make(chan error)
	errorsChannels[1] = make(chan error)

	go FuncAsync("message: 1st function, no errors!", isDoneChannels[0], errorsChannels[0])
	go FuncErrorAsync("message: 2nd function, ERRORS!", isDoneChannels[1], errorsChannels[1])

	for index, _ := range isDoneChannels {
		// select structure is a special structure for channels
		select {
		case err := <-errorsChannels[index]:
			if err != nil {
				fmt.Println("Error description:", err)
			}
		case <-isDoneChannels[index]:
			fmt.Println(index, "number of function is Done!")
		}
	}

}
