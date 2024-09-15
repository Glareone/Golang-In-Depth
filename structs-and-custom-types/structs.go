package main

import (
	"fmt"
	"time"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// "Struct Literal" (or also known as "Composite Literal") - initialization of custom type using curly brackets
	var appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)
}

func outputUserDetails(firstName string, lastName string, birthDate string) {

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
