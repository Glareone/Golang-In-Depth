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

	// Shorthand initialization
	var appUser2ShortHandInit = User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}

	var appUser3EmptyStruct = User{}

	// Calling User method
	fmt.Println("Calling appUser User method: ")
	appUser.OutputUserDetails()

	// Print user details
	fmt.Println("Using regular method below for appUser: ")
	outputUserDetails(appUser)

	fmt.Println("clear user (user's COPY!!!) details, which will not affect fields in the struct, only in 'Receiver Argument' copy:")
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	fmt.Println("this is how to properly clear user details fields:")
	appUser.ClearUserNameAsterisk()
	appUser.OutputUserDetails()

	// using output method with pointer in parameters
	// little comments below
	fmt.Println("appUser2ShortHandInit: ")
	outputUserDetailsPointer(&appUser2ShortHandInit)
	fmt.Println("appUser2ShortHandInit: ", appUser2ShortHandInit.firstName, appUser2ShortHandInit.lastName, appUser2ShortHandInit.birthDate)
	fmt.Println("appUser3EmptyStruct, User Struct to string:", appUser3EmptyStruct)
}

// "user" is a shallow copy of the struct we sent to the function calling it with "outputUserDetails(appUser)"
// in order to use a pointer we need to change the declaration
// it is a shallow(!) copy of the original struct
func outputUserDetails(user User) {
	// here we use the copy of original user Struct
	fmt.Println("appUser: ", user.firstName, user.lastName, user.birthDate, user.createdAt)
}

// using pointer on user
func outputUserDetailsPointer(user *User) {
	// technically, calling user.firstName without using asterisk there like *user.firstName
	// the technically correct way to use it would be using dereference: (*user).firstName
	// but Go allows us to use the shorthand
	fmt.Println("appUser: ", user.firstName, user.lastName, user.birthDate, user.createdAt)
	// returns the same, normally you can dereference the value, but Go allows us to use the shorthand from above
	fmt.Println("appUser using reference: ", (*user).firstName, (*user).lastName, (*user).birthDate, (*user).createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
