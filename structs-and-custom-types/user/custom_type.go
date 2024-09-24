package user

import "fmt"

// my own custom type customString and use alias "string"
// later I can use it in my application
// This means customString "inherits" all the properties and methods of string.
//
// 1. Purpose of Custom Types:
// Type Safety: You can use custom types to enforce stricter type checking. By defining a new type, you're signaling to the compiler (and to other developers)
// that these values, even though they might use the same underlying type, are meant to be used in distinct ways.
// Code Clarity: Custom types can make your code more readable and self-documenting. The type name itself conveys meaning and intent.
// Methods: You can define methods specifically on your custom types, even if they share an underlying type.
type customString string

// ALIASES
// Go also supports type aliases, which are more like nicknames for existing types:
type myString = string // MyString is now an alias for string

// The error "invalid receiver type 'string'" occurs because when you use (text myString), the compiler treats it as (text string),
// which is invalid because you cannot define a method with a receiver type from another package (string is from the fmt package).
//
//func (text myString) log() {
//	fmt.Println(text)
//}

func (text customString) log() {
	fmt.Println(text)
}

// Additional Example which demonstrates the difference how Golang treat different types
func myFunc() {
	var myString string = "Hello"
	var myCustomString customString = "World"

	// fmt.Println(myString + myCustomString)  // ERROR: cannot use myCustomString (variable of type customString) as type string
	fmt.Println(myString + string(myCustomString)) // Type conversion is required
}
