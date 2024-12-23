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

func main() {
	// Sync execution
	greet("*fast sync* Alex, nice to meet you!")
	greet("*fast sync* How are you?")
	slowGreet("*slow sync* What are you doing right now?")
	greet("*fast sync* Bye!")

	// Async execution using goroutines
	// With no output (because of the nature of go, we have to get result from them)
	go greet("*fast sync* Alex, nice to meet you!")
	go greet("*fast sync* How are you?")
	go slowGreet("*slow sync* What are you doing right now?")
	go greet("*fast sync* Bye!")
}
