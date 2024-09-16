package main

import (
	"fmt"
	"structs-and-custom-types/user"
	"time"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// "Struct Literal" (or also known as "Composite Literal") - initialization of custom type using curly brackets
	var appUser = user.User{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthdate,
		CreatedAt: time.Now(),
	}

	// Shorthand initialization
	var appUser2ShortHandInit = user.User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}

	// Empty instance
	var appUser3EmptyStruct = user.User{}

	// initialization using constructor
	// error is omitted, therefore _
	var appUserConstructor, _ = user.NewUser(userFirstName, userLastName, userBirthdate)
	fmt.Println("appUser2ShortHandInit: ", appUserConstructor.FirstName, appUserConstructor.LastName, appUserConstructor.BirthDate)

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
	fmt.Println("appUser2ShortHandInit: ", appUser2ShortHandInit.FirstName, appUser2ShortHandInit.LastName, appUser2ShortHandInit.BirthDate)
	fmt.Println("appUser3EmptyStruct, User Struct to string:", appUser3EmptyStruct)
}

// "user" is a shallow copy of the struct we sent to the function calling it with "outputUserDetails(appUser)"
// in order to use a pointer we need to change the declaration
// it is a shallow(!) copy of the original struct
func outputUserDetails(user user.User) {
	// here we use the copy of original user Struct
	fmt.Println("appUser: ", user.FirstName, user.LastName, user.BirthDate, user.CreatedAt)
}

// using pointer on user
func outputUserDetailsPointer(user *user.User) {
	// technically, calling user.firstName without using asterisk there like *user.firstName
	// the technically correct way to use it would be using dereference: (*user).firstName
	// but Go allows us to use the shorthand
	fmt.Println("appUser: ", user.FirstName, user.LastName, user.BirthDate, user.CreatedAt)
	// returns the same, normally you can dereference the value, but Go allows us to use the shorthand from above
	fmt.Println("appUser using reference: ", (*user).FirstName, (*user).LastName, (*user).BirthDate, (*user).CreatedAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	// Scanln stops the scanning of the value if you click the Enter.
	// otherwise Enter will be ignored
	fmt.Scanln(&value)
	return value
}
