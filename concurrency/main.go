// entry point should be in package "main", otherwise it will not create and will not be executed
package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("hello! ", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("slow Hello! ", phrase)
}

// Second parameter here is Channel, which we use as "await" to let other functions know when operations
// within this function ARE DONE
func slowOneFuncWithChannel(information string, isDoneChannel chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println(information)

	// we send "true" to the channel
	// "<-" operator declares the direction of the information flow
	isDoneChannel <- true
}

func fastFuncWithOneChannel(phrase string, isDone chan bool) {
	fmt.Println("hello! ", phrase)
	isDone <- true
}

func main() {
	// Sync execution
	greet("*fast sync* Alex, nice to meet you!")
	greet("*fast sync* How are you?")
	slowGreet("*slow sync* What are you doing right now?")
	greet("*fast sync* Bye!")

	// =========== Async execution using goroutines ====================
	// With no output (because of the nature of go, we have to get result from them)
	// Here we do not await them. To await them we have to use "channels" using "make" keyword
	go greet("*async func without channel* Goroutine")
	go greet("*async func without channel* We do not await them")
	go slowGreet("*slow async func* We even can see no responses from them because we execute from the program faster")
	go greet("*fast async func* That's it!")

	// =========== Single Func single Channel ====================
	// We create channel which expects boolean to be returned
	var isDoneChannel = make(chan bool)
	go slowOneFuncWithChannel("*slow async with channel* execute function and await channel", isDoneChannel)
	// we await the response from the channel
	// without printing it, without anything, sending the result to nowhere
	// but we also can print it!

	//<-isDoneChannel
	fmt.Println(<-isDoneChannel, " result from the channel for single function!")

	// ============ Multiple Functions Single Channel ================
	var isDoneOneChannelMultiFunction = make(chan bool)
	go fastFuncWithOneChannel("*async fast with channel* 1 call", isDoneOneChannelMultiFunction)
	go slowOneFuncWithChannel("*slow async with channel* 2 function and await channel", isDoneOneChannelMultiFunction)
	go fastFuncWithOneChannel("*async fast with channel* 3 function", isDoneOneChannelMultiFunction)
	// *Race condition here* In this case we will not wait for all functions, so we will receive 1-3 results here without guarantees
	// fmt.Println(<-isDoneOneChannelMultiFunction, " result from the multiple functions!")

	// To await all of them we can await the channel 3 times because of 3 functions
	// obviously we receive them in order they finished
	<-isDoneOneChannelMultiFunction
	<-isDoneOneChannelMultiFunction
	<-isDoneOneChannelMultiFunction

	// ============ Multiple Functions Single Channel with Slice ===================
	var isDoneChannelSlice = make([]chan bool, 4)
	isDoneChannelSlice[0] = make(chan bool)
	go fastFuncWithOneChannel("*async fast with channel* 1 call", isDoneChannelSlice[0])
	isDoneChannelSlice[1] = make(chan bool)
	go slowOneFuncWithChannel("*slow async with channel* 2 function and await channel", isDoneChannelSlice[1])
	isDoneChannelSlice[2] = make(chan bool)
	go fastFuncWithOneChannel("*async fast with channel* 3 function", isDoneChannelSlice[2])
	isDoneChannelSlice[3] = make(chan bool)
	go fastFuncWithOneChannel("*async fast with channel* 4 function", isDoneChannelSlice[3])

	// await all functions
	for _, done := range isDoneChannelSlice {
		<-done
	}
}
